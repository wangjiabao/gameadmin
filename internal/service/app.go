package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	pb "game/api/app/v1"
	"game/internal/biz"
	"game/internal/conf"
	"game/internal/pkg/middleware/auth"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"math/big"
	"regexp"
	"strconv"
	"time"
)

type AppService struct {
	pb.UnimplementedAppServer
	log *log.Helper
	ac  *biz.AppUsecase
	ca  *conf.Auth
}

func NewAppService(ac *biz.AppUsecase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{
		ac:  ac,
		log: log.NewHelper(logger),
		ca:  ca,
	}
}

var ethClient *ethclient.Client

func init() {
	var err error
	ethClient, err = ethclient.Dial("https://bsc-dataseed4.binance.org/")
	if err != nil {
		panic("eth client err")
	}
}

func addressCheck(addressParam string) (bool, error) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(addressParam) {
		return false, nil
	}

	var (
		err      error
		bytecode []byte
	)

	if nil == ethClient {
		ethClient, err = ethclient.Dial("https://bsc-dataseed4.binance.org/")
		if err != nil {
			panic("eth client err")
		}
	}

	// a random user account address
	address := common.HexToAddress(addressParam)
	bytecode, err = ethClient.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		return false, err
	}

	if len(bytecode) > 0 {
		return false, nil
	}

	return true, nil
}

// TestSign testSign.
func (a *AppService) TestSign(ctx context.Context, req *pb.TestSignRequest) (*pb.TestSignReply, error) {
	privateKey, err := crypto.HexToECDSA(req.Secret)
	if err != nil {
		return &pb.TestSignReply{Sign: ""}, err
	}

	data := []byte(req.SignContent)
	hash := accounts.TextHash(data)

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return &pb.TestSignReply{Sign: ""}, err
	}

	return &pb.TestSignReply{Sign: string(signature)}, nil
}

func verifySig(sigHex string, msg []byte) (bool, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("签名，捕获到异常:")
		}
	}()

	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false, ""
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	sigPublicKeyBytes := crypto.FromECDSAPub(recovered)
	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKeyBytes, msg, signatureNoRecoverID)
	return verified, recoveredAddr.Hex()
}

// EthAuthorize ethAuthorize.
func (a *AppService) EthAuthorize(ctx context.Context, req *pb.EthAuthorizeRequest) (*pb.EthAuthorizeReply, error) {
	userAddress := req.SendBody.Address // 以太坊账户

	// 验证
	var (
		res bool
		err error
	)
	res, err = addressCheck(userAddress)
	if nil != err {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "地址验证失败",
		}, nil
	}

	if !res {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "地址格式错误",
		}, nil
	}

	if 20 >= len(req.SendBody.Sign) {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "签名错误",
		}, nil
	}

	var (
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(userAddress))
	if !res || addressFromSign != userAddress {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "地址签名错误",
		}, nil
	}

	// 根据地址查询用户，不存在时则创建
	var (
		user *biz.User
		msg  string
	)
	user, err, msg = a.ac.GetExistUserByAddressOrCreate(ctx, userAddress, req)
	if err != nil {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: msg,
		}, nil
	}

	claims := auth.CustomClaims{
		Address: user.Address,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(48 * time.Hour)), // 2天过期
			Issuer:    "game",
		},
	}

	var (
		token string
	)
	token, err = auth.CreateToken(claims, a.ca.JwtKey)
	if err != nil {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "生成token失败",
		}, nil
	}

	return &pb.EthAuthorizeReply{
		Token:  token,
		Status: "ok",
	}, nil
}

// UserInfo userInfo.
func (a *AppService) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserInfoReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserInfoReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserInfoReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserInfoReply{Status: "无效token"}, nil
	}

	return a.ac.UserInfo(ctx, address)
}

// UserRecommend userRecommend.
func (a *AppService) UserRecommend(ctx context.Context, req *pb.UserRecommendRequest) (*pb.UserRecommendReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserRecommendReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserRecommendReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserRecommendReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserRecommendReply{Status: "无效token"}, nil
	}

	return a.ac.UserRecommend(ctx, address, req)
}

// UserRecommendL userRecommendL.
func (a *AppService) UserRecommendL(ctx context.Context, req *pb.UserRecommendLRequest) (*pb.UserRecommendLReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserRecommendLReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserRecommendLReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserRecommendLReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserRecommendLReply{Status: "无效token"}, nil
	}

	return a.ac.UserRecommendL(ctx, address, req)
}

// UserLand userLand.
func (a *AppService) UserLand(ctx context.Context, req *pb.UserLandRequest) (*pb.UserLandReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserLandReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserLandReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserLandReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserLandReply{Status: "无效token"}, nil
	}

	return a.ac.UserLand(ctx, address, req)
}

// UserStakeRewardList userStakeRewardList.
func (a *AppService) UserStakeRewardList(ctx context.Context, req *pb.UserStakeRewardListRequest) (*pb.UserStakeRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
	}

	return a.ac.UserStakeRewardList(ctx, address, req)
}

// UserBoxList userBoxList.
func (a *AppService) UserBoxList(ctx context.Context, req *pb.UserBoxListRequest) (*pb.UserBoxListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBoxListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBoxListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBoxListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBoxListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBoxList(ctx, address, req)
}

// UserBackList userBackList.
func (a *AppService) UserBackList(ctx context.Context, req *pb.UserBackListRequest) (*pb.UserBackListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}

// UserMarketSeedList userMarketSeedList.
func (a *AppService) UserMarketSeedList(ctx context.Context, req *pb.UserMarketSeedListRequest) (*pb.UserMarketSeedListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserMarketSeedListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserMarketSeedListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserMarketSeedListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserMarketSeedListReply{Status: "无效token"}, nil
	}

	return a.ac.UserMarketSeedList(ctx, address, req)
}

// UserMarketLandList userMarketLandList.
func (a *AppService) UserMarketLandList(ctx context.Context, req *pb.UserMarketLandListRequest) (*pb.UserMarketLandListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserMarketLandListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserMarketLandListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserMarketLandListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserMarketLandListReply{Status: "无效token"}, nil
	}

	return a.ac.UserMarketLandList(ctx, address, req)
}

// UserMarketPropList userMarketPropList.
func (a *AppService) UserMarketPropList(ctx context.Context, req *pb.UserMarketPropListRequest) (*pb.UserMarketPropListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserMarketPropListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserMarketPropListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserMarketPropListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserMarketPropListReply{Status: "无效token"}, nil
	}

	return a.ac.UserMarketPropList(ctx, address, req)
}

// UserMarketRentLandList userMarketRentLandList.
func (a *AppService) UserMarketRentLandList(ctx context.Context, req *pb.UserMarketRentLandListRequest) (*pb.UserMarketRentLandListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserMarketRentLandListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserMarketRentLandListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserMarketRentLandListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserMarketRentLandListReply{Status: "无效token"}, nil
	}

	return a.ac.UserMarketRentLandList(ctx, address, req)
}

// UserMyMarketList userMyMarketList.
func (a *AppService) UserMyMarketList(ctx context.Context, req *pb.UserMyMarketListRequest) (*pb.UserMyMarketListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserMyMarketListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserMyMarketListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserMyMarketListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserMyMarketListReply{Status: "无效token"}, nil
	}

	return a.ac.UserMyMarketList(ctx, address, req)
}

// UserNoticeList NoticeList.
func (a *AppService) UserNoticeList(ctx context.Context, req *pb.UserNoticeListRequest) (*pb.UserNoticeListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserNoticeListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserNoticeListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserNoticeListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserNoticeListReply{Status: "无效token"}, nil
	}

	return a.ac.UserNoticeList(ctx, address, req)
}

// UserStakeGitRewardList UserStakeGitRewardList.
func (a *AppService) UserStakeGitRewardList(ctx context.Context, req *pb.UserStakeGitRewardListRequest) (*pb.UserStakeGitRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserStakeGitRewardListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserStakeGitRewardListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserStakeGitRewardListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserStakeGitRewardListReply{Status: "无效token"}, nil
	}

	return a.ac.UserStakeGitRewardList(ctx, address, req)
}

// UserStakeGitStakeList UserStakeGitStakeList.
func (a *AppService) UserStakeGitStakeList(ctx context.Context, req *pb.UserStakeGitStakeListRequest) (*pb.UserStakeGitStakeListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserStakeGitStakeListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserStakeGitStakeListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserStakeGitStakeListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserStakeGitStakeListReply{Status: "无效token"}, nil
	}

	return a.ac.UserStakeGitStakeList(ctx, address, req)
}

// UserIndexList UserIndexList.
func (a *AppService) UserIndexList(ctx context.Context, req *pb.UserIndexListRequest) (*pb.UserIndexListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserIndexListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserIndexListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserIndexListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserIndexListReply{Status: "无效token"}, nil
	}

	return a.ac.UserIndexList(ctx, address, req)
}

// UserOrderList  userOrderList.
func (a *AppService) UserOrderList(ctx context.Context, req *pb.UserOrderListRequest) (*pb.UserOrderListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserOrderListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserOrderListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserOrderListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserOrderListReply{Status: "无效token"}, nil
	}

	return a.ac.UserOrderList(ctx, address, req)
}

func (a *AppService) BuyBox(ctx context.Context, req *pb.BuyBoxRequest) (*pb.BuyBoxReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.BuyBoxReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.BuyBoxReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.BuyBoxReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.BuyBoxReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.BuyBoxReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.BuyBox(ctx, address, req)
}

func (a *AppService) OpenBox(ctx context.Context, req *pb.OpenBoxRequest) (*pb.OpenBoxReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.OpenBoxReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.OpenBoxReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.OpenBoxReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.OpenBoxReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.OpenBoxReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.OpenBox(ctx, address, req)
}

func (a *AppService) LandPlay(ctx context.Context, req *pb.LandPlayRequest) (*pb.LandPlayReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlayReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlayReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlayReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlayReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlayReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlay(ctx, address, req)
}

func (a *AppService) LandPlayOne(ctx context.Context, req *pb.LandPlayOneRequest) (*pb.LandPlayOneReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlayOneReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlayOneReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlayOneReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlayOneReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlayOneReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlayOne(ctx, address, req)
}

func (a *AppService) LandPlayTwo(ctx context.Context, req *pb.LandPlayTwoRequest) (*pb.LandPlayTwoReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlayTwoReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlayTwoReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlayTwoReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlayTwoReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlayTwoReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlayTwo(ctx, address, req)
}

func (a *AppService) LandPlayThree(ctx context.Context, req *pb.LandPlayThreeRequest) (*pb.LandPlayThreeReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlayThreeReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlayThreeReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlayThreeReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlayThreeReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlayThreeReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlayThree(ctx, address, req)
}

func (a *AppService) LandPlayFour(ctx context.Context, req *pb.LandPlayFourRequest) (*pb.LandPlayFourReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlayFourReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlayFourReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlayFourReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlayFourReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlayFourReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlayFour(ctx, address, req)
}

func (a *AppService) LandPlayFive(ctx context.Context, req *pb.LandPlayFiveRequest) (*pb.LandPlayFiveReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlayFiveReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlayFiveReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlayFiveReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlayFiveReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlayFiveReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlayFive(ctx, address, req)
}

func (a *AppService) LandPlaySix(ctx context.Context, req *pb.LandPlaySixRequest) (*pb.LandPlaySixReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlaySixReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlaySixReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlaySixReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlaySixReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlaySixReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlaySix(ctx, address, req)
}

func (a *AppService) LandPlaySeven(ctx context.Context, req *pb.LandPlaySevenRequest) (*pb.LandPlaySevenReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandPlaySevenReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandPlaySevenReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandPlaySevenReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandPlaySevenReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandPlaySevenReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandPlaySeven(ctx, address, req)
}

func (a *AppService) Buy(ctx context.Context, req *pb.BuyRequest) (*pb.BuyReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.BuyReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.BuyReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.BuyReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.BuyReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.BuyReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.Buy(ctx, address, req)
}

func (a *AppService) Sell(ctx context.Context, req *pb.SellRequest) (*pb.SellReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.SellReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.SellReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.SellReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.SellReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.SellReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.Sell(ctx, address, req)
}

func (a *AppService) StakeGit(ctx context.Context, req *pb.StakeGitRequest) (*pb.StakeGitReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.StakeGitReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.StakeGitReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.StakeGitReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.StakeGitReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.StakeGitReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.StakeGit(ctx, address, req)
}

func (a *AppService) RentLand(ctx context.Context, req *pb.RentLandRequest) (*pb.RentLandReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.RentLandReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.RentLandReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.RentLandReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.RentLandReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.RentLandReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.RentLand(ctx, address, req)
}

func (a *AppService) LandAddOutRate(ctx context.Context, req *pb.LandAddOutRateRequest) (*pb.LandAddOutRateReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.LandAddOutRateReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.LandAddOutRateReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.LandAddOutRateReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.LandAddOutRateReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.LandAddOutRateReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.LandAddOutRate(ctx, address, req)
}

func (a *AppService) GetLand(ctx context.Context, req *pb.GetLandRequest) (*pb.GetLandReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.GetLandReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.GetLandReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.GetLandReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.GetLandReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.GetLandReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.GetLand(ctx, address, req)
}

func (a *AppService) StakeGet(ctx context.Context, req *pb.StakeGetRequest) (*pb.StakeGetReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.StakeGetReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.StakeGetReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.StakeGetReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.StakeGetReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.StakeGetReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.StakeGet(ctx, address, req)
}

func (a *AppService) StakeGetPlay(ctx context.Context, req *pb.StakeGetPlayRequest) (*pb.StakeGetPlayReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.StakeGetPlayReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.StakeGetPlayReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.StakeGetPlayReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.StakeGetPlayReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.StakeGetPlayReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.StakeGetPlay(ctx, address, req)
}

func (a *AppService) Exchange(ctx context.Context, req *pb.ExchangeRequest) (*pb.ExchangeReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.ExchangeReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.ExchangeReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.ExchangeReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.ExchangeReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.ExchangeReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.Exchange(ctx, address, req)
}

func (a *AppService) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.WithdrawReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.WithdrawReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.WithdrawReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.WithdrawReply{Status: "无效token"}, nil
	}

	var (
		res             bool
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(address))
	if !res || addressFromSign != address {
		return &pb.WithdrawReply{
			Status: "地址签名错误",
		}, nil
	}

	return a.ac.Withdraw(ctx, address, req)
}

func (a *AppService) SetGiw(ctx context.Context, req *pb.SetGiwRequest) (*pb.SetGiwReply, error) {
	return a.ac.SetGiw(ctx, req)
}

func (a *AppService) SetGit(ctx context.Context, req *pb.SetGitRequest) (*pb.SetGitReply, error) {
	return a.ac.SetGit(ctx, req)
}

func (a *AppService) SetLand(ctx context.Context, req *pb.SetLandRequest) (*pb.SetLandReply, error) {
	return a.ac.SetLand(ctx, req)
}

func (a *AppService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginReply, error) {
	claims := auth.CustomClaims{
		Address: "admin",
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(48 * time.Hour)), // 2天过期
			Issuer:    "game",
		},
	}

	var (
		err   error
		token string
	)
	token, err = auth.CreateToken(claims, a.ca.JwtKey)
	if err != nil {
		return &pb.AdminLoginReply{
			Token:  "",
			Status: "生成token失败",
		}, nil
	}

	return a.ac.AdminLogin(ctx, req, token)
}

func (a *AppService) AdminUserList(ctx context.Context, req *pb.AdminUserListRequest) (*pb.AdminUserListReply, error) {
	return a.ac.AdminUserList(ctx, req)
}

func (a *AppService) AdminUserRecommend(ctx context.Context, req *pb.AdminUserRecommendRequest) (*pb.AdminUserRecommendReply, error) {
	return a.ac.AdminRecommendList(ctx, req)
}

func (a *AppService) AdminRecordList(ctx context.Context, req *pb.RecordListRequest) (*pb.RecordListReply, error) {
	return a.ac.AdminRecordList(ctx, req)
}

func (a *AppService) AdminWithdrawList(ctx context.Context, req *pb.AdminWithdrawListRequest) (*pb.AdminWithdrawListReply, error) {
	return a.ac.AdminWithdrawList(ctx, req)
}

func (a *AppService) AdminLandConfigList(ctx context.Context, req *pb.AdminLandConfigRequest) (*pb.AdminLandConfigReply, error) {
	return a.ac.AdminLandConfig(ctx, req)
}

func (a *AppService) AdminLandConfigSet(ctx context.Context, req *pb.AdminLandConfigSetRequest) (*pb.AdminLandConfigSetReply, error) {
	return a.ac.AdminLandConfigSet(ctx, req)
}

func (a *AppService) AdminSeedConfigList(ctx context.Context, req *pb.AdminSeedConfigRequest) (*pb.AdminSeedConfigReply, error) {
	return a.ac.AdminSeedConfig(ctx, req)
}

func (a *AppService) AdminSeedConfigSet(ctx context.Context, req *pb.AdminSeedConfigSetRequest) (*pb.AdminSeedConfigSetReply, error) {
	return a.ac.AdminSeedConfigSet(ctx, req)
}

func (a *AppService) AdminPropConfigList(ctx context.Context, req *pb.AdminPropConfigRequest) (*pb.AdminPropConfigReply, error) {
	return a.ac.AdminPropConfig(ctx, req)
}

func (a *AppService) AdminPropConfigSet(ctx context.Context, req *pb.AdminPropConfigSetRequest) (*pb.AdminPropConfigSetReply, error) {
	return a.ac.AdminPropConfigSet(ctx, req)
}

func (a *AppService) AdminGetBox(ctx context.Context, req *pb.AdminGetBoxRequest) (*pb.AdminGetBoxReply, error) {
	return a.ac.AdminGetBox(ctx, req)
}

func (a *AppService) AdminSetBox(ctx context.Context, req *pb.AdminSetBoxRequest) (*pb.AdminSetBoxReply, error) {
	return a.ac.AdminSetBox(ctx, req)
}

func (a *AppService) AdminGetConfig(ctx context.Context, req *pb.AdminGetConfigRequest) (*pb.AdminGetConfigReply, error) {
	return a.ac.AdminGetConfig(ctx, req)
}

func (a *AppService) AdminSetConfig(ctx context.Context, req *pb.AdminSetConfigRequest) (*pb.AdminSetConfigReply, error) {
	return a.ac.AdminSetConfig(ctx, req)
}

func (a *AppService) AdminSetGiw(ctx context.Context, req *pb.AdminSetGiwRequest) (*pb.AdminSetGiwReply, error) {
	return a.ac.AdminSetGiw(ctx, req)
}

func (a *AppService) AdminSetGit(ctx context.Context, req *pb.AdminSetGitRequest) (*pb.AdminSetGitReply, error) {
	return a.ac.AdminSetGit(ctx, req)
}

func (a *AppService) AdminDaily(ctx context.Context, req *pb.AdminDailyRequest) (*pb.AdminDailyReply, error) {
	return a.ac.AdminDaily(ctx, req)
}

func (a *AppService) AdminSetLand(ctx context.Context, req *pb.AdminSetLandRequest) (*pb.AdminSetLandReply, error) {
	return a.ac.AdminSetLand(ctx, req)
}

func (a *AppService) AdminSetProp(ctx context.Context, req *pb.AdminSetPropRequest) (*pb.AdminSetPropReply, error) {
	return a.ac.AdminSetProp(ctx, req)
}

func (a *AppService) AdminSetSeed(ctx context.Context, req *pb.AdminSetSeedRequest) (*pb.AdminSetSeedReply, error) {
	return a.ac.AdminSetSeed(ctx, req)
}

func (a *AppService) AdminSetBuyLand(ctx context.Context, req *pb.AdminSetBuyLandRequest) (*pb.AdminSetBuyLandReply, error) {
	return a.ac.AdminSetBuyLand(ctx, req)
}

func getUserLength(address string) (int64, error) {
	url1 := "https://bsc-dataseed4.binance.org/"

	var balInt int64
	for i := 0; i < 5; i++ {
		if 1 == i {
			url1 = "https://binance.llamarpc.com/"
		} else if 2 == i {
			url1 = "https://bscrpc.com/"
		} else if 3 == i {
			url1 = "https://bsc-pokt.nodies.app/"
		} else if 4 == i {
			url1 = "https://data-seed-prebsc-1-s3.binance.org:8545/"
		}

		client, err := ethclient.Dial(url1)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewBuySomething(tokenAddress, client)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		bals, err := instance.GetUserLength(&bind.CallOpts{})
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		balInt = bals.Int64()
		break
	}

	return balInt, nil
}

type userDeposit struct {
	Address string
	Amount  int64
}

func getUserInfo(start int64, end int64, address string) ([]*userDeposit, error) {
	url1 := "https://bsc-dataseed4.binance.org/"

	var (
		bals  []common.Address
		bals2 []*big.Int
	)
	users := make([]*userDeposit, 0)

	for i := 0; i < 5; i++ {
		if 1 == i {
			url1 = "https://binance.llamarpc.com/"
		} else if 2 == i {
			url1 = "https://bscrpc.com/"
		} else if 3 == i {
			url1 = "https://bsc-pokt.nodies.app/"
		} else if 4 == i {
			url1 = "https://data-seed-prebsc-1-s3.binance.org:8545/"
		}

		client, err := ethclient.Dial(url1)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewBuySomething(tokenAddress, client)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		bals, err = instance.GetUsersByIndex(&bind.CallOpts{}, new(big.Int).SetInt64(start), new(big.Int).SetInt64(end))
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		break
	}

	for i := 0; i < 5; i++ {
		if 1 == i {
			url1 = "https://binance.llamarpc.com/"
		} else if 2 == i {
			url1 = "https://bscrpc.com/"
		} else if 3 == i {
			url1 = "https://bsc-pokt.nodies.app/"
		} else if 4 == i {
			url1 = "https://data-seed-prebsc-1-s3.binance.org:8545/"
		}

		client, err := ethclient.Dial(url1)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		tokenAddress := common.HexToAddress(address)
		instance, err := NewBuySomething(tokenAddress, client)
		if err != nil {
			fmt.Println(nil, err)
			continue
		}

		bals2, err = instance.GetUsersAmountByIndex(&bind.CallOpts{}, new(big.Int).SetInt64(start), new(big.Int).SetInt64(end))
		if err != nil {
			fmt.Println(err)
			//url1 = "https://bsc-dataseed4.binance.org"
			continue
		}

		break
	}

	if len(bals) != len(bals2) {
		fmt.Println("数量不一致，错误")
		return users, nil
	}

	for k, v := range bals {
		users = append(users, &userDeposit{
			Address: v.String(),
			Amount:  bals2[k].Int64(),
		})
	}

	return users, nil
}

func (a *AppService) AdminDeposit(ctx context.Context, req *pb.AdminDepositRequest) (*pb.AdminDepositReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)

	for i := 1; i <= 10; i++ {
		var (
			depositUsdtResult []*userDeposit
			depositUsers      map[string]*biz.User
			userLength        int64
			last              int64
			err               error
		)

		last, err = a.ac.GetEthUserRecordLast(ctx)
		if nil != err {
			fmt.Println(err)
			continue
		}

		if -1 == last {
			fmt.Println(err)
			continue
		}

		userLength, err = getUserLength("0xb3f58887f4Ce734DBd1A7dAa6A10e7901e641c7f")
		if nil != err {
			fmt.Println(err)
		}

		if -1 == userLength {
			continue
		}

		if 0 == userLength {
			break
		}

		if last >= userLength {
			break
		}

		depositUsdtResult, err = getUserInfo(last, userLength-1, "0xb3f58887f4Ce734DBd1A7dAa6A10e7901e641c7f")
		if nil != err {
			break
		}

		now := time.Now().UTC()
		//fmt.Println(now, end)
		if end.Before(now) {
			break
		}

		if 0 >= len(depositUsdtResult) {
			break
		}

		fromAccount := make([]string, 0)
		for _, vUser := range depositUsdtResult {
			fromAccount = append(fromAccount, vUser.Address)
		}

		depositUsers, err = a.ac.GetUserByAddress(ctx, fromAccount)
		if nil != depositUsers {
			// 统计开始
			for _, vUser := range depositUsdtResult { // 主查usdt
				if _, ok := depositUsers[vUser.Address]; !ok { // 用户不存在
					continue
				}

				var (
					tmpValue uint64
				)

				if 100 <= vUser.Amount {
					tmpValue = uint64(vUser.Amount)
				} else {
					continue
				}

				// 充值
				err = a.ac.DepositNew(ctx, &biz.EthRecord{ // 两种币的记录
					UserId:  depositUsers[vUser.Address].ID,
					Address: vUser.Address,
					Amount:  tmpValue,
					Last:    uint64(userLength),
				})
				if nil != err {
					fmt.Println(err)
				}
			}
		}

		time.Sleep(5 * time.Second)
	}

	return nil, nil
}

func (a *AppService) AdminWithdraw(ctx context.Context, req *pb.AdminWithdrawRequest) (*pb.AdminWithdrawReply, error) {
	end := time.Now().UTC().Add(45 * time.Second)
	for {
		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			withdraw *biz.Withdraw
			users    map[uint64]*biz.User
			err      error
		)

		withdraw, err = a.ac.GetWithdrawPassOrRewardedFirst(ctx)
		if nil == withdraw {
			break
		}

		if 0 >= withdraw.RelAmount {
			continue
		}

		userIds := []uint64{withdraw.UserId}
		users, err = a.ac.GetUserByUserIds(ctx, userIds)
		if nil != err {
			fmt.Println(err)
			continue
		}

		if _, ok := users[withdraw.UserId]; !ok {
			continue
		}

		err = a.ac.UpdateWithdrawDoing(ctx, withdraw.ID)
		if nil != err {
			fmt.Println(err)
			continue
		}

		tmpUrl1 := "https://bsc-dataseed4.binance.org/"

		withDrawAmount := strconv.FormatUint(withdraw.RelAmount, 10) + "000000000000000000"
		if len(withDrawAmount) <= 15 {
			fmt.Println(withDrawAmount, withdraw)
			err = a.ac.UpdateWithdrawSuccess(ctx, withdraw.ID)
			fmt.Println(err)
			continue
		}

		for i := 0; i <= 5; i++ {
			_, err = toToken("", users[withdraw.UserId].Address, withDrawAmount, "0x36081a3d9d4abA91D2075e4aD980c740e1B9BE76", tmpUrl1)
			if err == nil {
				err = a.ac.UpdateWithdrawSuccess(ctx, withdraw.ID)
				fmt.Println(err)
				break
			} else {
				fmt.Println(err)
				if 0 == i {
					tmpUrl1 = "https://bsc-dataseed1.binance.org"
				} else if 1 == i {
					tmpUrl1 = "https://bsc-dataseed3.binance.org"
				} else if 2 == i {
					tmpUrl1 = "https://bsc-dataseed2.binance.org"
				} else if 3 == i {
					tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com/"
				} else if 4 == i {
					tmpUrl1 = "https://bsc-dataseed.binance.org"
				}
				fmt.Println(33331, err, users[withdraw.UserId].Address, withDrawAmount)
				time.Sleep(3 * time.Second)
			}
		}

	}

	return nil, nil
}

func toToken(userPrivateKey string, toAccount string, withdrawAmount string, withdrawTokenAddress string, url1 string) (string, error) {
	client, err := ethclient.Dial(url1)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress(withdrawTokenAddress)
	instance, err := NewDfil(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tmpWithdrawAmount, _ := new(big.Int).SetString(withdrawAmount, 10)
	_, err = instance.Transfer(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasLimit: 0,
	}, common.HexToAddress(toAccount), tmpWithdrawAmount)
	if err != nil {
		return "", err
	}

	return "", nil
}
