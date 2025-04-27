package biz

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	pb "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Pagination struct {
	PageNum  int
	PageSize int
}

type User struct {
	ID               uint64
	Address          string
	Level            uint64
	Giw              float64
	GiwAdd           float64
	Git              float64
	Total            float64
	TotalOne         float64
	TotalTwo         float64
	TotalThree       float64
	RewardOne        float64
	RewardTwo        float64
	RewardThree      float64
	RewardTwoOne     float64
	RewardTwoTwo     float64
	RewardTwoThree   float64
	RewardThreeOne   float64
	RewardThreeTwo   float64
	RewardThreeThree float64
	LastRewardTotal  float64
	Location         float64
	Recommend        float64
	RecommendTwo     float64
	Area             float64
	AreaTwo          float64
	All              float64
	Amount           float64
	AmountGet        float64
	AmountUsdt       float64
	MyTotalAmount    float64
	OutNum           uint64
	Vip              uint64
	VipAdmin         uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type BoxRecord struct {
	ID        uint64
	UserId    uint64
	Num       uint64
	GoodId    uint64
	GoodType  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}

type UserRecommend struct {
	ID            uint64
	UserId        uint64
	RecommendCode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type StakeGit struct {
	ID        uint64
	UserId    uint64
	Status    uint64
	Amount    float64
	Reward    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reward struct {
	ID        uint64
	UserId    uint64
	Reason    uint64
	One       uint64
	Two       uint64
	Three     float64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Seed struct {
	ID           uint64
	UserId       uint64
	SeedId       uint64
	Name         string
	OutAmount    float64
	OutOverTime  uint64
	OutMaxAmount float64
	OutMinAmount float64
	Status       uint64
	SellAmount   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PropInfo struct {
	ID       uint64
	PropType uint64

	// 化肥相关字段
	OneOne uint64
	OneTwo uint64

	// 铲子相关字段
	TwoOne uint64
	TwoTwo float64

	// 水相关字段
	ThreeOne uint64

	// 除虫剂相关字段
	FourOne uint64

	// 手套相关字段
	FiveOne uint64
	GetRate float64

	// 时间字段
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SeedInfo struct {
	ID           uint64
	Name         string
	OutMinAmount float64
	OutMaxAmount float64
	GetRate      float64
	OutOverTime  uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Land struct {
	ID             uint64
	UserId         uint64
	Level          uint64
	OutPutRate     float64
	RentOutPutRate float64
	MaxHealth      uint64
	PerHealth      uint64
	LimitDate      uint64
	Status         uint64
	LocationNum    uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	One            uint64
	Two            uint64
	Three          uint64
	SellAmount     float64
}

type LandInfo struct {
	ID                uint64
	Level             uint64
	OutPutRateMax     float64
	OutPutRateMin     float64
	RentOutPutRateMax float64
	MaxHealth         uint64
	PerHealth         uint64
	LimitDateMax      uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type LandUserUse struct {
	ID           uint64
	LandId       uint64
	Level        uint64
	UserId       uint64
	OwnerUserId  uint64
	SeedId       uint64
	SeedTypeId   uint64
	Status       uint64
	BeginTime    uint64
	TotalTime    uint64
	OverTime     uint64
	OutMaxNum    float64
	OutNum       float64
	InsectStatus uint64
	OutSubNum    float64
	StealNum     float64
	StopStatus   uint64
	StopTime     uint64
	SubTime      uint64
	UseChan      uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	One          uint64
	Two          uint64
}

type ExchangeRecord struct {
	ID        uint64
	UserId    int64
	Git       float64
	Giw       float64
	Fee       float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Market struct {
	ID        uint64
	UserId    uint64
	GoodId    uint64
	GoodType  int
	Amount    float64
	Status    int
	GetUserId uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Notice struct {
	ID            uint64
	UserId        uint64
	NoticeContent string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Prop struct {
	ID         uint64
	UserId     uint64
	Status     int
	PropType   int
	OneOne     int
	OneTwo     int
	TwoOne     int
	TwoTwo     float64
	SellAmount float64
	ThreeOne   int
	FourOne    int
	FiveOne    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type StakeGet struct {
	ID        uint64
	UserId    uint64
	Status    int
	StakeRate float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGetPlayRecord struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	Reward    float64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGetRecord struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	StakeRate float64
	Total     float64
	StakeType int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGetTotal struct {
	ID        uint64
	Amount    float64
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGitRecord struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	StakeType int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Withdraw struct {
	ID        uint64
	UserId    uint64
	Amount    uint64
	RelAmount uint64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EthRecord struct {
	ID        uint64
	UserId    uint64
	Amount    uint64
	Last      uint64
	Address   string
	Coin      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BuyLand struct {
	ID        uint64
	Amount    float64
	Status    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	AmountTwo float64
	Limit     uint64
	Level     uint64
}

type RandomSeed struct {
	ID        uint64
	Scene     uint64
	SeedValue uint64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type BuyLandRecord struct {
	ID        uint64
	BuyLandID uint64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    uint64
	UserID    uint64
}

type Admin struct {
	ID       int64
	Password string
	Account  string
	Type     string
}

type RewardTwo struct {
	ID        uint64
	UserId    uint64
	Reason    uint64
	One       uint64
	Two       uint64
	Three     float64
	Amount    float64
	Four      string
	Five      float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*User, error)
	GetUserByAddresses(ctx context.Context, Addresses []string) (map[string]*User, error)
	GetUserById(ctx context.Context, id uint64) (*User, error)
	GetUserPageCount(ctx context.Context, address string) (int64, error)
	GetWithdrawPageCount(ctx context.Context, userId uint64) (int64, error)
	GetRecordPageCount(ctx context.Context, address string) (int64, error)
	GetRecordPageCountTwo(ctx context.Context, address string) (int64, error)
	GetUserPage(ctx context.Context, userId string, b *Pagination) ([]*User, error)
	GetWithdrawPage(ctx context.Context, userId uint64, b *Pagination) ([]*Withdraw, error)
	GetEthUserRecordLast(ctx context.Context) (int64, error)
	GetEthUserRecordLastTwo(ctx context.Context) (int64, error)
	GetRecordPageTwo(ctx context.Context, address string, b *Pagination) ([]*EthRecord, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetRecordPage(ctx context.Context, address string, b *Pagination) ([]*EthRecord, error)
	GetUserRecommendByUserId(ctx context.Context, userId uint64) (*UserRecommend, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateStakeGet(ctx context.Context, sg *StakeGet) error
	CreateStakeGit(ctx context.Context, sg *StakeGit) error
	CreateUserRecommend(ctx context.Context, user *User, recommendUser *UserRecommend) (*UserRecommend, error)
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetStakeGitByUserId(ctx context.Context, userId uint64) (*StakeGit, error)
	GetBoxRecord(ctx context.Context, num uint64) ([]*BoxRecord, error)
	GetBoxRecordCount(ctx context.Context, num uint64) (int64, error)
	GetUserBoxRecord(ctx context.Context, userId, num uint64, b *Pagination) ([]*BoxRecord, error)
	GetUserBoxRecordOpen(ctx context.Context, userId, num uint64, open bool, b *Pagination) ([]*BoxRecord, error)
	GetStakeGetTotal(ctx context.Context) (*StakeGetTotal, error)
	GetUserStakeGet(ctx context.Context, userId uint64) (*StakeGet, error)
	GetTotalStakeRate(ctx context.Context) (float64, error)
	GetUserRecommendCount(ctx context.Context, code string) (int64, error)
	GetUserRecommendByCodePage(ctx context.Context, code string, b *Pagination) ([]*UserRecommend, error)
	GetLandByUserIDUsing(ctx context.Context, userID uint64, status []uint64) ([]*Land, error)
	GetLandByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Land, error)
	GetLandByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Land, error)
	GetUserRewardPage(ctx context.Context, userId uint64, reason []uint64, b *Pagination) ([]*Reward, error)
	GetUserRewardPageCount(ctx context.Context, userId uint64, reason []uint64) (int64, error)
	GetSeedByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Seed, error)
	GetSeedByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Seed, error)
	GetLandUserUseByUserIDUseing(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*LandUserUse, error)
	GetExchangeRecordsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*ExchangeRecord, error)
	GetLandUserUseByID(ctx context.Context, id uint64) (*LandUserUse, error)
	GetMarketRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*Market, error)
	GetNoticesByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*Notice, error)
	GetNoticesCountByUserID(ctx context.Context, userID uint64) (int64, error)
	GetPropsByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Prop, error)
	GetPropsByUserIDPropType(ctx context.Context, userID uint64, propType []uint64) ([]*Prop, error)
	GetPropsByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Prop, error)
	GetStakeGetsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*StakeGet, error)
	GetStakeGetPlayRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*StakeGetPlayRecord, error)
	GetStakeGetPlayRecordCount(ctx context.Context, userID uint64, status uint64) (int64, error)
	GetStakeGetRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*StakeGetRecord, error)
	GetStakeGitByUserID(ctx context.Context, userID int64) (*StakeGit, error)
	GetStakeGitRecordsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*StakeGitRecord, error)
	GetStakeGitRecordsByID(ctx context.Context, id, userId uint64) (*StakeGitRecord, error)
	GetWithdrawRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*Withdraw, error)
	GetUserOrderCount(ctx context.Context) (int64, error)
	GetUserOrder(ctx context.Context, b *Pagination) ([]*User, error)
	GetLandUserUseByLandIDsMapUsing(ctx context.Context, userId uint64, landIDs []uint64) (map[uint64]*LandUserUse, error)
	BuyBox(ctx context.Context, giw float64, originValue, value string, uc *BoxRecord) (uint64, error)
	GetUserBoxRecordById(ctx context.Context, id uint64) (*BoxRecord, error)
	OpenBoxSeed(ctx context.Context, id uint64, content string, seedInfo *Seed) (uint64, error)
	OpenBoxProp(ctx context.Context, id uint64, content string, propInfo *Prop) (uint64, error)
	GetAllSeedInfo(ctx context.Context) ([]*SeedInfo, error)
	GetAllPropInfo(ctx context.Context) ([]*PropInfo, error)
	GetAllRandomSeeds(ctx context.Context) ([]*RandomSeed, error)
	UpdateSeedValue(ctx context.Context, scene uint64, newSeed uint64) error
	GetSeedByID(ctx context.Context, seedID, userId, status uint64) (*Seed, error)
	GetLandByID(ctx context.Context, landID uint64) (*Land, error)
	GetLandByIDTwo(ctx context.Context, landID uint64) (*Land, error)
	GetLandByUserIdLocationNum(ctx context.Context, userId uint64, locationNum uint64) (*Land, error)
	Plant(ctx context.Context, status, originStatus, perHealth uint64, landUserUse *LandUserUse) error
	PlantPlatTwo(ctx context.Context, id, landId uint64, rent bool) error
	PlantPlatThree(ctx context.Context, id, overTime, propId uint64, one, two bool) error
	PlantPlatFour(ctx context.Context, outMax float64, id, propId, propStatus, propNum uint64) error
	PlantPlatFive(ctx context.Context, overTime, id, propId, propStatus, propNum uint64) error
	PlantPlatSix(ctx context.Context, id, propId, propStatus, propNum, landId uint64) error
	PlantPlatSeven(ctx context.Context, outMax, amount float64, subTime, lastTime, id, propId, propStatus, propNum, userId uint64) error
	PlantPlatTwoTwo(ctx context.Context, id, userId, rentUserId uint64, amount, rentAmount float64) error
	PlantPlatTwoTwoL(ctx context.Context, id, userId, lowUserId, num uint64, amount float64) error
	GetSeedBuyByID(ctx context.Context, seedID, status uint64) (*Seed, error)
	GetPropByID(ctx context.Context, propID, status uint64) (*Prop, error)
	GetPropByIDTwo(ctx context.Context, propID uint64) (*Prop, error)
	BuySeed(ctx context.Context, git, getGit float64, userId, userIdGet, seedId uint64) error
	BuyLand(ctx context.Context, git, getGit float64, userId, userIdGet, landId uint64) error
	BuyProp(ctx context.Context, git, getGit float64, userId, userIdGet, propId uint64) error
	SellLand(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error
	SellProp(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error
	SellSeed(ctx context.Context, seedId uint64, userId uint64, sellAmount float64) error
	GetPropByIDSell(ctx context.Context, propID, status uint64) (*Prop, error)
	UnSellLand(ctx context.Context, propId uint64, userId uint64) error
	UnSellProp(ctx context.Context, propId uint64, userId uint64) error
	UnSellSeed(ctx context.Context, seedId, userId uint64) error
	RentLand(ctx context.Context, landId uint64, userId uint64, rentRate float64) error
	UnRentLand(ctx context.Context, landId uint64, userId uint64) error
	LandPull(ctx context.Context, landId uint64, userId uint64) error
	LandPush(ctx context.Context, landId uint64, userId, locationNum uint64) error
	LandAddOutRate(ctx context.Context, id, landId, userId uint64) error
	CreateLand(ctx context.Context, lc *Land) (*Land, error)
	GetLand(ctx context.Context, id, id2, userId uint64) error
	GetLandInfoByLevels(ctx context.Context) (map[uint64]*LandInfo, error)
	GetLandInfo(ctx context.Context) ([]*LandInfo, error)
	SetGiw(ctx context.Context, address string, giw uint64) error
	SetGit(ctx context.Context, address string, git uint64) error
	SetUsdt(ctx context.Context, address string, usdt uint64) error
	SetStakeGetTotal(ctx context.Context, amount, balance float64) error
	SetStakeGetTotalSub(ctx context.Context, amount, balance float64) error
	SetStakeGet(ctx context.Context, userId uint64, git, amount float64) error
	SetStakeGetSub(ctx context.Context, userId uint64, git, amount float64) error
	SetStakeGetPlaySub(ctx context.Context, userId uint64, amount float64) error
	SetStakeGetPlay(ctx context.Context, userId uint64, git, amount float64) error
	SetStakeGit(ctx context.Context, userId uint64, amount float64) error
	SetUnStakeGit(ctx context.Context, id, userId uint64, amount float64) error
	Exchange(ctx context.Context, userId uint64, git, giw float64) error
	Withdraw(ctx context.Context, userId uint64, giw float64) error
	GetAllBuyLandRecords(ctx context.Context, id uint64) ([]*BuyLandRecord, error)
	GetBuyLandById(ctx context.Context) (*BuyLand, error)
	CreateBuyLandRecord(ctx context.Context, limit uint64, bl *BuyLandRecord) error
	GetAdminByAccount(ctx context.Context, account string, password string) (*Admin, error)
	CreateEth(ctx context.Context, e *EthRecord) error
	CreateEthTwo(ctx context.Context, e *EthRecord) error
	AddGiw(ctx context.Context, address string, giw uint64) error
	AddUsdt(ctx context.Context, address string, usdt uint64) error
	AddUserTotal(ctx context.Context, userId, num uint64, giw uint64) error
	RewardProp(ctx context.Context, typeProp int, userId uint64, lastRewardTotal float64) error
	SetAdminPropConfig(ctx context.Context, info *PropInfo) error
	SetAdminSeedConfig(ctx context.Context, info *SeedInfo) error
	SetAdminLandConfig(ctx context.Context, info *LandInfo) error
	UpdateConfig(ctx context.Context, id uint64, value string) error
	GetStakeGitRecords(ctx context.Context) ([]*StakeGitRecord, error)
	DailyReward(ctx context.Context, id, userId uint64, amount float64) error
	DailyRewardL(ctx context.Context, id, userId, lowUserId, num uint64, amount float64) error
	CreateNotice(ctx context.Context, userId uint64, content string, contentTwo string) error
	SetSeed(ctx context.Context, seedInfo *Seed) (uint64, error)
	SetProp(ctx context.Context, propInfo *Prop) (uint64, error)
	SetBuyLand(ctx context.Context, buyLand *BuyLand) error
	GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error)
	UpdateWithdraw(ctx context.Context, id uint64, status string) error
	UpdateUserRewardNew(ctx context.Context, userId uint64, giw, usdt2, usdt float64, amountOrigin float64, stop bool) error
	UpdateUserMyTotalAmountSub(ctx context.Context, userId int64, amount float64) error
	UpdateUserRewardArea(ctx context.Context, userId uint64, giw, usdt2, usdt float64, amountOrigin float64, stop bool, level bool, currentLevel, i uint64, address string) error
	UpdateUserRewardAreaTwo(ctx context.Context, userId uint64, giw, usdt2, usdt float64, amountOrigin float64, stop bool, i uint64, address string) error
	GetRewardYes(ctx context.Context) ([]*RewardTwo, error)
	UpdateUserRewardNewThree(ctx context.Context, userId uint64, giw, usdt2, usdt float64, amountOrigin float64, level uint64, stop bool) error
	GetUserRewardTwoPage(ctx context.Context, userId uint64, reason uint64, b *Pagination) ([]*RewardTwo, error)
	GetUserRewardTwoPageCount(ctx context.Context, userId uint64, reason uint64) (int64, error)
	GetUserRewardAdminPageCount(ctx context.Context, userId uint64, reason uint64) (int64, error)
	GetUserRewardAdminPage(ctx context.Context, userId uint64, reason uint64, b *Pagination) ([]*Reward, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
}

// AppUsecase is an app usecase.
type AppUsecase struct {
	userRepo UserRepo
	tx       Transaction
	log      *log.Helper
}

// NewAppUsecase new a app usecase.
func NewAppUsecase(userRepo UserRepo, tx Transaction, logger log.Logger) *AppUsecase {
	return &AppUsecase{userRepo: userRepo, tx: tx, log: log.NewHelper(logger)}
}

func RandomBoolCrypto() (bool, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return false, err
	}
	return n.Int64() < 5, nil
}

func (ac *AppUsecase) GetExistUserByAddressOrCreate(ctx context.Context, address string, req *pb.EthAuthorizeRequest) (*User, error, string) {
	var (
		user            *User
		rURecommendUser *UserRecommend
		err             error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			if 1 >= len(code) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1"), "无效的推荐码"
			}

			var (
				rUser *User
			)

			rUser, err = ac.userRepo.GetUserByAddress(ctx, code)
			if nil == rUser || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1"), "无效的推荐码"
			}

			//if 0 >= rUser.GiwAdd {
			//	return nil, errors.New(500, "USER_ERROR", "推荐人未入金"), "推荐人未入金"
			//}

			// 查询推荐人的相关信息
			rURecommendUser, err = ac.userRepo.GetUserRecommendByUserId(ctx, rUser.ID)
			if nil == rURecommendUser || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码3"), "无效的推荐码3"
			}
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = ac.userRepo.CreateUser(ctx, &User{
				Address: address,
			})
			if err != nil {
				return err
			}

			_, err = ac.userRepo.CreateUserRecommend(ctx, user, rURecommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			err = ac.userRepo.CreateStakeGet(ctx, &StakeGet{
				UserId: user.ID,
			})
			if err != nil {
				return err
			}

			//err = ac.userRepo.CreateStakeGit(ctx, &StakeGit{
			//	UserId: user.ID,
			//})
			//if err != nil {
			//	return err
			//}

			return nil
		}); err != nil {
			return nil, err, "错误"
		}
	}

	return user, nil, ""
}

func (ac *AppUsecase) UserInfo(ctx context.Context, address string) (*pb.UserInfoReply, error) {
	var (
		user            *User
		boxNum          uint64
		boxSellNum      uint64
		configs         []*Config
		bPrice          float64
		exchangeFeeRate float64
		rewardStakeRate float64
		boxMax          uint64
		boxAmount       float64
		boxStart        string
		boxEnd          string
		stakeOverRate   float64
		sellFeeRate     float64
		withdrawMin     uint64
		withdrawMax     uint64
		err             error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserInfoReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
		"b_price",
		"exchange_fee_rate",
		"reward_stake_rate",
		"box_max",
		"box_sell_num",
		"box_start",
		"box_end",
		"box_amount",
		"stake_over_rate",
		"sell_fee_rate",
		"withdraw_amount_min",
		"withdraw_amount_max",
	)
	if nil != err || nil == configs {
		return &pb.UserInfoReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "box_sell_num" == vConfig.KeyName {
			boxSellNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_min" == vConfig.KeyName {
			withdrawMin, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_max" == vConfig.KeyName {
			withdrawMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "b_price" == vConfig.KeyName {
			bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "exchange_fee_rate" == vConfig.KeyName {
			exchangeFeeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "reward_stake_rate" == vConfig.KeyName {
			rewardStakeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "box_start" == vConfig.KeyName {
			boxStart = vConfig.Value
		}
		if "box_end" == vConfig.KeyName {
			boxEnd = vConfig.Value
		}
		if "box_amount" == vConfig.KeyName {
			boxAmount, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "box_max" == vConfig.KeyName {
			boxMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "stake_over_rate" == vConfig.KeyName {
			stakeOverRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "sell_fee_rate" == vConfig.KeyName {
			sellFeeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	// 推荐
	var (
		userRecommend   *UserRecommend
		myUserRecommend []*UserRecommend
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.UserInfoReply{
			Status: "推荐错误查询",
		}, nil
	}

	myUserRecommend, err = ac.userRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil == myUserRecommend || nil != err {
		return &pb.UserInfoReply{
			Status: "推荐错误查询",
		}, nil
	}

	RecommendTotalRewardOne := user.RewardOne + user.RewardTwo + user.RewardThree
	RecommendTotalRewardTwo := user.RewardTwoOne + user.RewardTwoTwo + user.RewardTwoThree
	RecommendTotalRewardThree := user.RewardThreeOne + user.RewardThreeTwo + user.RewardThreeThree
	RecommendTotalReward := RecommendTotalRewardOne + RecommendTotalRewardTwo + RecommendTotalRewardThree

	var (
		stakeGitRecord []*StakeGitRecord
	)
	stakeGitRecord, err = ac.userRepo.GetStakeGitRecordsByUserID(ctx, user.ID, nil)
	if nil != err {
		return &pb.UserInfoReply{
			Status: "粮仓错误查询",
		}, nil
	}
	stakeGitAmount := float64(0)
	stakeGitAmountToday := float64(0)

	// 获取中国时区（Asia/Shanghai）
	locShanghai, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return &pb.UserInfoReply{
			Status: "时区查询",
		}, nil
	}
	// 获取当前时间（假设服务器时间是 UTC）
	nowUTC := time.Now()
	// 转换当前时间为中国时区
	nowInShanghai := nowUTC.In(locShanghai)
	for _, v := range stakeGitRecord {
		stakeGitAmount += v.Amount
		// 转换用户注册时间到中国时区
		userRegisterInShanghai := v.CreatedAt.In(locShanghai)
		// 计算用户注册当天在中国时区的 24:00（即第二天 00:00:00）
		nextMidnight := time.Date(userRegisterInShanghai.Year(), userRegisterInShanghai.Month(), userRegisterInShanghai.Day()+1, 0, 0, 0, 0, locShanghai)
		// 判断是否超过注册当天的 24:00
		if nowInShanghai.After(nextMidnight) {
			stakeGitAmountToday += v.Amount
		}
	}

	todayStakeGitAmount := stakeGitAmountToday * rewardStakeRate
	if boxNum > 0 {
		//var (
		//	countBox int64
		//)
		//countBox, err = ac.userRepo.GetBoxRecordCount(ctx, boxNum)
		//if nil != err {
		//	return &pb.UserInfoReply{
		//		Status: "盲盒错误查询",
		//	}, nil
		//}
		//
		//boxSellNum = uint64(countBox)
	}

	var (
		stakeGetTotalMy = float64(0)
		stakeGet        *StakeGet

		stakeGetTotalAmount float64
		stakeGetTotal       *StakeGetTotal
	)

	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil != err || nil == stakeGetTotal {
		return &pb.UserInfoReply{
			Status: "我的放大器错误查询",
		}, nil
	}
	stakeGetTotalAmount = stakeGetTotal.Balance

	stakeGet, err = ac.userRepo.GetUserStakeGet(ctx, user.ID)
	if nil != err {
		return &pb.UserInfoReply{
			Status: "我的放大器错误查询",
		}, nil
	}
	if nil != stakeGet {
		if 0 < stakeGetTotal.Amount {
			// 每份价值
			valuePerShare := stakeGetTotalAmount / stakeGetTotal.Amount
			// 用户最大可提取金额
			stakeGetTotalMy = stakeGet.StakeRate * valuePerShare
		}
	}

	return &pb.UserInfoReply{
		Status:                    "ok",
		MyAddress:                 user.Address,
		Level:                     user.Level,
		Giw:                       user.Giw,
		Git:                       user.Git,
		RecommendTotal:            uint64(len(myUserRecommend)),
		RecommendTotalBiw:         user.Total,
		RecommendTotalReward:      RecommendTotalReward,
		RecommendTotalBiwOne:      user.TotalOne,
		RecommendTotalRewardOne:   RecommendTotalRewardOne,
		RecommendTotalBiwTwo:      user.TotalTwo,
		RecommendTotalRewardTwo:   RecommendTotalRewardTwo,
		RecommendTotalBiwThree:    user.TotalThree,
		RecommendTotalRewardThree: RecommendTotalRewardThree,
		MyStakeGit:                stakeGitAmount,
		TodayRewardSkateGit:       todayStakeGitAmount,
		RewardStakeRate:           rewardStakeRate,
		Box:                       boxMax,
		BoxSell:                   boxSellNum,
		Start:                     boxStart,
		End:                       boxEnd,
		BoxSellAmount:             boxAmount,
		ExchangeRate:              bPrice,
		ExchangeFeeRate:           exchangeFeeRate,
		StakeGetTotal:             stakeGetTotalAmount,
		MyStakeGetTotal:           stakeGetTotalMy,
		StakeGetOverFeeRate:       stakeOverRate,
		SellFeeRate:               sellFeeRate,
		WithdrawMax:               withdrawMax,
		WithdrawMin:               withdrawMin,
	}, nil
}

func (ac *AppUsecase) UserRecommend(ctx context.Context, address string, req *pb.UserRecommendRequest) (*pb.UserRecommendReply, error) {
	res := make([]*pb.UserRecommendReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserRecommendReply{
			Status: "不存在用户",
		}, nil
	}

	// 推荐
	var (
		userRecommend  *UserRecommend
		userRecommends []*UserRecommend
		count          int64
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐错误查询",
		}, nil
	}

	count, err = ac.userRepo.GetUserRecommendCount(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐错误查询",
		}, nil
	}

	userRecommends, err = ac.userRepo.GetUserRecommendByCodePage(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10), &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range userRecommends {
		userIds = append(userIds, v.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐用户查询错误",
		}, nil
	}

	for _, v := range usersMap {
		res = append(res, &pb.UserRecommendReply_List{
			Address:   v.Address,
			Level:     v.Level,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserRecommendReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserRecommendL(ctx context.Context, address string, req *pb.UserRecommendLRequest) (*pb.UserRecommendLReply, error) {
	res := make([]*pb.UserRecommendLReply_List, 0)
	var (
		user  *User
		err   error
		count int64
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserRecommendLReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		reward []*Reward
	)

	status := []uint64{}
	if 1 == req.Num {
		status = append(status, 4, 5, 6)
	} else if 2 == req.Num {
		status = append(status, 7, 8, 9)
	} else if 3 == req.Num {
		status = append(status, 10, 11, 12)
	} else {
		return &pb.UserRecommendLReply{
			Status: "参数错误",
		}, nil
	}

	count, err = ac.userRepo.GetUserRewardPageCount(ctx, user.ID, status)
	if nil != err {
		return &pb.UserRecommendLReply{
			Status: "不存在数据L，count",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardPage(ctx, user.ID, status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserRecommendLReply{
			Status: "不存在数据L",
		}, nil
	}

	userIds := []uint64{}
	for _, v := range reward {
		userIds = append(userIds, v.One)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserRecommendLReply{
			Status: "不存在数据L,用户",
		}, nil
	}

	for _, v := range reward {
		tmpAddress := ""
		if _, ok := usersMap[v.One]; ok {
			tmpAddress = usersMap[v.One].Address
		}

		res = append(res, &pb.UserRecommendLReply_List{
			Address:   tmpAddress,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserRecommendLReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserLand(ctx context.Context, address string, req *pb.UserLandRequest) (*pb.UserLandReply, error) {
	res := make([]*pb.UserLandReply_List, 0)
	var (
		user  *User
		lands []*Land
		err   error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserLandReply{
			Status: "不存在用户",
		}, nil
	}
	status := []uint64{0, 1, 2, 3, 4, 5, 8}
	lands, err = ac.userRepo.GetLandByUserID(ctx, user.ID, status, nil)
	if nil != err {
		return &pb.UserLandReply{
			Status: "不存在用户",
		}, nil
	}

	for _, v := range lands {
		statusTmp := v.Status
		if 8 == v.Status {
			statusTmp = 3
		}

		res = append(res, &pb.UserLandReply_List{
			Id:        v.ID,
			Level:     v.Level,
			Health:    v.MaxHealth,
			Status:    statusTmp,
			OutRate:   v.OutPutRate * 100,
			PerHealth: v.PerHealth,
			One:       v.One,
			Two:       v.Two,
			Three:     v.Three,
		})
	}

	return &pb.UserLandReply{
		Status: "ok",
		Count:  uint64(len(res)),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserStakeGitStakeList(ctx context.Context, address string, req *pb.UserStakeGitStakeListRequest) (*pb.UserStakeGitStakeListReply, error) {
	res := make([]*pb.UserStakeGitStakeListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserStakeGitStakeListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		stakeGitRecord []*StakeGitRecord
	)
	stakeGitRecord, err = ac.userRepo.GetStakeGitRecordsByUserID(ctx, user.ID, nil)
	if nil != err {
		return &pb.UserStakeGitStakeListReply{
			Status: "粮仓错误查询",
		}, nil
	}

	for _, v := range stakeGitRecord {
		res = append(res, &pb.UserStakeGitStakeListReply_List{
			Id:        v.ID,
			Stake:     v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserStakeGitStakeListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, err
}

func (ac *AppUsecase) UserStakeGitRewardList(ctx context.Context, address string, req *pb.UserStakeGitRewardListRequest) (*pb.UserStakeGitRewardListReply, error) {
	res := make([]*pb.UserStakeGitRewardListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserStakeGitRewardListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		reward []*Reward
		count  int64
	)

	status := []uint64{3}
	count, err = ac.userRepo.GetUserRewardPageCount(ctx, user.ID, status)
	if nil != err {
		return &pb.UserStakeGitRewardListReply{
			Status: "粮仓查询错误",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardPage(ctx, user.ID, status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserStakeGitRewardListReply{
			Status: "粮仓查询错误",
		}, nil
	}

	for _, v := range reward {
		res = append(res, &pb.UserStakeGitRewardListReply_List{
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserStakeGitRewardListReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserBoxList(ctx context.Context, address string, req *pb.UserBoxListRequest) (*pb.UserBoxListReply, error) {
	res := make([]*pb.UserBoxListReply_List, 0)
	var (
		boxNum  uint64
		configs []*Config
		user    *User
		err     error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBoxListReply{
			Status: "不存在用户",
		}, nil
	}

	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
	)
	if nil != err || nil == configs {
		return &pb.UserBoxListReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	if 0 < boxNum {
		var (
			box []*BoxRecord
		)
		box, err = ac.userRepo.GetUserBoxRecordOpen(ctx, user.ID, boxNum, true, &Pagination{
			PageNum:  1,
			PageSize: 20,
		})
		if nil != err {
			return &pb.UserBoxListReply{
				Status: "查询错误",
			}, nil
		}

		for _, v := range box {
			res = append(res, &pb.UserBoxListReply_List{
				Content:   v.Content,
				CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &pb.UserBoxListReply{
		Status: "ok",
		Count:  2,
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserBackList(ctx context.Context, address string, req *pb.UserBackListRequest) (*pb.UserBackListReply, error) {
	res := make([]*pb.UserBackListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBackListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		seed []*Seed
	)
	seedStatus := []uint64{0, 4}
	seed, err = ac.userRepo.GetSeedByUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.UserBackListReply{
			Status: "查询种子错误",
		}, nil
	}

	for _, vSeed := range seed {
		tmpStatus := uint64(1)
		if 4 == vSeed.Status {
			tmpStatus = 4
		}

		res = append(res, &pb.UserBackListReply_List{
			Id:     vSeed.ID,
			Type:   1,
			Num:    vSeed.SeedId,
			UseNum: 0,
			Status: tmpStatus,
			OutMax: vSeed.OutMaxAmount,
			Time:   vSeed.OutOverTime,
			Amount: vSeed.SellAmount,
		})
	}

	var (
		prop []*Prop
	)
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	propStatus := []uint64{1, 2, 4}
	prop, err = ac.userRepo.GetPropsByUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.UserBackListReply{
			Status: "道具错误",
		}, nil
	}

	for _, vProp := range prop {

		useNum := uint64(0)
		if 12 == vProp.PropType {
			useNum = uint64(vProp.ThreeOne) // 水
		} else if 13 == vProp.PropType {
			useNum = uint64(vProp.FiveOne) // 手套
		} else if 14 == vProp.PropType {
			useNum = uint64(vProp.FourOne) // 除虫剂
		} else if 15 == vProp.PropType {
			useNum = uint64(vProp.TwoOne) // 铲子
		}

		res = append(res, &pb.UserBackListReply_List{
			Id:     vProp.ID,
			Type:   2,
			Num:    uint64(vProp.PropType),
			UseNum: useNum,
			Status: uint64(vProp.Status),
			OutMax: 0,
			Amount: vProp.SellAmount,
		})
	}

	var (
		box []*BoxRecord
	)

	box, err = ac.userRepo.GetUserBoxRecordOpen(ctx, user.ID, 0, false, nil)
	if nil != err {
		return &pb.UserBackListReply{
			Status: "查询盒子错误",
		}, nil
	}

	for _, v := range box {
		res = append(res, &pb.UserBackListReply_List{
			Id:     v.ID,
			Type:   2,
			Num:    16,
			UseNum: 0,
			Status: 0,
			OutMax: 0,
		})
	}

	return &pb.UserBackListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketSeedList userMarketSeedList.
func (ac *AppUsecase) UserMarketSeedList(ctx context.Context, address string, req *pb.UserMarketSeedListRequest) (*pb.UserMarketSeedListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketSeedListReply{
			Status: "不存在用户",
		}, nil
	}
	res := make([]*pb.UserMarketSeedListReply_List, 0)
	var (
		seed []*Seed
	)

	seedStatus := []uint64{4}
	seed, err = ac.userRepo.GetSeedByExUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.UserMarketSeedListReply{
			Status: "查询错误",
		}, nil
	}

	for _, vSeed := range seed {
		res = append(res, &pb.UserMarketSeedListReply_List{
			Id:     vSeed.ID,
			Num:    vSeed.SeedId,
			Amount: vSeed.SellAmount,
			OutMax: vSeed.OutMaxAmount,
			Time:   vSeed.OutOverTime,
		})
	}

	return &pb.UserMarketSeedListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketLandList userMarketLandList.
func (ac *AppUsecase) UserMarketLandList(ctx context.Context, address string, req *pb.UserMarketLandListRequest) (*pb.UserMarketLandListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketLandListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserMarketLandListReply_List, 0)
	var (
		land []*Land
	)
	landStatus := []uint64{4}
	land, err = ac.userRepo.GetLandByExUserID(ctx, user.ID, landStatus, nil)
	if nil != err {
		return &pb.UserMarketLandListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		res = append(res, &pb.UserMarketLandListReply_List{
			Id:         vLand.ID,
			Level:      vLand.Level,
			MaxHealth:  vLand.MaxHealth,
			Amount:     vLand.SellAmount,
			PerHealth:  vLand.PerHealth,
			OutPutRate: uint64(vLand.OutPutRate) * 100,
		})
	}

	return &pb.UserMarketLandListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketPropList userMarketPropList.
func (ac *AppUsecase) UserMarketPropList(ctx context.Context, address string, req *pb.UserMarketPropListRequest) (*pb.UserMarketPropListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketPropListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserMarketPropListReply_List, 0)
	var (
		prop []*Prop
	)
	propStatus := []uint64{4}
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropsByExUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.UserMarketPropListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range prop {
		useNum := uint64(0)
		if 12 == v.PropType {
			useNum = uint64(v.ThreeOne) // 水
		} else if 13 == v.PropType {
			useNum = uint64(v.FiveOne) // 手套
		} else if 14 == v.PropType {
			useNum = uint64(v.FourOne) // 除虫剂
		} else if 15 == v.PropType {
			useNum = uint64(v.TwoOne) // 铲子
		}

		res = append(res, &pb.UserMarketPropListReply_List{
			Id:     v.ID,
			Num:    uint64(v.PropType),
			Amount: v.SellAmount,
			UseMax: useNum,
		})
	}

	return &pb.UserMarketPropListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketRentLandList userMarketRentLandList.
func (ac *AppUsecase) UserMarketRentLandList(ctx context.Context, address string, req *pb.UserMarketRentLandListRequest) (*pb.UserMarketRentLandListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketRentLandListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserMarketRentLandListReply_List, 0)
	var (
		land []*Land
	)
	landStatus := []uint64{3}
	land, err = ac.userRepo.GetLandByExUserID(ctx, user.ID, landStatus, nil)
	if nil != err {
		return &pb.UserMarketRentLandListReply{
			Status: "错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vLand := range land {
		userIds = append(userIds, vLand.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserMarketRentLandListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		addressTmp := ""
		if _, ok := usersMap[vLand.UserId]; ok {
			addressTmp = usersMap[vLand.UserId].Address
		}

		res = append(res, &pb.UserMarketRentLandListReply_List{
			Id:         vLand.ID,
			Level:      vLand.Level,
			MaxHealth:  vLand.MaxHealth,
			RentAmount: vLand.RentOutPutRate,
			Address:    addressTmp,
			OutPutRate: vLand.OutPutRate,
			PerHealth:  vLand.PerHealth,
		})
	}

	return &pb.UserMarketRentLandListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMyMarketList userMyMarketList.
func (ac *AppUsecase) UserMyMarketList(ctx context.Context, address string, req *pb.UserMyMarketListRequest) (*pb.UserMyMarketListReply, error) {
	res := make([]*pb.UserMyMarketListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMyMarketListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		seed []*Seed
	)
	seedStatus := []uint64{4}
	seed, err = ac.userRepo.GetSeedByUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.UserMyMarketListReply{
			Status: "查询种子错误",
		}, nil
	}

	for _, vSeed := range seed {
		res = append(res, &pb.UserMyMarketListReply_List{
			Id:         vSeed.ID,
			Type:       1,
			Num:        vSeed.SeedId,
			UseNum:     0,
			OutMax:     vSeed.OutMaxAmount,
			Level:      0,
			Status:     0,
			MaxHealth:  0,
			Amount:     vSeed.SellAmount,
			RentAmount: 0,
			Time:       vSeed.OutOverTime,
		})
	}

	var (
		prop []*Prop
	)
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	propStatus := []uint64{4}
	prop, err = ac.userRepo.GetPropsByUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.UserMyMarketListReply{
			Status: "道具错误",
		}, nil
	}

	for _, vProp := range prop {

		useNum := uint64(0)
		if 12 == vProp.PropType {
			useNum = uint64(vProp.ThreeOne) // 水
		} else if 13 == vProp.PropType {
			useNum = uint64(vProp.FiveOne) // 手套
		} else if 14 == vProp.PropType {
			useNum = uint64(vProp.FourOne) // 除虫剂
		} else if 15 == vProp.PropType {
			useNum = uint64(vProp.TwoOne) // 铲子
		}

		res = append(res, &pb.UserMyMarketListReply_List{
			Id:         vProp.ID,
			Type:       2,
			Num:        uint64(vProp.PropType),
			UseNum:     useNum,
			OutMax:     0,
			Level:      0,
			Status:     0,
			MaxHealth:  0,
			Amount:     vProp.SellAmount,
			RentAmount: 0,
		})
	}

	var (
		land []*Land
	)
	landStatus := []uint64{3, 4, 8}
	land, err = ac.userRepo.GetLandByUserID(ctx, user.ID, landStatus, nil)
	if nil != err {
		return &pb.UserMyMarketListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		statusTmp := uint64(1)
		if 4 == vLand.Status {
			statusTmp = 2
		}

		res = append(res, &pb.UserMyMarketListReply_List{
			Id:         vLand.ID,
			Type:       3,
			Num:        0,
			UseNum:     0,
			OutMax:     0,
			Level:      vLand.Level,
			Status:     statusTmp,
			MaxHealth:  vLand.MaxHealth,
			Amount:     vLand.SellAmount,
			RentAmount: vLand.RentOutPutRate,
			PerHealth:  vLand.PerHealth,
			OutPutRate: uint64(vLand.OutPutRate) * 100,
		})
	}

	return &pb.UserMyMarketListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserNoticeList NoticeList.
func (ac *AppUsecase) UserNoticeList(ctx context.Context, address string, req *pb.UserNoticeListRequest) (*pb.UserNoticeListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserNoticeListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserNoticeListReply_List, 0)

	var (
		notice []*Notice
		count  int64
	)

	count, err = ac.userRepo.GetNoticesCountByUserID(ctx, user.ID)
	if nil != err {
		return &pb.UserNoticeListReply{
			Status: "推荐错误查询",
		}, nil
	}

	notice, err = ac.userRepo.GetNoticesByUserID(ctx, user.ID, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserNoticeListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vNotice := range notice {
		res = append(res, &pb.UserNoticeListReply_List{
			Content:   vNotice.NoticeContent,
			CreatedAt: vNotice.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserNoticeListReply{
		Count:  uint64(count),
		List:   res,
		Status: "ok",
	}, nil
}

// UserStakeRewardList userStakeRewardList.
func (ac *AppUsecase) UserStakeRewardList(ctx context.Context, address string, req *pb.UserStakeRewardListRequest) (*pb.UserStakeRewardListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserStakeRewardListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserStakeRewardListReply_List, 0)

	var (
		stakeGetPlayRecord []*StakeGetPlayRecord
		count              int64
	)

	count, err = ac.userRepo.GetStakeGetPlayRecordCount(ctx, user.ID, 0)
	if nil != err {
		return &pb.UserStakeRewardListReply{
			Status: "推荐错误查询",
		}, nil
	}

	stakeGetPlayRecord, err = ac.userRepo.GetStakeGetPlayRecordsByUserID(ctx, user.ID, 0, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserStakeRewardListReply{
			Status: "错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range stakeGetPlayRecord {
		userIds = append(userIds, v.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserStakeRewardListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range stakeGetPlayRecord {
		addressTmp := ""
		if _, ok := usersMap[v.UserId]; ok {
			addressTmp = usersMap[v.UserId].Address
		}

		res = append(res, &pb.UserStakeRewardListReply_List{
			Address: addressTmp,
			Content: "",
			Amount:  v.Amount,
			Reward:  v.Reward,
			Status:  uint64(v.Status),
		})
	}

	return &pb.UserStakeRewardListReply{
		Count:  uint64(count),
		List:   res,
		Status: "ok",
	}, nil
}

// UserIndexList UserIndexList.
func (ac *AppUsecase) UserIndexList(ctx context.Context, address string, req *pb.UserIndexListRequest) (*pb.UserIndexListReply, error) {
	res := make([]*pb.UserIndexListReply_List, 0)
	var (
		user  *User
		lands []*Land
		err   error
	)
	if 20 < len(req.Address) {
		address = req.Address
	}

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserIndexListReply{
			Status: "不存在用户",
		}, nil
	}

	status := []uint64{1, 2, 3, 8}
	lands, err = ac.userRepo.GetLandByUserIDUsing(ctx, user.ID, status)
	if nil != err {
		return &pb.UserIndexListReply{
			Status: "错误查询",
		}, nil
	}

	landIds := make([]uint64, 0)
	for _, vLand := range lands {
		landIds = append(landIds, vLand.ID)
	}

	var (
		landUserUse map[uint64]*LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByLandIDsMapUsing(ctx, user.ID, landIds)
	if nil != err {
		return &pb.UserIndexListReply{
			Status: "错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vLand := range landUserUse {
		userIds = append(userIds, vLand.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserIndexListReply{
			Status: "错误查询",
		}, nil
	}

	resTmp := make(map[uint64]*pb.UserIndexListReply_List, 0)
	now := time.Now().Unix()
	for _, vLand := range lands {
		plantUserAddressTmp := ""

		if _, ok := landUserUse[vLand.ID]; ok {
			if _, ok2 := usersMap[landUserUse[vLand.ID].UserId]; ok2 {
				plantUserAddressTmp = usersMap[landUserUse[vLand.ID].UserId].Address
			}

			tmpRewardStatus := uint64(2)
			rewardTmp := float64(0)
			statusTmp := uint64(1)
			if 0 != landUserUse[vLand.ID].One {
				if landUserUse[vLand.ID].One <= uint64(now) {
					statusTmp = 3
				}

			} else if 0 != landUserUse[vLand.ID].Two {
				if landUserUse[vLand.ID].Two <= uint64(now) {
					statusTmp = 2
				}

				// 有虫子但是已经结束
				if landUserUse[vLand.ID].OverTime <= uint64(now) {
					if uint64(now) > landUserUse[vLand.ID].Two {
						tmp := landUserUse[vLand.ID].OutMaxNum * 0.01 * float64(uint64(now)-landUserUse[vLand.ID].Two) / 300

						if tmp >= landUserUse[vLand.ID].OutMaxNum {
							rewardTmp = 0
						} else {
							rewardTmp = landUserUse[vLand.ID].OutMaxNum - tmp
						}
					}

					tmpRewardStatus = 1
				}
			} else {
				if landUserUse[vLand.ID].OverTime <= uint64(now) {
					rewardTmp = landUserUse[vLand.ID].OutMaxNum
					tmpRewardStatus = 1
				}
			}

			resTmp[vLand.LocationNum] = &pb.UserIndexListReply_List{
				LocationNum:      vLand.LocationNum,
				LandId:           vLand.ID,
				LandLevel:        vLand.Level,
				Health:           vLand.MaxHealth,
				OutRate:          vLand.OutPutRate,
				PerHealth:        vLand.PerHealth,
				LandUseId:        landUserUse[vLand.ID].ID,
				SeedId:           landUserUse[vLand.ID].SeedTypeId,
				Start:            landUserUse[vLand.ID].BeginTime,
				End:              landUserUse[vLand.ID].OverTime,
				CurrentTime:      uint64(now),
				Status:           statusTmp,
				Reward:           rewardTmp,
				PlantUserAddress: plantUserAddressTmp,
				RewardStatus:     tmpRewardStatus,
				LandStatus:       vLand.Status,
			}
		} else {
			resTmp[vLand.LocationNum] = &pb.UserIndexListReply_List{
				LocationNum:      vLand.LocationNum,
				LandId:           vLand.ID,
				LandLevel:        vLand.Level,
				Health:           vLand.MaxHealth,
				OutRate:          vLand.OutPutRate,
				PerHealth:        vLand.PerHealth,
				LandStatus:       vLand.Status,
				LandUseId:        0,
				SeedId:           0,
				Start:            0,
				End:              0,
				CurrentTime:      0,
				Status:           0,
				Reward:           0,
				PlantUserAddress: plantUserAddressTmp,
				RewardStatus:     2,
			}
		}
	}

	for i := uint64(1); i <= 9; i++ {
		if _, ok := resTmp[i]; !ok {
			res = append(res, &pb.UserIndexListReply_List{
				LocationNum:      0,
				LandId:           0,
				LandLevel:        0,
				Health:           0,
				OutRate:          0,
				PerHealth:        0,
				LandUseId:        0,
				SeedId:           0,
				Start:            0,
				End:              0,
				CurrentTime:      0,
				Status:           0,
				Reward:           0,
				PlantUserAddress: "",
			})
		} else {
			res = append(res, resTmp[i])
		}
	}

	return &pb.UserIndexListReply{
		Status: "ok",
		Count:  9,
		List:   res,
	}, nil
}

// UserOrderList userOrderList.
func (ac *AppUsecase) UserOrderList(ctx context.Context, address string, req *pb.UserOrderListRequest) (*pb.UserOrderListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserOrderListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		count int64
		users []*User
	)
	count, err = ac.userRepo.GetUserOrderCount(ctx)
	if nil != err {
		return &pb.UserOrderListReply{
			Status: "查询错误",
		}, nil
	}

	users, err = ac.userRepo.GetUserOrder(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserOrderListReply{
			Status: "查询错误",
		}, nil
	}

	res := make([]*pb.UserOrderListReply_List, 0)

	for _, v := range users {
		res = append(res, &pb.UserOrderListReply_List{
			Address: v.Address,
			Git:     v.Git,
		})
	}

	return &pb.UserOrderListReply{
		Count:  uint64(count),
		List:   res,
		Status: "ok",
	}, nil
}

// 随机数生成器的初始化锁
var rngMutexBuyBox sync.Mutex

func (ac *AppUsecase) BuyBox(ctx context.Context, address string, req *pb.BuyBoxRequest) (*pb.BuyBoxReply, error) {
	rngMutexBuyBox.Lock()
	defer rngMutexBuyBox.Unlock()

	var (
		user             *User
		err              error
		boxNum           uint64
		boxSellNum       uint64
		boxSellNumOrigin string
		configs          []*Config
		boxMax           uint64
		boxAmount        float64
		boxStart         string
		boxEnd           string
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyBoxReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
		"box_max",
		"box_sell_num",
		"box_start",
		"box_end",
		"box_amount",
	)
	if nil != err || nil == configs {
		return &pb.BuyBoxReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "box_sell_num" == vConfig.KeyName {
			boxSellNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		boxSellNumOrigin = vConfig.Value
		if "box_start" == vConfig.KeyName {
			boxStart = vConfig.Value
		}
		if "box_end" == vConfig.KeyName {
			boxEnd = vConfig.Value
		}
		if "box_amount" == vConfig.KeyName {
			boxAmount, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "box_max" == vConfig.KeyName {
			boxMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}
	// 解析时间字符串

	var (
		parsedboxStart time.Time
		parsedboxEnd   time.Time
	)
	parsedboxStart, err = time.Parse("2006-01-02 15:04:05", boxStart)
	if err != nil {
		return &pb.BuyBoxReply{
			Status: "解析时间失败",
		}, nil
	}

	parsedboxEnd, err = time.Parse("2006-01-02 15:04:05", boxEnd)
	if err != nil {
		return &pb.BuyBoxReply{
			Status: "解析时间失败",
		}, nil
	}

	// 获取当前时间
	now := time.Now()

	// 比较时间
	if now.After(parsedboxEnd) {
		return &pb.BuyBoxReply{
			Status: "已结束",
		}, nil
	}

	if now.Before(parsedboxStart) {
		return &pb.BuyBoxReply{
			Status: "未开始",
		}, nil
	}

	if boxSellNum >= boxMax {
		return &pb.BuyBoxReply{
			Status: "已售空",
		}, nil
	}

	if boxAmount >= user.Giw {
		return &pb.BuyBoxReply{
			Status: "余额不足",
		}, nil
	}

	tmpSellNumNew := strconv.FormatUint(boxSellNum+1, 10)
	boxId := uint64(0)
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		boxId, err = ac.userRepo.BuyBox(ctx, boxAmount, boxSellNumOrigin, tmpSellNumNew, &BoxRecord{
			UserId: user.ID,
			Num:    boxNum,
		})
		if nil != err {
			return err
		}
		return nil
	}); nil != err {
		fmt.Println(err, "buybox", user)
		return &pb.BuyBoxReply{
			Status: "购买失败",
		}, nil
	}

	return &pb.BuyBoxReply{
		Status: "ok",
		Id:     boxId,
	}, nil
}

// 随机数生成器
var rngBox *rand2.Rand
var rngPlant *rand2.Rand

// 随机数生成器的初始化锁
var rngMutexBox sync.Mutex

func (ac *AppUsecase) OpenBox(ctx context.Context, address string, req *pb.OpenBoxRequest) (*pb.OpenBoxReply, error) {
	rngMutexBox.Lock()
	defer rngMutexBox.Unlock()

	var (
		user *User
		box  *BoxRecord
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.OpenBoxReply{
			Status: "不存在用户",
		}, nil
	}

	box, err = ac.userRepo.GetUserBoxRecordById(ctx, req.SendBody.Id)
	if nil != err || nil == box {
		return &pb.OpenBoxReply{
			Status: "不存在盲盒",
		}, nil
	}

	if user.ID != box.UserId {
		return &pb.OpenBoxReply{
			Status: "非用户盲盒",
		}, nil
	}

	if 0 != box.GoodId {
		return &pb.OpenBoxReply{
			Status: "已开盲盒",
		}, nil
	}

	// 盲盒道具池
	blindBoxItems := make([]struct {
		Name   uint64
		Weight float64
	}, 0)

	var (
		seedInfos    []*SeedInfo
		seedInfosMap map[uint64]*SeedInfo
	)
	seedInfos, err = ac.userRepo.GetAllSeedInfo(ctx)
	if nil != err {
		return &pb.OpenBoxReply{
			Status: "异常配置",
		}, nil
	}

	seedInfosMap = make(map[uint64]*SeedInfo)
	for _, v := range seedInfos {
		seedInfosMap[v.ID] = v

		blindBoxItems = append(blindBoxItems, struct {
			Name   uint64
			Weight float64
		}{Name: v.ID, Weight: v.GetRate})
	}

	var (
		propInfos    []*PropInfo
		propInfosMap map[uint64]*PropInfo
	)
	propInfos, err = ac.userRepo.GetAllPropInfo(ctx)
	if nil != err {
		return &pb.OpenBoxReply{
			Status: "异常配置",
		}, nil
	}

	propInfosMap = make(map[uint64]*PropInfo)
	for _, v := range propInfos {
		propInfosMap[v.PropType] = v

		blindBoxItems = append(blindBoxItems, struct {
			Name   uint64
			Weight float64
		}{Name: v.PropType, Weight: v.GetRate})
	}

	if 0 >= len(blindBoxItems) {
		return &pb.OpenBoxReply{
			Status: "异常配置概率",
		}, nil
	}

	if nil == rngBox {
		var (
			seedInt     int64
			randomSeeds []*RandomSeed
		)
		randomSeeds, err = ac.userRepo.GetAllRandomSeeds(ctx)
		if nil != err {
			return &pb.OpenBoxReply{
				Status: "异常",
			}, nil
		}

		for _, v := range randomSeeds {
			if 1 == v.Scene {
				seedInt = int64(v.SeedValue)
				break
			}
		}

		if 0 >= seedInt {
			seedInt = time.Now().UnixNano()
			err = ac.userRepo.UpdateSeedValue(ctx, 1, uint64(seedInt))
			if nil != err {
				return &pb.OpenBoxReply{
					Status: "异常",
				}, nil
			}
		}

		rngBox = rand2.New(rand2.NewSource(seedInt))
	}

	r := rngBox.Float64() // 生成 0.0 ~ 1.0 之间的随机数
	// 计算总权重
	var totalWeight float64
	for _, item := range blindBoxItems {
		totalWeight += item.Weight
	}

	// 按照概率随机选择
	result := uint64(0)
	var sum float64
	for _, item := range blindBoxItems {
		sum += item.Weight / totalWeight // 归一化
		if r < sum {
			result = item.Name
			break
		}
	}

	if 0 >= result || 15 < result {
		return &pb.OpenBoxReply{
			Status: "错误盲盒",
		}, nil
	}

	if 1 <= result && result <= 10 {
		if _, ok := seedInfosMap[result]; !ok {
			return &pb.OpenBoxReply{
				Status: "不存在盲盒信息",
			}, nil
		}
		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))

		outMin := int64(seedInfosMap[result].OutMinAmount)
		outMax := int64(seedInfosMap[result].OutMaxAmount)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		// 种子
		boxId := uint64(0)
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			boxId, err = ac.userRepo.OpenBoxSeed(ctx, box.ID, "", &Seed{
				UserId:       user.ID,
				SeedId:       result,
				Name:         seedInfosMap[result].Name,
				OutOverTime:  seedInfosMap[result].OutOverTime,
				OutMaxAmount: float64(randomNumber),
			})
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			fmt.Println(err, "openBox", user)
			return &pb.OpenBoxReply{
				Status: "开启失败",
			}, nil
		}

		return &pb.OpenBoxReply{
			Id:       boxId,
			Status:   "ok",
			OpenType: 1,
			Num:      result,
			OutMax:   float64(randomNumber),
			Time:     seedInfosMap[result].OutOverTime,
		}, nil
	} else if 11 <= result && result <= 15 {
		if _, ok := propInfosMap[result]; !ok {
			return &pb.OpenBoxReply{
				Status: "不存在盲盒信息",
			}, nil
		}

		// 种子
		boxId := uint64(0)
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			boxId, err = ac.userRepo.OpenBoxProp(ctx, box.ID, "", &Prop{
				UserId:   user.ID,
				PropType: int(result),
				OneOne:   int(propInfosMap[result].OneOne),
				OneTwo:   int(propInfosMap[result].OneTwo),
				TwoOne:   int(propInfosMap[result].TwoOne),
				TwoTwo:   propInfosMap[result].TwoTwo,
				ThreeOne: int(propInfosMap[result].ThreeOne),
				FourOne:  int(propInfosMap[result].FourOne),
				FiveOne:  int(propInfosMap[result].FiveOne),
			})
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			fmt.Println(err, "openBox", user)
			return &pb.OpenBoxReply{
				Status: "开启失败",
			}, nil
		}

		useNum := uint64(0)
		if 12 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].ThreeOne // 水
		} else if 13 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].FiveOne // 手套
		} else if 14 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].FourOne // 除虫剂
		} else if 15 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].TwoOne // 铲子
		}

		return &pb.OpenBoxReply{
			Id:       boxId,
			Status:   "ok",
			OpenType: 2,
			Num:      result,
			OutMax:   0,
			UseNum:   useNum,
		}, nil
	} else {
		return &pb.OpenBoxReply{
			Status: "开盲盒失败",
		}, nil
	}
}

var rngMutexPlant sync.Mutex

// LandPlayOne 种植
func (ac *AppUsecase) LandPlayOne(ctx context.Context, address string, req *pb.LandPlayOneRequest) (*pb.LandPlayOneReply, error) {
	rngMutexPlant.Lock()
	defer rngMutexPlant.Unlock()

	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayOneReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		seed *Seed
	)
	seed, err = ac.userRepo.GetSeedByID(ctx, req.SendBody.SeedId, user.ID, 0)
	if nil != err || nil == seed {
		return &pb.LandPlayOneReply{
			Status: "不存种子",
		}, nil
	}

	var (
		land *Land
	)
	land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
	if nil != err || nil == land {
		return &pb.LandPlayOneReply{
			Status: "土地信息错误",
		}, nil
	}

	if land.PerHealth > land.MaxHealth {
		return &pb.LandPlayOneReply{
			Status: "肥沃度不足",
		}, nil
	}

	if land.UserId != user.ID {
		if 3 != land.Status {
			return &pb.LandPlayOneReply{
				Status: "未出租土地",
			}, nil
		}
	} else if land.UserId == user.ID {
		if 1 != land.Status {
			return &pb.LandPlayOneReply{
				Status: "已出租土地",
			}, nil
		}
	} else {
		return &pb.LandPlayOneReply{
			Status: "错误参数",
		}, nil
	}

	if 0 == land.LocationNum {
		return &pb.LandPlayOneReply{
			Status: "未布置土地",
		}, nil
	}

	if nil == rngPlant {
		var (
			seedInt     int64
			randomSeeds []*RandomSeed
		)
		randomSeeds, err = ac.userRepo.GetAllRandomSeeds(ctx)
		if nil != err {
			return &pb.LandPlayOneReply{
				Status: "异常",
			}, nil
		}

		for _, v := range randomSeeds {
			if 2 == v.Scene {
				seedInt = int64(v.SeedValue)
				break
			}
		}

		if 0 >= seedInt {
			seedInt = time.Now().UnixNano()
			err = ac.userRepo.UpdateSeedValue(ctx, 2, uint64(seedInt))
			if nil != err {
				return &pb.LandPlayOneReply{
					Status: "异常",
				}, nil
			}
		}

		rngPlant = rand2.New(rand2.NewSource(seedInt))
	}

	now := uint64(time.Now().Unix())
	rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	outMin := int64(now)
	outMax := int64(now + seed.OutOverTime)

	// 计算随机范围
	tmpNum := outMax - outMin
	if tmpNum <= 0 {
		tmpNum = 1 // 避免 Int63n(0) panic
	}
	// 生成随机数
	randomNumber := outMin + rngTmp.Int63n(tmpNum)

	one := uint64(0)
	two := uint64(0)
	r := rngPlant.Float64() // 生成 0.0 ~ 1.0 之间的随机数
	if r < 0.05 {
		one = uint64(randomNumber)
	} else if r < 0.1 {
		two = uint64(randomNumber)
	}

	originStatusTmp := land.Status
	statusTmp := uint64(2)
	if 3 == originStatusTmp {
		statusTmp = 8
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.Plant(ctx, statusTmp, originStatusTmp, land.PerHealth, &LandUserUse{
			LandId:      land.ID,
			Level:       land.Level,
			UserId:      user.ID,
			OwnerUserId: land.UserId,
			SeedId:      seed.ID,
			SeedTypeId:  seed.SeedId,
			Status:      1,
			BeginTime:   now,
			TotalTime:   seed.OutOverTime,
			OverTime:    now + seed.OutOverTime,
			OutMaxNum:   seed.OutMaxAmount * land.OutPutRate,
			One:         one, // 水时间
			Two:         two, // 虫子时间
		})
	}); nil != err {
		fmt.Println(err, "openBox", user)
		return &pb.LandPlayOneReply{
			Status: "种植失败",
		}, nil
	}

	return &pb.LandPlayOneReply{
		Status: "ok",
	}, nil

}

var rngMutexPlantTwo sync.Mutex

// LandPlayTwo 收果实
func (ac *AppUsecase) LandPlayTwo(ctx context.Context, address string, req *pb.LandPlayTwoRequest) (*pb.LandPlayTwoReply, error) {
	rngMutexPlantTwo.Lock()
	defer rngMutexPlantTwo.Unlock()

	var (
		configs   []*Config
		user      *User
		oneRate   float64
		twoRate   float64
		threeRate float64
		err       error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayTwoReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"one_rate", "two_rate", "three_rate",
	)
	if nil != err || nil == configs {
		return &pb.LandPlayTwoReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "one_rate" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "two_rate" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "three_rate" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayTwoReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayTwoReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayTwoReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if uint64(current) < landUserUse.OverTime {
		return &pb.LandPlayTwoReply{
			Status: "种植未结束",
		}, nil
	}

	if 0 != landUserUse.One {
		return &pb.LandPlayTwoReply{
			Status: "停止生长状态",
		}, nil
	}

	// 已结束
	reward := landUserUse.OutMaxNum
	now := time.Now().Unix()
	// 有虫子 todo 杀虫和浇水更新数量和结束时间 偷的时候注意梳理
	if 0 != landUserUse.Two {
		if uint64(now) > landUserUse.Two {
			tmp := landUserUse.OutMaxNum * 0.01 * float64(uint64(now)-landUserUse.Two) / 300

			if tmp >= landUserUse.OutMaxNum {
				reward = 0
			} else {
				reward = landUserUse.OutMaxNum - tmp
			}
		}
	}

	var (
		land *Land
	)
	land, err = ac.userRepo.GetLandByIDTwo(ctx, landUserUse.LandId)
	if nil != err || nil == land {
		return &pb.LandPlayTwoReply{
			Status: "土地信息错误",
		}, nil
	}

	// 租的
	rentReward := float64(0)
	if landUserUse.UserId != landUserUse.OwnerUserId {
		if 0 < reward {
			rentReward = reward * land.RentOutPutRate
			if reward > rentReward {
				reward = reward - rentReward
			}
		}
	}

	// 推荐
	var (
		userRecommend *UserRecommend
	)
	tmpRecommendUserIds := make([]string, 0)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.UserId)
	if nil == userRecommend || nil != err {
		return &pb.LandPlayTwoReply{
			Status: "查询推荐错误",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	// 收租推荐
	tmpRecommendUserIdsRent := make([]string, 0)
	tmpRent := false
	rentUserId := uint64(0)
	if landUserUse.UserId != landUserUse.OwnerUserId {
		tmpRent = true
		rentUserId = landUserUse.OwnerUserId
		var (
			userRecommendRent *UserRecommend
		)
		userRecommendRent, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.OwnerUserId)
		if nil == userRecommendRent || nil != err {
			return &pb.LandPlayTwoReply{
				Status: "查询推荐错误",
			}, nil
		}
		if "" != userRecommendRent.RecommendCode {
			tmpRecommendUserIdsRent = strings.Split(userRecommendRent.RecommendCode, "D")
		}
	}

	// 分红，状态变更
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		// 资源释放
		err = ac.userRepo.PlantPlatTwo(ctx, landUserUse.ID, land.ID, tmpRent)
		if nil != err {
			return err
		}

		// 奖励
		err = ac.userRepo.PlantPlatTwoTwo(ctx, landUserUse.ID, landUserUse.UserId, rentUserId, reward, rentReward)
		if nil != err {
			return err
		}

		// l1-l3，奖励发放
		if reward > 0 {
			tmpI := 1
			for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
				if 4 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}
				tmpReward := float64(0)

				tmpNum := uint64(4)
				tmpReward = reward * oneRate
				if 2 == tmpI {
					tmpReward = reward * twoRate
					tmpNum = 7
				} else if 3 == tmpI {
					tmpReward = reward * threeRate
					tmpNum = 10
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, uint64(tmpUserId), landUserUse.UserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}
			}
		}

		// l1-l3，奖励发放
		if rentReward > 0 {
			tmpI := 1
			for i := len(tmpRecommendUserIdsRent) - 1; i >= 0; i-- {
				if 4 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIdsRent[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}
				tmpReward := float64(0)

				tmpNum := uint64(5)
				tmpReward = rentReward * oneRate
				if 2 == tmpI {
					tmpReward = rentReward * twoRate
					tmpNum = 8
				} else if 3 == tmpI {
					tmpReward = rentReward * threeRate
					tmpNum = 11
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, uint64(tmpUserId), landUserUse.OwnerUserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}
			}
		}

		return nil
	}); nil != err {
		fmt.Println(err, "LandPlayTwo", landUserUse)
		return &pb.LandPlayTwoReply{
			Status: "种植失败",
		}, nil
	}

	return &pb.LandPlayTwoReply{
		Status: "ok",
	}, nil
}

// LandPlayThree 施肥
func (ac *AppUsecase) LandPlayThree(ctx context.Context, address string, req *pb.LandPlayThreeRequest) (*pb.LandPlayThreeReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayThreeReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		prop *Prop
	)
	prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 1)
	if nil != err || nil == prop {
		return &pb.LandPlayThreeReply{
			Status: "不存道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlayThreeReply{
			Status: "不是自己的",
		}, nil
	}

	if 11 != prop.PropType {
		return &pb.LandPlayThreeReply{
			Status: "不是化肥",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayThreeReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayThreeReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayThreeReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if 0 != landUserUse.One && uint64(current) >= landUserUse.One {
		return &pb.LandPlayThreeReply{
			Status: "停止生长状态",
		}, nil
	}

	if 0 != landUserUse.Two && uint64(current) >= landUserUse.Two {
		return &pb.LandPlayThreeReply{
			Status: "生虫状态",
		}, nil
	}

	if uint64(current) >= landUserUse.OverTime {
		return &pb.LandPlayThreeReply{
			Status: "种植已经结束了",
		}, nil
	}

	if landUserUse.OverTime < uint64(prop.OneTwo) {
		return &pb.LandPlayThreeReply{
			Status: "道具配置错误",
		}, nil
	}

	overTime := landUserUse.OverTime - uint64(prop.OneTwo)
	one := false
	if overTime <= landUserUse.One {
		one = true
	}

	two := false
	if overTime <= landUserUse.Two {
		two = true
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatThree(ctx, landUserUse.ID, overTime, prop.ID, one, two)
	}); nil != err {
		fmt.Println(err, "LandPlayThree", user)
		return &pb.LandPlayThreeReply{
			Status: "施肥失败",
		}, nil
	}

	return &pb.LandPlayThreeReply{
		Status: "ok",
	}, nil
}

// LandPlayFour 杀虫
func (ac *AppUsecase) LandPlayFour(ctx context.Context, address string, req *pb.LandPlayFourRequest) (*pb.LandPlayFourReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayFourReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlayFourReply{
			Status: "不存在道具",
		}, nil
	}

	if 14 != prop.PropType {
		return &pb.LandPlayFourReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlayFourReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.FourOne {
		return &pb.LandPlayFourReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlayFourReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayFourReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayFourReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayFourReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if 0 >= landUserUse.Two {
		return &pb.LandPlayFourReply{
			Status: "无需杀虫",
		}, nil
	}

	if uint64(current) < landUserUse.Two {
		return &pb.LandPlayFourReply{
			Status: "无需杀虫",
		}, nil
	}

	// 剩余最大产出
	rewardTmp := float64(0)
	if uint64(current) > landUserUse.Two {
		tmp := landUserUse.OutMaxNum * 0.01 * float64(uint64(uint64(current)-landUserUse.Two)/300)
		if tmp < landUserUse.OutMaxNum {
			rewardTmp = landUserUse.OutMaxNum - tmp
		}
	}

	one := uint64(0)
	if 1 <= prop.FourOne {
		one = uint64(prop.FourOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatFour(ctx, rewardTmp, landUserUse.ID, prop.ID, two, one)
	}); nil != err {
		fmt.Println(err, "LandPlayFour", user)
		return &pb.LandPlayFourReply{
			Status: "杀虫失败",
		}, nil
	}

	return &pb.LandPlayFourReply{
		Status: "ok",
	}, nil
}

// LandPlayFive 浇水
func (ac *AppUsecase) LandPlayFive(ctx context.Context, address string, req *pb.LandPlayFiveRequest) (*pb.LandPlayFiveReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayFiveReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlayFiveReply{
			Status: "不存在道具",
		}, nil
	}

	if 12 != prop.PropType {
		return &pb.LandPlayFiveReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlayFiveReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.ThreeOne {
		return &pb.LandPlayFiveReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlayFiveReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayFiveReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayFiveReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayFiveReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if 0 >= landUserUse.One {
		return &pb.LandPlayFiveReply{
			Status: "无需浇水",
		}, nil
	}

	if uint64(current) < landUserUse.One {
		return &pb.LandPlayFiveReply{
			Status: "无需浇水",
		}, nil
	}

	tmpOverTime := landUserUse.OverTime + uint64(current) - landUserUse.One

	one := uint64(0)
	if 1 <= prop.ThreeOne {
		one = uint64(prop.ThreeOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatFive(ctx, tmpOverTime, landUserUse.ID, prop.ID, two, one)
	}); nil != err {
		fmt.Println(err, "LandPlayFive", user)
		return &pb.LandPlayFiveReply{
			Status: "浇水失败",
		}, nil
	}

	return &pb.LandPlayFiveReply{
		Status: "ok",
	}, nil
}

// LandPlaySix 铲子
func (ac *AppUsecase) LandPlaySix(ctx context.Context, address string, req *pb.LandPlaySixRequest) (*pb.LandPlaySixReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlaySixReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlaySixReply{
			Status: "不存在道具",
		}, nil
	}

	if 15 != prop.PropType {
		return &pb.LandPlaySixReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlaySixReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.TwoOne {
		return &pb.LandPlaySixReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlaySixReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlaySixReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.OwnerUserId != user.ID {
		return &pb.LandPlaySixReply{
			Status: "非土地用户",
		}, nil
	}

	if landUserUse.UserId == user.ID {
		return &pb.LandPlaySixReply{
			Status: "非出租土地",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlaySixReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if uint64(current) < landUserUse.OverTime {
		return &pb.LandPlaySixReply{
			Status: "还未成熟",
		}, nil
	}

	one := uint64(0)
	if 1 <= prop.TwoOne {
		one = uint64(prop.TwoOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	tmpOverMax := float64(0)
	if landUserUse.OutMaxNum > landUserUse.OutMaxNum*prop.TwoTwo {
		tmpOverMax = landUserUse.OutMaxNum - landUserUse.OutMaxNum*prop.TwoTwo
	}
	tmpOverMaxTwo := landUserUse.OutMaxNum * prop.TwoTwo

	if 0 < landUserUse.One {
		tmpOverMax = 0
		tmpOverMaxTwo = 0
	} else if 0 < landUserUse.Two {
		// 剩余最大产出
		rewardTmp := float64(0)
		if uint64(current) > landUserUse.Two {
			tmp := landUserUse.OutMaxNum * 0.01 * float64(uint64(uint64(current)-landUserUse.Two)/300)
			if tmp < landUserUse.OutMaxNum {
				rewardTmp = landUserUse.OutMaxNum - tmp
			}
		}

		if 0 >= rewardTmp {
			tmpOverMax = 0
			tmpOverMaxTwo = 0
		} else {
			if rewardTmp > rewardTmp*prop.TwoTwo {
				tmpOverMax = rewardTmp - rewardTmp*prop.TwoTwo
			}
			tmpOverMaxTwo = rewardTmp * prop.TwoTwo
		}
	}

	// 推荐
	var (
		userRecommend *UserRecommend
	)
	tmpRecommendUserIds := make([]string, 0)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.UserId)
	if nil == userRecommend || nil != err {
		return &pb.LandPlaySixReply{
			Status: "查询推荐错误",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	// 收租推荐
	tmpRecommendUserIdsRent := make([]string, 0)
	if landUserUse.UserId != landUserUse.OwnerUserId {
		var (
			userRecommendRent *UserRecommend
		)
		userRecommendRent, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.OwnerUserId)
		if nil == userRecommendRent || nil != err {
			return &pb.LandPlaySixReply{
				Status: "查询推荐错误",
			}, nil
		}
		if "" != userRecommendRent.RecommendCode {
			tmpRecommendUserIdsRent = strings.Split(userRecommendRent.RecommendCode, "D")
		}
	}

	// 配置
	var (
		configs   []*Config
		oneRate   float64
		twoRate   float64
		threeRate float64
	)
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"one_rate", "two_rate", "three_rate",
	)
	if nil != err || nil == configs {
		return &pb.LandPlaySixReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "one_rate" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "two_rate" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "three_rate" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.PlantPlatSix(ctx, landUserUse.ID, prop.ID, two, one, landUserUse.LandId)
		if nil != err {
			return err
		}

		// 奖励
		err = ac.userRepo.PlantPlatTwoTwo(ctx, landUserUse.ID, landUserUse.UserId, landUserUse.OwnerUserId, tmpOverMax, tmpOverMaxTwo)
		if nil != err {
			return err
		}

		// l1-l3，奖励发放
		if tmpOverMax > 0 {
			tmpI := 1
			for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
				if 4 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}
				tmpReward := float64(0)

				tmpNum := uint64(4)
				tmpReward = tmpOverMax * oneRate
				if 2 == tmpI {
					tmpReward = tmpOverMax * twoRate
					tmpNum = 7
				} else if 3 == tmpI {
					tmpReward = tmpOverMax * threeRate
					tmpNum = 10
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, uint64(tmpUserId), landUserUse.UserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}
			}
		}

		// l1-l3，奖励发放
		if tmpOverMaxTwo > 0 {
			tmpI := 1
			for i := len(tmpRecommendUserIdsRent) - 1; i >= 0; i-- {
				if 4 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIdsRent[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}
				tmpReward := float64(0)

				tmpNum := uint64(5)
				tmpReward = tmpOverMaxTwo * oneRate
				if 2 == tmpI {
					tmpReward = tmpOverMaxTwo * twoRate
					tmpNum = 8
				} else if 3 == tmpI {
					tmpReward = tmpOverMaxTwo * threeRate
					tmpNum = 11
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, uint64(tmpUserId), landUserUse.OwnerUserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}
			}
		}

		return nil
	}); nil != err {
		fmt.Println(err, "LandPlaySix", user)
		return &pb.LandPlaySixReply{
			Status: "铲除土地失败",
		}, nil
	}

	return &pb.LandPlaySixReply{
		Status: "ok",
	}, nil
}

// LandPlaySeven 手套
func (ac *AppUsecase) LandPlaySeven(ctx context.Context, address string, req *pb.LandPlaySevenRequest) (*pb.LandPlaySevenReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlaySevenReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlaySevenReply{
			Status: "不存在道具",
		}, nil
	}

	if 13 != prop.PropType {
		return &pb.LandPlaySevenReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlaySevenReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.FiveOne {
		return &pb.LandPlaySevenReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlaySevenReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlaySevenReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.OwnerUserId == user.ID {
		return &pb.LandPlaySevenReply{
			Status: "土地用户不能使用手套",
		}, nil
	}

	if landUserUse.UserId == user.ID {
		return &pb.LandPlaySevenReply{
			Status: "种植用户不能使用手套",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlaySevenReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if uint64(current) < landUserUse.OverTime {
		return &pb.LandPlaySevenReply{
			Status: "还未成熟",
		}, nil
	}

	if 0 < landUserUse.One {
		return &pb.LandPlaySevenReply{
			Status: "缺水暂停中",
		}, nil
	}

	if 0 < landUserUse.Two {
		return &pb.LandPlaySevenReply{
			Status: "虫蛀减产中",
		}, nil
	}

	lastTime := landUserUse.SubTime
	if 0 < lastTime {
		if uint64(current)-600 <= lastTime {
			return &pb.LandPlaySevenReply{
				Status: "偷盗过于频繁",
			}, nil
		}
	}

	tmpAmount := landUserUse.OutMaxNum * 0.1
	tmpOutMax := float64(0)
	if tmpAmount >= landUserUse.OutMaxNum {
		tmpOutMax = 0
	} else {
		tmpOutMax = landUserUse.OutMaxNum - tmpAmount
	}

	one := uint64(0)
	if 1 <= prop.FiveOne {
		one = uint64(prop.FiveOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatSeven(ctx, tmpOutMax, tmpAmount, uint64(current), lastTime, landUserUse.ID, prop.ID, two, one, user.ID)
	}); nil != err {
		fmt.Println(err, "LandPlaySeven", user)
		return &pb.LandPlaySevenReply{
			Status: "偷取失败",
		}, nil
	}

	return &pb.LandPlaySevenReply{
		Status: "ok",
		Amount: tmpAmount,
	}, nil
}

var rngMutexBuy sync.Mutex

func (ac *AppUsecase) Buy(ctx context.Context, address string, req *pb.BuyRequest) (*pb.BuyReply, error) {
	rngMutexBuy.Lock()
	defer rngMutexBuy.Unlock()

	var (
		user    *User
		feeRate float64
		configs []*Config
		err     error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"sell_fee_rate",
	)
	if nil != err || nil == configs {
		return &pb.BuyReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "sell_fee_rate" == vConfig.KeyName {
			feeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 1 == req.SendBody.BuyType {
		var (
			seed *Seed
		)
		seed, err = ac.userRepo.GetSeedBuyByID(ctx, req.SendBody.Id, 4)
		if nil != err || nil == seed {
			return &pb.BuyReply{
				Status: "不存种子",
			}, nil
		}

		if user.ID == seed.UserId {
			return &pb.BuyReply{
				Status: "不允许购买自己的",
			}, nil
		}
		if 0 >= seed.SellAmount {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if user.Git < seed.SellAmount {
			return &pb.BuyReply{
				Status: "余额不足",
			}, nil
		}
		// 种子
		tmpGet := seed.SellAmount - seed.SellAmount*feeRate
		if 0 >= tmpGet {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.BuySeed(ctx, seed.SellAmount, tmpGet, seed.UserId, user.ID, seed.ID)
		}); nil != err {
			fmt.Println(err, "buySeed", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else if 2 == req.SendBody.BuyType {
		var (
			prop *Prop
		)
		prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 4)
		if nil != err || nil == prop {
			return &pb.BuyReply{
				Status: "不存道具",
			}, nil
		}

		if user.ID == prop.UserId {
			return &pb.BuyReply{
				Status: "不允许购买自己的",
			}, nil
		}

		if 0 >= prop.SellAmount {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if user.Git < prop.SellAmount {
			return &pb.BuyReply{
				Status: "余额不足",
			}, nil
		}

		// 种子
		tmpGet := prop.SellAmount - prop.SellAmount*feeRate
		if 0 >= tmpGet {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.BuyProp(ctx, prop.SellAmount, tmpGet, prop.UserId, user.ID, prop.ID)
		}); nil != err {
			fmt.Println(err, "buyProp", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else if 3 == req.SendBody.BuyType {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
		if nil != err || nil == land {
			return &pb.BuyReply{
				Status: "不存道具",
			}, nil
		}

		if user.ID == land.UserId {
			return &pb.BuyReply{
				Status: "不允许购买自己的",
			}, nil
		}

		if 4 != land.Status {
			return &pb.BuyReply{
				Status: "未出售",
			}, nil
		}

		if 0 >= land.SellAmount {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if user.Git < land.SellAmount {
			return &pb.BuyReply{
				Status: "余额不足",
			}, nil
		}

		// 土地
		tmpGet := land.SellAmount - land.SellAmount*feeRate
		if 0 >= tmpGet {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.BuyLand(ctx, land.SellAmount, tmpGet, land.UserId, user.ID, land.ID)
		}); nil != err {
			fmt.Println(err, "buyLand", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else {
		return &pb.BuyReply{
			Status: "参数错误",
		}, nil
	}

	return &pb.BuyReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) Sell(ctx context.Context, address string, req *pb.SellRequest) (*pb.SellReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.SellReply{
			Status: "不存在用户",
		}, nil
	}

	tmpSellAmount := float64(req.SendBody.Amount)
	if 1 == req.SendBody.Num {
		if 0 >= tmpSellAmount {
			return &pb.SellReply{
				Status: "售价不能为0",
			}, nil
		}

		if 1 == req.SendBody.SellType {
			var (
				seed *Seed
			)
			seed, err = ac.userRepo.GetSeedBuyByID(ctx, req.SendBody.Id, 0)
			if nil != err || nil == seed {
				return &pb.SellReply{
					Status: "不存在种子",
				}, nil
			}

			if user.ID != seed.UserId {
				return &pb.SellReply{
					Status: "不是自己的种子",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.SellSeed(ctx, seed.ID, user.ID, tmpSellAmount)
			}); nil != err {
				fmt.Println(err, "sellSeed", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else if 2 == req.SendBody.SellType {
			var (
				prop *Prop
			)
			prop, err = ac.userRepo.GetPropByIDSell(ctx, req.SendBody.Id, 2)
			if nil != err || nil == prop {
				return &pb.SellReply{
					Status: "不存在道具",
				}, nil
			}

			if user.ID != prop.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.SellProp(ctx, prop.ID, user.ID, tmpSellAmount)
			}); nil != err {
				fmt.Println(err, "sellProp", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else if 3 == req.SendBody.SellType {
			var (
				land *Land
			)
			land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
			if nil != err || nil == land {
				return &pb.SellReply{
					Status: "不存在土地",
				}, nil
			}

			if user.ID != land.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if 0 != land.LocationNum {
				return &pb.SellReply{
					Status: "土地布置中",
				}, nil
			}

			if 0 != land.Status {
				return &pb.SellReply{
					Status: "不符合上架状态",
				}, nil
			}

			if 1 != land.Three {
				return &pb.SellReply{
					Status: "不可出售土地",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.SellLand(ctx, land.ID, user.ID, tmpSellAmount)
			}); nil != err {
				fmt.Println(err, "sellProp", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else {
			return &pb.SellReply{
				Status: "参数错误",
			}, nil
		}
	} else {
		if 1 == req.SendBody.SellType {
			var (
				seed *Seed
			)
			seed, err = ac.userRepo.GetSeedBuyByID(ctx, req.SendBody.Id, 4)
			if nil != err || nil == seed {
				return &pb.SellReply{
					Status: "不存在种子",
				}, nil
			}

			if user.ID != seed.UserId {
				return &pb.SellReply{
					Status: "不是自己的种子",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.UnSellSeed(ctx, seed.ID, user.ID)
			}); nil != err {
				fmt.Println(err, "unSellSeed", user)
				return &pb.SellReply{
					Status: "下架失败",
				}, nil
			}
		} else if 2 == req.SendBody.SellType {
			var (
				prop *Prop
			)
			prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 4)
			if nil != err || nil == prop {
				return &pb.SellReply{
					Status: "不存在道具",
				}, nil
			}

			if user.ID != prop.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.UnSellProp(ctx, prop.ID, user.ID)
			}); nil != err {
				fmt.Println(err, "unSellProp", user)
				return &pb.SellReply{
					Status: "下架失败",
				}, nil
			}
		} else if 3 == req.SendBody.SellType {
			var (
				land *Land
			)
			land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
			if nil != err || nil == land {
				return &pb.SellReply{
					Status: "不存在土地",
				}, nil
			}

			if user.ID != land.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if 0 != land.LocationNum {
				return &pb.SellReply{
					Status: "土地布置中",
				}, nil
			}

			if 4 != land.Status {
				return &pb.SellReply{
					Status: "不符合下架要求",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.UnSellLand(ctx, land.ID, user.ID)
			}); nil != err {
				fmt.Println(err, "unSellLand", user)
				return &pb.SellReply{
					Status: "下架失败",
				}, nil
			}
		} else {
			return &pb.SellReply{
				Status: "参数错误",
			}, nil
		}
	}

	return &pb.SellReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) StakeGit(ctx context.Context, address string, req *pb.StakeGitRequest) (*pb.StakeGitReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGitReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == req.SendBody.Num {
		if 100 > req.SendBody.Amount {
			return &pb.StakeGitReply{
				Status: "git金额要多于100",
			}, nil
		}

		if req.SendBody.Amount > uint64(user.Git) {
			return &pb.StakeGitReply{
				Status: "git余额不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGit(ctx, user.ID, float64(req.SendBody.Amount))
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGitReply{
				Status: "stakeGit失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			record *StakeGitRecord
		)
		record, err = ac.userRepo.GetStakeGitRecordsByID(ctx, req.SendBody.Id, user.ID) // 查询用户
		if nil != err || nil == record {
			return &pb.StakeGitReply{
				Status: "不存在记录",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetUnStakeGit(ctx, record.ID, user.ID, record.Amount)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGitReply{
				Status: "stakeGit失败",
			}, nil
		}
	} else {
		return &pb.StakeGitReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.StakeGitReply{Status: "ok"}, nil
}

func (ac *AppUsecase) RentLand(ctx context.Context, address string, req *pb.RentLandRequest) (*pb.RentLandReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.RentLandReply{
			Status: "不存在用户",
		}, nil
	}

	rentRate := 0.05
	if 1 == req.SendBody.Rate {
		rentRate = 0.05
	} else if 2 == req.SendBody.Rate {
		rentRate = 0.1
	} else if 3 == req.SendBody.Rate {
		rentRate = 0.2
	} else if 4 == req.SendBody.Rate {
		rentRate = 0.3
	} else if 5 == req.SendBody.Rate {
		rentRate = 0.4
	} else if 6 == req.SendBody.Rate {
		rentRate = 0.5
	} else {
		return &pb.RentLandReply{
			Status: "比例错误",
		}, nil
	}

	if 1 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.RentLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.RentLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 1 != land.Status {
			return &pb.RentLandReply{
				Status: "请将土地布置在农场",
			}, nil
		}

		if 1 != land.One {
			return &pb.RentLandReply{
				Status: "不允许出租类型",
			}, nil
		}

		if land.PerHealth > land.MaxHealth {
			return &pb.RentLandReply{
				Status: "肥沃度不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.RentLand(ctx, land.ID, user.ID, rentRate)
		}); nil != err {
			fmt.Println(err, "rendLand", user)
			return &pb.RentLandReply{
				Status: "上架失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.RentLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.RentLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 3 != land.Status {
			return &pb.RentLandReply{
				Status: "土地使用中",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.UnRentLand(ctx, land.ID, user.ID)
		}); nil != err {
			fmt.Println(err, "unRendLand", user)
			return &pb.RentLandReply{
				Status: "下架失败",
			}, nil
		}
	} else {
		return &pb.RentLandReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.RentLandReply{Status: "ok"}, nil
}

func (ac *AppUsecase) LandPlay(ctx context.Context, address string, req *pb.LandPlayRequest) (*pb.LandPlayReply, error) {
	var (
		user *User
		//box  *BoxRecord
		err error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByIDTwo(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if 1 != land.Status {
			return &pb.LandPlayReply{
				Status: "请将土地布置在农场",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.LandPull(ctx, land.ID, user.ID)
		}); nil != err {
			fmt.Println(err, "LandPull", user)
			return &pb.LandPlayReply{
				Status: "上架失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			land  *Land
			land2 *Land
		)
		land, err = ac.userRepo.GetLandByIDTwo(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if 1 <= land.LocationNum && 9 >= land.LocationNum {

		} else {
			return &pb.LandPlayReply{
				Status: "非布置土地",
			}, nil
		}

		land2, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandTwoId)
		if nil != err || nil == land2 {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if 0 != land2.LocationNum {
			return &pb.LandPlayReply{
				Status: "已布置土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if user.ID != land2.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if 1 != land.Status {
			return &pb.LandPlayReply{
				Status: "请将土地布置在农场",
			}, nil
		}

		if 0 != land2.Status {
			return &pb.LandPlayReply{
				Status: "土地使用中",
			}, nil
		}

		if 1 > land.LocationNum || 9 < land.LocationNum {
			return &pb.LandPlayReply{
				Status: "位置参数错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.LandPull(ctx, land.ID, user.ID)
			if nil != err {
				return err
			}

			return ac.userRepo.LandPush(ctx, land2.ID, user.ID, land.LocationNum)
		}); nil != err {
			fmt.Println(err, "LandPullPush", user)
			return &pb.LandPlayReply{
				Status: "替换失败",
			}, nil
		}
	} else if 3 == req.SendBody.Num {
		if 1 > req.SendBody.LocationNum || 9 < req.SendBody.LocationNum {
			return &pb.LandPlayReply{
				Status: "位置参数错误",
			}, nil
		}

		var (
			tmpLand *Land
			land    *Land
		)
		tmpLand, err = ac.userRepo.GetLandByUserIdLocationNum(ctx, user.ID, req.SendBody.LocationNum)
		if nil != err {
			return &pb.LandPlayReply{
				Status: "错误查询",
			}, nil
		}

		if nil != tmpLand {
			return &pb.LandPlayReply{
				Status: "存在布置土地",
			}, nil
		}

		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if 0 != land.LocationNum {
			return &pb.LandPlayReply{
				Status: "不是空闲的土地",
			}, nil
		}

		if 0 != land.Status {
			return &pb.LandPlayReply{
				Status: "不是空闲的土地",
			}, nil
		}

		if land.PerHealth > land.MaxHealth {
			return &pb.LandPlayReply{
				Status: "肥沃度不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.LandPush(ctx, land.ID, user.ID, req.SendBody.LocationNum)
		}); nil != err {
			fmt.Println(err, "LandPush", user)
			return &pb.LandPlayReply{
				Status: "失败",
			}, nil
		}
	} else {
		return &pb.LandPlayReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.LandPlayReply{Status: "ok"}, nil
}

func (ac *AppUsecase) LandAddOutRate(ctx context.Context, address string, req *pb.LandAddOutRateRequest) (*pb.LandAddOutRateReply, error) {
	var (
		user *User
		//box  *BoxRecord
		err error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandAddOutRateReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		land *Land
	)
	land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
	if nil != err || nil == land {
		return &pb.LandAddOutRateReply{
			Status: "不存在土地",
		}, nil
	}

	if user.ID != land.UserId {
		return &pb.LandAddOutRateReply{
			Status: "不是自己的",
		}, nil
	}

	if 1 != land.Status && 0 != land.Status && 3 != land.Status {
		return &pb.LandAddOutRateReply{
			Status: "土地不符合培育条件",
		}, nil
	}

	// 化肥道具
	var (
		prop *Prop
	)
	prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 1)
	if nil != err || nil == prop {
		return &pb.LandAddOutRateReply{
			Status: "不存道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandAddOutRateReply{
			Status: "不是自己的道具",
		}, nil
	}

	if 11 != prop.PropType {
		return &pb.LandAddOutRateReply{
			Status: "不是化肥",
		}, nil
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.LandAddOutRate(ctx, prop.ID, land.ID, user.ID)
	}); nil != err {
		fmt.Println(err, "LandAddOutRate", user)
		return &pb.LandAddOutRateReply{
			Status: "培育失败",
		}, nil
	}

	return &pb.LandAddOutRateReply{Status: "ok"}, nil
}

func (ac *AppUsecase) GetLand(ctx context.Context, address string, req *pb.GetLandRequest) (*pb.GetLandReply, error) {
	var (
		user      *User
		landInfos map[uint64]*LandInfo
		err       error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.GetLandReply{
			Status: "不存在用户",
		}, nil
	}

	landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
	if nil != err {
		return &pb.GetLandReply{
			Status: "信息错误",
		}, nil
	}

	if 0 >= len(landInfos) {
		return &pb.GetLandReply{
			Status: "配置信息错误",
		}, nil
	}

	now := time.Now().Unix()
	if 1 == req.SendBody.Num {

		var (
			prop []*Prop
		)
		// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
		propType := []uint64{11, 17}
		prop, err = ac.userRepo.GetPropsByUserIDPropType(ctx, user.ID, propType)
		if nil != err {
			return &pb.GetLandReply{
				Status: "道具错误",
			}, nil
		}

		one := uint64(0)
		two := make([]uint64, 0)
		for _, vProp := range prop {
			if vProp.PropType == 17 {
				one = vProp.ID
			}
			if vProp.PropType == 11 {
				two = append(two, vProp.ID)
			}
		}

		if 0 >= one {
			return &pb.GetLandReply{
				Status: "缺少地契",
			}, nil
		}

		if 5 > len(two) {
			return &pb.GetLandReply{
				Status: "缺少化肥",
			}, nil
		}

		tmpLevel := uint64(1)
		if _, ok := landInfos[tmpLevel]; !ok {
			return &pb.GetLandReply{
				Status: "不存在级别土地的配置信息",
			}, nil
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			_, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     user.ID,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
				Status:     0,
				One:        1,
				Two:        1,
				Three:      1,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.GetLandReply{
				Status: "培育失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandOneId)
		if nil != err || nil == land {
			return &pb.GetLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.GetLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 0 != land.Status {
			return &pb.GetLandReply{
				Status: "土地不符合合成条件",
			}, nil
		}

		if 1 != land.Two {
			return &pb.GetLandReply{
				Status: "不可合成土地类型",
			}, nil
		}

		var (
			land2 *Land
		)
		land2, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandTwoId)
		if nil != err || nil == land2 {
			return &pb.GetLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land2.UserId {
			return &pb.GetLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 0 != land2.Status {
			return &pb.GetLandReply{
				Status: "土地不符合合成条件",
			}, nil
		}

		if 1 != land2.Two {
			return &pb.GetLandReply{
				Status: "不可合成土地类型",
			}, nil
		}

		if land.Level != land2.Level {
			return &pb.GetLandReply{
				Status: "土地等级不一致",
			}, nil
		}

		if 10 <= land.Level {
			return &pb.GetLandReply{
				Status: "土地已到最高级别",
			}, nil
		}

		tmpLevel := land.Level + 1
		if _, ok := landInfos[tmpLevel]; !ok {
			return &pb.GetLandReply{
				Status: "不存在级别土地的配置信息",
			}, nil
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.GetLand(ctx, land.ID, land2.ID, user.ID)
			if nil != err {
				return err
			}

			_, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     user.ID,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
				Status:     0,
				One:        1,
				Two:        1,
				Three:      1,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.GetLandReply{
				Status: "培育失败",
			}, nil
		}
	} else {
		return &pb.GetLandReply{
			Status: "参数错误",
		}, nil
	}

	return &pb.GetLandReply{Status: "ok"}, nil
}

var stakeAndPlay sync.Mutex

func (ac *AppUsecase) StakeGet(ctx context.Context, address string, req *pb.StakeGetRequest) (*pb.StakeGetReply, error) {
	stakeAndPlay.Lock()
	defer stakeAndPlay.Unlock()

	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGetReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		stakeGetTotal *StakeGetTotal
	)
	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil == stakeGetTotal || nil != err {
		return &pb.StakeGetReply{
			Status: "放大器总额错误查询",
		}, nil
	}

	var (
		stakeGet *StakeGet
	)
	stakeGet, err = ac.userRepo.GetUserStakeGet(ctx, user.ID)
	if nil == stakeGet || nil != err {
		return &pb.StakeGetReply{
			Status: "我的放大器错误查询",
		}, nil
	}

	// 质押
	if 1 == req.SendBody.Num {
		if 100 > req.SendBody.Amount {
			return &pb.StakeGetReply{
				Status: "git金额要多于100",
			}, nil
		}

		if req.SendBody.Amount > uint64(user.Git) {
			return &pb.StakeGetReply{
				Status: "git余额不足",
			}, nil
		}

		var mintedShares float64
		tmpAmount := float64(req.SendBody.Amount)
		if 0 >= stakeGetTotal.Balance || 0 >= stakeGetTotal.Amount {
			mintedShares = tmpAmount
		} else {
			mintedShares = tmpAmount * stakeGetTotal.Amount / stakeGetTotal.Balance
		}
		if 0 >= mintedShares {
			return &pb.StakeGetReply{
				Status: "份额计算不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetTotal(ctx, mintedShares, tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.SetStakeGet(ctx, user.ID, tmpAmount, mintedShares)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetReply{
				Status: "git余额不足",
			}, nil
		}
	} else if 2 == req.SendBody.Num {

		if 0 >= stakeGetTotal.Balance || 0 >= stakeGetTotal.Amount {
			return &pb.StakeGetReply{
				Status: "池子为空",
			}, nil
		}

		if 100 > req.SendBody.Amount {
			return &pb.StakeGetReply{
				Status: "git最小提现100",
			}, nil
		}

		if 0 >= stakeGet.StakeRate {
			return &pb.StakeGetReply{
				Status: "用户无可提git",
			}, nil
		}

		// 每份价值
		valuePerShare := stakeGetTotal.Balance / stakeGetTotal.Amount
		// 用户最大可提取金额
		maxWithdraw := stakeGet.StakeRate * valuePerShare
		if req.SendBody.Amount > uint64(maxWithdraw) {
			return &pb.StakeGetReply{
				Status: "可提git不足",
			}, nil
		}

		sharesToRemove := float64(req.SendBody.Amount) / valuePerShare

		tmpGit := float64(req.SendBody.Amount)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetTotalSub(ctx, sharesToRemove, tmpGit)
			if nil != err {
				return err
			}

			err = ac.userRepo.SetStakeGetSub(ctx, user.ID, tmpGit, sharesToRemove)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetReply{
				Status: "git余额不足",
			}, nil
		}
	} else {
		return &pb.StakeGetReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.StakeGetReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) StakeGetPlay(ctx context.Context, address string, req *pb.StakeGetPlayRequest) (*pb.StakeGetPlayReply, error) {
	stakeAndPlay.Lock()
	defer stakeAndPlay.Unlock()

	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGetPlayReply{
			Status: "不存在用户",
		}, nil
	}

	if req.SendBody.Amount > uint64(user.Git) {
		return &pb.StakeGetPlayReply{
			Status: "git余额不足",
		}, nil
	}

	if 100 > req.SendBody.Amount {
		return &pb.StakeGetPlayReply{
			Status: "最少100",
		}, nil
	}

	var (
		stakeGetTotal *StakeGetTotal
	)
	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil == stakeGetTotal || nil != err {
		return &pb.StakeGetPlayReply{
			Status: "放大器总额错误查询",
		}, nil
	}

	if 0 == uint64(stakeGetTotal.Balance) {
		return &pb.StakeGetPlayReply{
			Status: "资金池不足",
		}, nil
	}

	if uint64(stakeGetTotal.Balance) < req.SendBody.Amount {
		return &pb.StakeGetPlayReply{
			Status: "资金池不足",
		}, nil
	}

	rand2.New(rand2.NewSource(time.Now().UnixNano()))
	outcome := rand2.Intn(2)

	if outcome == 0 { // 赢：需要池子中有足够资金支付奖金
		var (
			configs       []*Config
			stakeOverRate float64
		)

		// 配置
		configs, err = ac.userRepo.GetConfigByKeys(ctx,
			"stake_over_rate",
		)
		if nil != err || nil == configs {
			return &pb.StakeGetPlayReply{
				Status: "配置错误",
			}, nil
		}
		for _, vConfig := range configs {
			if "stake_over_rate" == vConfig.KeyName {
				stakeOverRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}

		tmpGit := float64(req.SendBody.Amount) - float64(req.SendBody.Amount)*stakeOverRate
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetPlay(ctx, user.ID, tmpGit, float64(req.SendBody.Amount))
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetPlayReply{
				Status: "git余额不足",
			}, nil
		}

		return &pb.StakeGetPlayReply{Status: "ok", PlayStatus: 1, Amount: tmpGit}, nil
	} else { // 输：下注金额加入池子
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetPlaySub(ctx, user.ID, float64(req.SendBody.Amount))
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetPlayReply{
				Status: "git余额不足",
			}, nil
		}

		return &pb.StakeGetPlayReply{Status: "ok", PlayStatus: 2}, nil
	}
}

func (ac *AppUsecase) SetGiw(ctx context.Context, req *pb.SetGiwRequest) (*pb.SetGiwReply, error) {
	return &pb.SetGiwReply{Status: "ok"}, ac.userRepo.SetGiw(ctx, req.Address, req.Giw)
}

func (ac *AppUsecase) SetGit(ctx context.Context, req *pb.SetGitRequest) (*pb.SetGitReply, error) {
	return &pb.SetGitReply{Status: "ok"}, ac.userRepo.SetGit(ctx, req.Address, req.Git)
}

func (ac *AppUsecase) AdminSetGiw(ctx context.Context, req *pb.AdminSetGiwRequest) (*pb.AdminSetGiwReply, error) {
	return &pb.AdminSetGiwReply{Status: "ok"}, ac.userRepo.SetGiw(ctx, req.Address, req.Biw)
}

func (ac *AppUsecase) AdminSetGit(ctx context.Context, req *pb.AdminSetGitRequest) (*pb.AdminSetGitReply, error) {
	return &pb.AdminSetGitReply{Status: "ok"}, ac.userRepo.SetGit(ctx, req.Address, req.Giw)
}

func (ac *AppUsecase) AdminSetUsdt(ctx context.Context, req *pb.AdminSetUsdtRequest) (*pb.AdminSetUsdtReply, error) {
	return &pb.AdminSetUsdtReply{Status: "ok"}, ac.userRepo.SetUsdt(ctx, req.Address, req.Usdt)
}

func (ac *AppUsecase) Exchange(ctx context.Context, address string, req *pb.ExchangeRequest) (*pb.ExchangeReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.ExchangeReply{
			Status: "不存在用户",
		}, nil
	}

	if req.SendBody.Amount > uint64(user.Git) {
		return &pb.ExchangeReply{
			Status: "git余额不足",
		}, nil
	}

	if 100 > req.SendBody.Amount {
		return &pb.ExchangeReply{
			Status: "最少100",
		}, nil
	}

	var (
		configs []*Config
		rate    float64
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"exchange_fee_rate",
	)
	if nil != err || nil == configs {
		return &pb.ExchangeReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "exchange_fee_rate" == vConfig.KeyName {
			rate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	giw := float64(req.SendBody.Amount) - float64(req.SendBody.Amount)*rate
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.Exchange(ctx, user.ID, float64(req.SendBody.Amount), giw)
		if nil != err {
			return err
		}
		return nil
	}); nil != err {
		return &pb.ExchangeReply{
			Status: "兑换错误",
		}, nil
	}

	return &pb.ExchangeReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) Withdraw(ctx context.Context, address string, req *pb.WithdrawRequest) (*pb.WithdrawReply, error) {
	var (
		user        *User
		configs     []*Config
		err         error
		withdrawMin uint64
		withdrawMax uint64
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.WithdrawReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"withdraw_amount_min",
		"withdraw_amount_max",
	)
	if nil != err || nil == configs {
		return &pb.WithdrawReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "withdraw_amount_min" == vConfig.KeyName {
			withdrawMin, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_max" == vConfig.KeyName {
			withdrawMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	if req.SendBody.Amount > uint64(user.Giw) {
		return &pb.WithdrawReply{
			Status: "giw余额不足",
		}, nil
	}

	if withdrawMin > req.SendBody.Amount {
		return &pb.WithdrawReply{
			Status: "低于最小值",
		}, nil
	}

	if withdrawMax < req.SendBody.Amount {
		return &pb.WithdrawReply{
			Status: "高于最大值",
		}, nil
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.Withdraw(ctx, user.ID, float64(req.SendBody.Amount))
		if nil != err {
			return err
		}
		return nil
	}); nil != err {
		return &pb.WithdrawReply{
			Status: "兑换错误",
		}, nil
	}

	return &pb.WithdrawReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) SetLand(ctx context.Context, req *pb.SetLandRequest) (*pb.SetLandReply, error) {
	var (
		user *User
		err  error

		landInfos map[uint64]*LandInfo
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, req.Address)
	if nil != err {
		return &pb.SetLandReply{Status: "地址不存在用户"}, nil
	}

	if nil == user {
		return &pb.SetLandReply{Status: "地址不存在用户"}, nil
	}

	landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
	if nil != err {
		return &pb.SetLandReply{
			Status: "信息错误",
		}, nil
	}

	if 0 >= len(landInfos) {
		return &pb.SetLandReply{
			Status: "配置信息错误",
		}, nil
	}

	tmpLevel := req.Level
	if _, ok := landInfos[tmpLevel]; !ok {
		return &pb.SetLandReply{
			Status: "级别错误",
		}, nil
	}

	rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	outMin := int64(landInfos[tmpLevel].OutPutRateMin)
	outMax := int64(landInfos[tmpLevel].OutPutRateMax)

	// 计算随机范围
	tmpNum := outMax - outMin
	if tmpNum <= 0 {
		tmpNum = 1 // 避免 Int63n(0) panic
	}

	// 生成随机数
	randomNumber := outMin + rngTmp.Int63n(tmpNum)

	now := time.Now().Unix()
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = ac.userRepo.CreateLand(ctx, &Land{
			UserId:     user.ID,
			Level:      landInfos[tmpLevel].Level,
			OutPutRate: float64(randomNumber),
			MaxHealth:  landInfos[tmpLevel].MaxHealth,
			PerHealth:  landInfos[tmpLevel].PerHealth,
			LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
			Status:     0,
			One:        1,
			Two:        1,
			Three:      1,
		})
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "setLand", user)
		return &pb.SetLandReply{
			Status: "发放失败",
		}, nil
	}

	return &pb.SetLandReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest, token string) (*pb.AdminLoginReply, error) {
	var (
		admin *Admin
		err   error
	)

	res := &pb.AdminLoginReply{
		Status: "ok",
	}
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	admin, err = ac.userRepo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil == admin || nil != err {
		return res, err
	}

	res.Token = token
	return res, nil
}

func (ac *AppUsecase) AdminUserList(ctx context.Context, req *pb.AdminUserListRequest) (*pb.AdminUserListReply, error) {
	var (
		users []*User
		count int64
		err   error
	)

	useRes := make([]*pb.AdminUserListReply_List, 0)

	count, err = ac.userRepo.GetUserPageCount(ctx, req.Address)
	if nil != err {
		return &pb.AdminUserListReply{
			Status: "错误",
		}, nil
	}

	users, err = ac.userRepo.GetUserPage(ctx, req.Address, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if nil != err {
		return &pb.AdminUserListReply{
			Status: "错误",
		}, nil
	}

	var (
		stakeGetTotal *StakeGetTotal
	)
	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil != err || nil == stakeGetTotal {
		return &pb.AdminUserListReply{
			Status: "错误",
		}, nil
	}

	for _, v := range users {
		// 推荐
		var (
			userRecommend   *UserRecommend
			myUserRecommend []*UserRecommend
		)
		userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, v.ID)
		if nil == userRecommend || nil != err {
			continue
		}

		myUserRecommend, err = ac.userRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(v.ID, 10))
		if nil == myUserRecommend || nil != err {
			continue
		}

		var (
			stakeGitRecord []*StakeGitRecord
		)
		stakeGitRecord, err = ac.userRepo.GetStakeGitRecordsByUserID(ctx, v.ID, nil)
		if nil != err {
			continue
		}
		stakeGitAmount := float64(0)
		for _, vS := range stakeGitRecord {
			stakeGitAmount += vS.Amount
		}

		var (
			stakeGetTotalMy     = float64(0)
			stakeGet            *StakeGet
			stakeGetTotalAmount float64
		)

		stakeGetTotalAmount = stakeGetTotal.Balance
		stakeGet, err = ac.userRepo.GetUserStakeGet(ctx, v.ID)
		if nil != err {
			continue
		}
		if nil != stakeGet {
			if 0 < stakeGetTotal.Amount {
				// 每份价值
				valuePerShare := stakeGetTotalAmount / stakeGetTotal.Amount
				// 用户最大可提取金额
				stakeGetTotalMy = stakeGet.StakeRate * valuePerShare
			}
		}

		useRes = append(useRes, &pb.AdminUserListReply_List{
			UserId:                    v.ID,
			Address:                   v.Address,
			Level:                     v.Level,
			Giw:                       v.Giw,
			Git:                       v.Git,
			RecommendTotal:            uint64(len(myUserRecommend)),
			RecommendTotalBiw:         v.Total,
			RecommendTotalReward:      v.RewardOne + v.RewardTwo + v.RewardThree,
			RecommendTotalBiwOne:      v.TotalOne,
			RecommendTotalRewardOne:   v.RewardOne,
			RecommendTotalBiwTwo:      v.TotalTwo,
			RecommendTotalRewardTwo:   v.RewardTwo,
			RecommendTotalBiwThree:    v.TotalThree,
			RecommendTotalRewardThree: v.RewardThree,
			MyStakeGit:                stakeGitAmount,
			MyStakeGetTotal:           stakeGetTotalMy,
			AmountUsdt:                v.AmountUsdt,
		})
	}

	return &pb.AdminUserListReply{
		Status: "ok",
		Users:  useRes,
		Count:  count,
	}, nil
}

func (ac *AppUsecase) AdminRecommendList(ctx context.Context, req *pb.AdminUserRecommendRequest) (*pb.AdminUserRecommendReply, error) {
	var (
		userRecommends []*UserRecommend
		userRecommend  *UserRecommend
		userIdsMap     map[uint64]uint64
		userIds        []uint64
		users          map[uint64]*User
		err            error
	)

	res := &pb.AdminUserRecommendReply{
		Status: "ok",
		Users:  make([]*pb.AdminUserRecommendReply_List, 0),
	}

	// 地址查询
	if 0 >= req.UserId {
		return res, nil
	}

	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, req.UserId)
	if nil == userRecommend {
		return res, nil
	}

	userRecommends, err = ac.userRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(userRecommend.UserId, 10))
	if nil != err {
		return res, nil
	}

	userIdsMap = make(map[uint64]uint64, 0)
	for _, v := range userRecommends {
		userIdsMap[v.UserId] = v.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return res, nil
	}

	for _, v := range userRecommends {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Users = append(res.Users, &pb.AdminUserRecommendReply_List{
			Address:           users[v.UserId].Address,
			UserId:            v.UserId,
			CreatedAt:         v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			RecommendTotalBiw: users[v.UserId].Total,
		})
	}

	return res, nil
}

func (ac *AppUsecase) AdminWithdrawList(ctx context.Context, req *pb.AdminWithdrawListRequest) (*pb.AdminWithdrawListReply, error) {
	var (
		user         *User
		count        int64
		err          error
		userId       uint64
		withdrawList []*Withdraw
	)

	if 0 < len(req.Address) {
		user, err = ac.userRepo.GetUserByAddress(ctx, req.Address) // 查询用户
		if nil != err || nil == user {
			return &pb.AdminWithdrawListReply{
				Status: "不存在用户",
			}, nil
		}
		userId = user.ID
	}

	withdrawRes := make([]*pb.AdminWithdrawListReply_List, 0)

	count, err = ac.userRepo.GetWithdrawPageCount(ctx, userId)
	if nil != err {
		return &pb.AdminWithdrawListReply{
			Status: "错误",
		}, nil
	}

	withdrawList, err = ac.userRepo.GetWithdrawPage(ctx, userId, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if nil != err {
		return &pb.AdminWithdrawListReply{
			Status: "错误",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range withdrawList {
		userIds = append(userIds, v.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.AdminWithdrawListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range withdrawList {
		addressTmp := ""
		if _, ok := usersMap[v.UserId]; ok {
			addressTmp = usersMap[v.UserId].Address
		}

		withdrawRes = append(withdrawRes, &pb.AdminWithdrawListReply_List{
			Address:   addressTmp,
			Id:        v.ID,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    v.Amount,
			Status:    v.Status,
		})
	}

	return &pb.AdminWithdrawListReply{
		Status:   "ok",
		Withdraw: withdrawRes,
		Count:    count,
	}, nil
}

func (ac *AppUsecase) AdminRecordList(ctx context.Context, req *pb.RecordListRequest) (*pb.RecordListReply, error) {
	var (
		count int64
		err   error
		list  []*EthRecord
	)

	res := make([]*pb.RecordListReply_List, 0)

	if "usdt" == req.Coin {
		count, err = ac.userRepo.GetRecordPageCountTwo(ctx, req.Address)
		if nil != err {
			return &pb.RecordListReply{
				Status: "错误",
			}, nil
		}

		list, err = ac.userRepo.GetRecordPageTwo(ctx, req.Address, &Pagination{
			PageNum:  int(req.Page),
			PageSize: 10,
		})
		if nil != err {
			return &pb.RecordListReply{
				Status: "错误",
			}, nil
		}
	} else {
		count, err = ac.userRepo.GetRecordPageCount(ctx, req.Address)
		if nil != err {
			return &pb.RecordListReply{
				Status: "错误",
			}, nil
		}

		list, err = ac.userRepo.GetRecordPage(ctx, req.Address, &Pagination{
			PageNum:  int(req.Page),
			PageSize: 10,
		})
		if nil != err {
			return &pb.RecordListReply{
				Status: "错误",
			}, nil
		}
	}

	for _, v := range list {
		res = append(res, &pb.RecordListReply_List{
			Address:   v.Address,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    v.Amount,
			Coin:      v.Coin,
		})
	}

	return &pb.RecordListReply{
		Status:     "ok",
		RecordList: res,
		Count:      count,
	}, nil
}

func (ac *AppUsecase) AdminLandConfigSet(ctx context.Context, req *pb.AdminLandConfigSetRequest) (*pb.AdminLandConfigSetReply, error) {
	if 0 >= req.SendBody.Level {
		return &pb.AdminLandConfigSetReply{
			Status: "参数错误",
		}, nil
	}

	err := ac.userRepo.SetAdminLandConfig(ctx, &LandInfo{
		Level:         req.SendBody.Level,
		OutPutRateMax: req.SendBody.OutPutRateMax,
		OutPutRateMin: req.SendBody.OutPutRateMin,
		MaxHealth:     req.SendBody.MaxHealth,
		PerHealth:     req.SendBody.PerHealth,
	})
	if nil != err {
		return &pb.AdminLandConfigSetReply{
			Status: "修改失败",
		}, nil
	}

	return &pb.AdminLandConfigSetReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminLandConfig(ctx context.Context, req *pb.AdminLandConfigRequest) (*pb.AdminLandConfigReply, error) {
	var (
		err       error
		landInfos []*LandInfo
	)

	landInfos, err = ac.userRepo.GetLandInfo(ctx)
	if nil != err {
		return &pb.AdminLandConfigReply{Status: "ok"}, err
	}

	if 0 >= len(landInfos) {
		return &pb.AdminLandConfigReply{Status: "ok"}, nil
	}

	res := make([]*pb.AdminLandConfigReply_List, 0)

	for _, v := range landInfos {
		res = append(res, &pb.AdminLandConfigReply_List{
			CreatedAt:     v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			OutPutRateMax: v.OutPutRateMax,
			OutPutRateMin: v.OutPutRateMin,
			MaxHealth:     v.MaxHealth,
			PerHealth:     v.PerHealth,
			Level:         v.Level,
		})
	}

	return &pb.AdminLandConfigReply{
		Status:     "ok",
		RecordList: res,
		Count:      10,
	}, nil
}

func (ac *AppUsecase) AdminSeedConfigSet(ctx context.Context, req *pb.AdminSeedConfigSetRequest) (*pb.AdminSeedConfigSetReply, error) {
	if 0 >= req.SendBody.SeedId {
		return &pb.AdminSeedConfigSetReply{
			Status: "参数错误",
		}, nil
	}

	err := ac.userRepo.SetAdminSeedConfig(ctx, &SeedInfo{
		ID:           req.SendBody.SeedId,
		OutMinAmount: req.SendBody.OutMinAmount,
		OutMaxAmount: req.SendBody.OutMaxAmount,
		OutOverTime:  req.SendBody.OutOverTime,
		GetRate:      req.SendBody.Rate,
	})

	if nil != err {
		return &pb.AdminSeedConfigSetReply{
			Status: "修改失败",
		}, nil
	}

	return &pb.AdminSeedConfigSetReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminSeedConfig(ctx context.Context, req *pb.AdminSeedConfigRequest) (*pb.AdminSeedConfigReply, error) {
	var (
		err       error
		seedInfos []*SeedInfo
	)

	seedInfos, err = ac.userRepo.GetAllSeedInfo(ctx)
	if nil != err {
		return &pb.AdminSeedConfigReply{Status: "ok"}, err
	}

	if 0 >= len(seedInfos) {
		return &pb.AdminSeedConfigReply{Status: "ok"}, nil
	}

	res := make([]*pb.AdminSeedConfigReply_List, 0)

	for _, v := range seedInfos {
		res = append(res, &pb.AdminSeedConfigReply_List{
			Rate:         v.GetRate,
			CreatedAt:    v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			OutMaxAmount: v.OutMaxAmount,
			OutMinAmount: v.OutMinAmount,
			OutOverTime:  v.OutOverTime,
		})
	}

	return &pb.AdminSeedConfigReply{
		Status:     "ok",
		RecordList: res,
		Count:      10,
	}, nil
}

func (ac *AppUsecase) AdminPropConfigSet(ctx context.Context, req *pb.AdminPropConfigSetRequest) (*pb.AdminPropConfigSetReply, error) {

	tmp := &PropInfo{
		PropType: req.SendBody.PropType,
		OneOne:   0,
		OneTwo:   0,
		TwoOne:   0,
		TwoTwo:   0,
		ThreeOne: 0,
		FourOne:  0,
		FiveOne:  0,
		GetRate:  0,
	}

	if 12 == req.SendBody.PropType {
		tmp.ThreeOne = req.SendBody.Max
	} else if 13 == req.SendBody.PropType {
		tmp.FiveOne = req.SendBody.Max
	} else if 14 == req.SendBody.PropType {
		tmp.FourOne = req.SendBody.Max
	} else if 15 == req.SendBody.PropType {
		tmp.TwoOne = req.SendBody.Max
	} else if 11 == req.SendBody.PropType {

	} else {
		return &pb.AdminPropConfigSetReply{
			Status: "参数错误",
		}, nil
	}

	tmp.GetRate = req.SendBody.Rate
	err := ac.userRepo.SetAdminPropConfig(ctx, tmp)
	if nil != err {
		return &pb.AdminPropConfigSetReply{
			Status: "修改失败",
		}, nil
	}

	return &pb.AdminPropConfigSetReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminPropConfig(ctx context.Context, req *pb.AdminPropConfigRequest) (*pb.AdminPropConfigReply, error) {
	var (
		err       error
		propInfos []*PropInfo
	)

	propInfos, err = ac.userRepo.GetAllPropInfo(ctx)
	if nil != err {
		return &pb.AdminPropConfigReply{Status: "ok"}, err
	}

	if 0 >= len(propInfos) {
		return &pb.AdminPropConfigReply{Status: "ok"}, nil
	}

	res := make([]*pb.AdminPropConfigReply_List, 0)

	for _, v := range propInfos {
		useNum := uint64(1)

		if 12 == v.PropType {
			useNum = v.ThreeOne // 水
		} else if 13 == v.PropType {
			useNum = v.FiveOne // 手套
		} else if 14 == v.PropType {
			useNum = v.FourOne // 除虫剂
		} else if 15 == v.PropType {
			useNum = v.TwoOne // 铲子
		}

		res = append(res, &pb.AdminPropConfigReply_List{
			Rate:      v.GetRate,
			PropType:  v.PropType,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Max:       useNum,
		})
	}

	res = append(res, &pb.AdminPropConfigReply_List{
		PropType:  17,
		CreatedAt: "",
		Max:       1,
	})

	return &pb.AdminPropConfigReply{
		Status:     "ok",
		RecordList: res,
		Count:      10,
	}, nil
}

func (ac *AppUsecase) AdminGetBox(ctx context.Context, req *pb.AdminGetBoxRequest) (*pb.AdminGetBoxReply, error) {
	var (
		boxNum     uint64
		boxSellNum uint64
		configs    []*Config
		boxMax     uint64
		boxAmount  float64
		boxStart   string
		boxEnd     string
		err        error
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
		"box_max",
		"box_sell_num",
		"box_start",
		"box_end",
		"box_amount",
	)
	if nil != err || nil == configs {
		return &pb.AdminGetBoxReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "box_sell_num" == vConfig.KeyName {
			boxSellNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "box_start" == vConfig.KeyName {
			boxStart = vConfig.Value
		}
		if "box_end" == vConfig.KeyName {
			boxEnd = vConfig.Value
		}
		if "box_amount" == vConfig.KeyName {
			boxAmount, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "box_max" == vConfig.KeyName {
			boxMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	return &pb.AdminGetBoxReply{
		Start:   boxStart,
		End:     boxEnd,
		Total:   boxMax,
		Amount:  boxAmount,
		SellNum: boxSellNum,
		Term:    boxNum,
	}, nil
}

func (ac *AppUsecase) AdminSetBox(ctx context.Context, req *pb.AdminSetBoxRequest) (*pb.AdminSetBoxReply, error) {
	var (
		boxNum  uint64
		configs []*Config
		err     error
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
	)
	if nil != err || nil == configs {
		return &pb.AdminSetBoxReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		if 1 == req.SendBody.NewTerm {
			err = ac.userRepo.UpdateConfig(ctx, 19, strconv.FormatUint(boxNum+1, 10))
			if nil != err {
				return err
			}

			err = ac.userRepo.UpdateConfig(ctx, 21, "0")
			if nil != err {
				return err
			}
		}

		if 10 < len(req.SendBody.Start) {
			err = ac.userRepo.UpdateConfig(ctx, 16, req.SendBody.Start)
			if nil != err {
				return err
			}
		}

		if 10 < len(req.SendBody.End) {
			err = ac.userRepo.UpdateConfig(ctx, 17, req.SendBody.End)
			if nil != err {
				return err
			}
		}

		if 0 < req.SendBody.Total {
			err = ac.userRepo.UpdateConfig(ctx, 20, strconv.FormatUint(req.SendBody.Total, 10))
			if nil != err {
				return err
			}
		}

		if 0 < req.SendBody.Amount {
			err = ac.userRepo.UpdateConfig(ctx, 18, strconv.FormatFloat(req.SendBody.Amount, 'f', -1, 64))
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return &pb.AdminSetBoxReply{
			Status: "配置修改错误",
		}, nil
	}

	return &pb.AdminSetBoxReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminGetConfig(ctx context.Context, req *pb.AdminGetConfigRequest) (*pb.AdminGetConfigReply, error) {
	var (
		configs []*Config
		err     error
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"withdraw_amount_min",
		"withdraw_amount_max",
		"exchange_fee_rate",
		"reward_stake_rate",
		"stake_over_rate",
		"sell_fee_rate",
		"one_rate",
		"two_rate",
		"three_rate",
		"buy_one", "buy_two", "buy_three", "buy_four", "buy_five", "buy_six", "buy_seven", "buy_eight",
		"recommend_two", "recommend_two_sub",
		"area_one", "area_two", "area_three", "area_four", "area_five", "area_zero",
		"all_each",
		"u_price",
		"b_price",
		"recommend",
		"low_reward_u",
	)
	if nil != err || nil == configs {
		return &pb.AdminGetConfigReply{
			Status: "配置错误",
		}, nil
	}

	res := make([]*pb.AdminGetConfigReply_List, 0)
	for _, vConfig := range configs {
		res = append(res, &pb.AdminGetConfigReply_List{
			Value: vConfig.Value,
			Name:  vConfig.Name,
			Id:    uint64(vConfig.ID),
		})
	}

	return &pb.AdminGetConfigReply{Status: "ok", List: res}, nil
}

func (ac *AppUsecase) AdminSetConfig(ctx context.Context, req *pb.AdminSetConfigRequest) (*pb.AdminSetConfigReply, error) {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		if 1 <= req.SendBody.Id {
			err = ac.userRepo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return &pb.AdminSetConfigReply{
			Status: "配置修改错误",
		}, nil
	}

	return &pb.AdminSetConfigReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminDailyReward(ctx context.Context, req *pb.AdminDailyRewardRequest) (*pb.AdminDailyRewardReply, error) {
	var (
		configs             []*Config
		oneRate             float64
		twoRate             float64
		threeRate           float64
		fourRate            float64
		fiveRate            float64
		sixRate             float64
		sevenRate           float64
		eightRate           float64
		recommendTwoRate    float64
		recommendTwoRateSub float64
		areaOne             float64
		areaTwo             float64
		areaThree           float64
		areaFour            float64
		areaFive            float64
		areaZero            float64
		allEach             float64
		uPrice              float64
		err                 error
	)
	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"buy_one", "buy_two", "buy_three", "buy_four", "buy_five", "buy_six", "buy_seven", "buy_eight",
		"recommend_two", "recommend_two_sub",
		"area_one", "area_two", "area_three", "area_four", "area_five", "area_zero",
		"all_each",
		"u_price",
	)
	if nil != err || nil == configs {
		fmt.Println("错误分红，配置", err)
		return &pb.AdminDailyRewardReply{}, nil
	}
	for _, vConfig := range configs {
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_one" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_two" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_three" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_four" == vConfig.KeyName {
			fourRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_five" == vConfig.KeyName {
			fiveRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_six" == vConfig.KeyName {
			sixRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_seven" == vConfig.KeyName {
			sevenRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_eight" == vConfig.KeyName {
			eightRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "recommend_two" == vConfig.KeyName {
			recommendTwoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "recommend_two_sub" == vConfig.KeyName {
			recommendTwoRateSub, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_one" == vConfig.KeyName {
			areaOne, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_two" == vConfig.KeyName {
			areaTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_three" == vConfig.KeyName {
			areaThree, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_four" == vConfig.KeyName {
			areaFour, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_five" == vConfig.KeyName {
			areaFive, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_zero" == vConfig.KeyName {
			areaZero, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "all_each" == vConfig.KeyName {
			allEach, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[uint64]*UserRecommend
	)
	userRecommends, err = ac.userRepo.GetUserRecommends(ctx)
	if nil != err {
		return &pb.AdminDailyRewardReply{}, nil
	}

	myLowUser := make(map[uint64][]*UserRecommend, 0)
	userRecommendsMap = make(map[uint64]*UserRecommend, 0)
	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId uint64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		users       []*User
		usersReward []*User
		usersMap    map[uint64]*User
	)
	users, err = ac.userRepo.GetAllUsers(ctx)
	if nil == users {
		return &pb.AdminDailyRewardReply{}, nil
	}

	usersMap = make(map[uint64]*User, 0)
	usersReward = make([]*User, 0)
	levelOne := make([]*User, 0)
	levelTwo := make([]*User, 0)
	levelThree := make([]*User, 0)
	levelFour := make([]*User, 0)
	levelFive := make([]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers

		if 0 < vUsers.Amount {
			usersReward = append(usersReward, vUsers)
		}
	}

	for _, vUsers := range users {
		if 0 < vUsers.Vip {
			if 1 == vUsers.Vip {
				levelOne = append(levelOne, vUsers)
			} else if 2 == vUsers.Vip {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
			} else if 3 == vUsers.Vip {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
				levelThree = append(levelThree, vUsers)
			} else if 4 == vUsers.Vip {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
				levelThree = append(levelThree, vUsers)
				levelFour = append(levelFour, vUsers)
			} else if 5 == vUsers.Vip {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
				levelThree = append(levelThree, vUsers)
				levelFour = append(levelFour, vUsers)
				levelFive = append(levelFive, vUsers)
			} else {
				// 跳过，没级别
				continue
			}

			continue
		}

		if 1 >= len(myLowUser[vUsers.ID]) {
			continue
		}

		// 获取业绩
		tmpAreaMax := float64(0)
		tmpAreaMin := float64(0)
		tmpMaxId := uint64(0)
		for _, vMyLowUser := range myLowUser[vUsers.ID] {
			if _, ok := usersMap[vMyLowUser.UserId]; !ok {
				continue
			}

			if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].Amount {
				tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
				tmpMaxId = vMyLowUser.ID
			}
		}

		if 0 >= tmpMaxId {
			continue
		}

		for _, vMyLowUser := range myLowUser[vUsers.ID] {
			if _, ok := usersMap[vMyLowUser.UserId]; !ok {
				continue
			}

			if tmpMaxId != vMyLowUser.ID {
				tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
			}
		}

		if 1000000 <= tmpAreaMin {
			levelFive = append(levelFive, vUsers)
			levelFour = append(levelFour, vUsers)
			levelThree = append(levelThree, vUsers)
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 500000 <= tmpAreaMin {
			levelFour = append(levelFour, vUsers)
			levelThree = append(levelThree, vUsers)
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 150000 <= tmpAreaMin {
			levelThree = append(levelThree, vUsers)
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 50000 <= tmpAreaMin {
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 10000 <= tmpAreaMin {
			levelOne = append(levelOne, vUsers)
		} else {
			// 跳过，没级别
			continue
		}
	}

	stopUserIds := make(map[uint64]bool, 0)

	// 静态
	for _, v := range usersReward {
		tmpUsers := v

		if 0 >= uPrice {
			continue
		}

		// 出局的
		if 0 >= tmpUsers.Amount {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[tmpUsers.ID]; ok {
			continue
		}

		numTwo := float64(0)
		if 30000 <= tmpUsers.Amount {
			numTwo = eightRate
		} else if 15000 <= tmpUsers.Amount {
			numTwo = sevenRate
		} else if 10000 <= tmpUsers.Amount {
			numTwo = sixRate
		} else if 5000 <= tmpUsers.Amount {
			numTwo = fiveRate
		} else if 1000 <= tmpUsers.Amount {
			numTwo = fourRate
		} else if 500 <= tmpUsers.Amount {
			numTwo = threeRate
		} else if 300 <= tmpUsers.Amount {
			numTwo = twoRate
		} else if 100 <= tmpUsers.Amount {
			numTwo = oneRate
		} else {
			continue
		}

		stop := false
		tmp := tmpUsers.Amount * numTwo
		if tmp+tmpUsers.AmountGet >= tmpUsers.Amount*2.5 {
			tmp = math.Abs(tmpUsers.Amount*2.5 - tmpUsers.AmountGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000

		tmp2 := tmp * 0.1
		tmp2 = math.Round(tmp2*10000000) / 10000000

		tmpBiw := tmp / uPrice * 0.9
		tmpBiw = math.Round(tmpBiw*10000000) / 10000000

		if 0 >= tmp || 0 >= tmpBiw || 0 >= tmp2 {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
			userRecommend = userRecommendsMap[tmpUsers.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, tmpUsers)
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateUserRewardNew(ctx, tmpUsers.ID, tmpBiw, tmp2, tmp, tmpUsers.Amount, stop)
			if err != nil {
				fmt.Println("错误分红静态：", err, tmpUsers)
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, tmpUsers)
			continue
		}

		if stop {
			stopUserIds[tmpUsers.ID] = true // 出局

			if nil != userRecommend && "" != userRecommend.RecommendCode {
				var tmpRecommendUserIds []string
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
				for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
					if 0 >= len(tmpRecommendUserIds[j]) {
						continue
					}

					myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
					if 0 >= myUserRecommendUserId {
						continue
					}
					if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
						// 减掉业绩
						err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpUsers.Amount)
						if err != nil {
							fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
							return err
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily 业绩更新", err, tmpUsers)
						continue
					}
				}
			}

		}
	}

	// 团队和平级
	for _, v := range usersReward {
		tmpUsers := v

		// 出局的
		if 0 >= tmpUsers.Amount {
			continue
		}

		if 0 >= uPrice {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[tmpUsers.ID]; ok {
			continue
		}

		numTwo := float64(0)
		if 30000 <= tmpUsers.Amount {
			numTwo = eightRate
		} else if 15000 <= tmpUsers.Amount {
			numTwo = sevenRate
		} else if 10000 <= tmpUsers.Amount {
			numTwo = sixRate
		} else if 5000 <= tmpUsers.Amount {
			numTwo = fiveRate
		} else if 1000 <= tmpUsers.Amount {
			numTwo = fourRate
		} else if 500 <= tmpUsers.Amount {
			numTwo = threeRate
		} else if 300 <= tmpUsers.Amount {
			numTwo = twoRate
		} else if 100 <= tmpUsers.Amount {
			numTwo = oneRate
		} else {
			continue
		}
		tmp := tmpUsers.Amount * numTwo
		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
			userRecommend = userRecommendsMap[tmpUsers.ID]
		} else {
			fmt.Println("错误分红团队，信息缺失：", err, tmpUsers)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		lastKey := len(tmpRecommendUserIds) - 1
		tmpI := uint64(0)
		for i := lastKey; i >= 0; i-- {
			currentLevel := 0
			tmpI++

			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, tmpUsers)
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, tmpUsers)
				continue
			}

			if 0 >= tmpRecommendUser.Amount {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpAreaMin := float64(0)
			tmpMaxId := uint64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].Amount {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
					tmpMaxId = vMyLowUser.ID
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			// 如果是我大区的人，不拿，当前人的下级是不是大区的用户id
			if i == lastKey {
				// 直推，是我的大区
				if tmpMaxId == v.ID {
					continue
				}
			} else {
				if i+1 > lastKey {
					fmt.Println("错误分红小区，信息缺失44：", err, tmpUserId, lastKey, i+1, v)
					continue
				}

				tmpLastUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i+1], 10, 64) // 最后一位是直推人
				if 0 >= tmpLastUserId {
					fmt.Println("错误分红小区，信息缺失445：", err, tmpUserId, lastKey, i+1, v)
					continue
				}

				// 是我大区的人，跳过
				if tmpMaxId == tmpLastUserId {
					continue
				}
			}

			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpMaxId != vMyLowUser.ID {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = areaOne
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = areaTwo
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = areaThree
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = areaFour
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = areaFive
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000000 <= tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = areaFive
				} else if 500000 <= tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = areaFour
				} else if 150000 <= tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = areaFour
				} else if 50000 <= tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = areaFour
				} else if 10000 <= tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = areaFour
				} else {
					// 跳过，没级别
					continue
				}
			}

			var (
				tmpAreaAmount float64
			)
			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmpAreaAmount = tmp * areaZero
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmpAreaAmount = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			if 0 >= tmpAreaAmount {
				continue
			}

			var (
				stopArea bool
			)

			// 平级，结束
			tmpLevel := false
			if currentLevel == lastLevel {
				tmpLevel = true
			}

			if tmpAreaAmount+tmpRecommendUser.AmountGet >= tmpRecommendUser.Amount*2.5 {
				tmpAreaAmount = math.Abs(tmpRecommendUser.Amount*2.5 - tmpRecommendUser.AmountGet)
				stopArea = true
			}

			tmp2 := tmpAreaAmount * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmpAreaAmount / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmpAreaAmount || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

				err = ac.userRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpBiw, tmp2, tmpAreaAmount, tmpRecommendUser.Amount, stopArea, tmpLevel, uint64(currentLevel), tmpI, v.Address)
				if err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, tmpRecommendUser.Amount)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}

		}

	}

	// 直推加速
	for _, v := range usersReward {
		tmpUsers := v

		// 出局的
		if 0 >= tmpUsers.Amount {
			continue
		}

		if 0 >= uPrice {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[tmpUsers.ID]; ok {
			continue
		}

		numTwo := float64(0)
		if 30000 <= tmpUsers.Amount {
			numTwo = eightRate
		} else if 15000 <= tmpUsers.Amount {
			numTwo = sevenRate
		} else if 10000 <= tmpUsers.Amount {
			numTwo = sixRate
		} else if 5000 <= tmpUsers.Amount {
			numTwo = fiveRate
		} else if 1000 <= tmpUsers.Amount {
			numTwo = fourRate
		} else if 500 <= tmpUsers.Amount {
			numTwo = threeRate
		} else if 300 <= tmpUsers.Amount {
			numTwo = twoRate
		} else if 100 <= tmpUsers.Amount {
			numTwo = oneRate
		} else {
			continue
		}
		tmp := tmpUsers.Amount * numTwo
		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
			userRecommend = userRecommendsMap[tmpUsers.ID]
		} else {
			fmt.Println("错误分红团队，信息缺失：", err, tmpUsers)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		tmpI := 0
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			if 10 == tmpI {
				break
			}

			tmpI++

			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红加速，信息缺失,user：", err, tmpUsers)
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红加速，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红加速，信息缺失3：", err, tmpUserId, v)
				continue
			}

			// 2代1个，依次类推
			tmpLen := 0
			for _, vTmp := range myLowUser[tmpUserId] {
				if 0 >= usersMap[vTmp.UserId].Amount && 0 >= usersMap[vTmp.UserId].OutNum {
					continue
				}

				tmpLen++
			}

			if tmpLen < tmpI-1 {
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红加速，信息缺失,user1：", err, tmpUsers)
				continue
			}

			if 0 >= tmpRecommendUser.Amount {
				continue
			}

			var (
				tmpAreaAmount float64
				tmpRate       = 0.01
			)

			if recommendTwoRate >= recommendTwoRateSub*float64(tmpI-1) {
				tmpRate = recommendTwoRate - recommendTwoRateSub*float64(tmpI-1)
			}

			tmpAreaAmount = tmp * tmpRate
			if 0 >= tmpAreaAmount {
				continue
			}

			var (
				stopArea bool
			)

			if tmpAreaAmount+tmpRecommendUser.AmountGet >= tmpRecommendUser.Amount*2.5 {
				tmpAreaAmount = math.Abs(tmpRecommendUser.Amount*2.5 - tmpRecommendUser.AmountGet)
				stopArea = true
			}

			tmp2 := tmpAreaAmount * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmpAreaAmount / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmpAreaAmount || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

				err = ac.userRepo.UpdateUserRewardAreaTwo(ctx, tmpRecommendUser.ID, tmpBiw, tmp2, tmpAreaAmount, tmpRecommendUser.Amount, stopArea, uint64(tmpI), v.Address)
				if err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, tmpRecommendUser.Amount)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}
		}

	}

	var (
		reward []*RewardTwo
	)
	reward, err = ac.userRepo.GetRewardYes(ctx)
	if nil != err {
		return &pb.AdminDailyRewardReply{}, nil
	}

	tmpAmountAll := float64(0)
	for _, vReward := range reward {
		tmpAmountAll += vReward.Amount
	}

	fmt.Println("昨日入金:", tmpAmountAll, len(levelOne), len(levelTwo), len(levelThree), len(levelFour), len(levelFive))
	if 0 >= tmpAmountAll {
		return &pb.AdminDailyRewardReply{}, nil
	}

	tmpRewardAllEach := tmpAmountAll * allEach
	if 0 >= tmpRewardAllEach {
		return &pb.AdminDailyRewardReply{}, nil
	}

	if 0 < len(levelOne) {
		levelOneTmp := tmpRewardAllEach / float64(len(levelOne))

		for _, v := range levelOne {

			tmpUsers := v
			stop := false

			tmp := levelOneTmp
			if tmp+tmpUsers.AmountGet >= tmpUsers.Amount*2.5 {
				tmp = math.Abs(tmpUsers.Amount*2.5 - tmpUsers.AmountGet)
				stop = true
			}
			tmp = math.Round(tmp*10000000) / 10000000

			tmp2 := tmp * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmp / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmp || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 推荐人
			var (
				userRecommend *UserRecommend
			)
			if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
				userRecommend = userRecommendsMap[tmpUsers.ID]
			} else {
				fmt.Println("错误分红all，信息缺失：", err, tmpUsers)
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.UpdateUserRewardNewThree(ctx, tmpUsers.ID, tmpBiw, tmp2, tmp, tmpUsers.Amount, 1, stop)
				if err != nil {
					fmt.Println("错误分红all：", err, tmpUsers)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily all", err, tmpUsers)
				continue
			}

			if stop {
				stopUserIds[tmpUsers.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendUserIds[j]) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}
						if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpUsers.Amount)
							if err != nil {
								fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
								return err
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily all 业绩更新", err, tmpUsers)
							continue
						}
					}
				}

			}

		}
	}

	if 0 < len(levelTwo) {
		levelOneTmp := tmpRewardAllEach / float64(len(levelTwo))

		for _, v := range levelTwo {

			tmpUsers := v
			stop := false

			tmp := levelOneTmp
			if tmp+tmpUsers.AmountGet >= tmpUsers.Amount*2.5 {
				tmp = math.Abs(tmpUsers.Amount*2.5 - tmpUsers.AmountGet)
				stop = true
			}
			tmp = math.Round(tmp*10000000) / 10000000

			tmp2 := tmp * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmp / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmp || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 推荐人
			var (
				userRecommend *UserRecommend
			)
			if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
				userRecommend = userRecommendsMap[tmpUsers.ID]
			} else {
				fmt.Println("错误分红all，信息缺失：", err, tmpUsers)
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.UpdateUserRewardNewThree(ctx, tmpUsers.ID, tmpBiw, tmp2, tmp, tmpUsers.Amount, 2, stop)
				if err != nil {
					fmt.Println("错误分红all：", err, tmpUsers)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily all", err, tmpUsers)
				continue
			}

			if stop {
				stopUserIds[tmpUsers.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendUserIds[j]) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}
						if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpUsers.Amount)
							if err != nil {
								fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
								return err
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily all 业绩更新", err, tmpUsers)
							continue
						}
					}
				}

			}

		}
	}

	if 0 < len(levelThree) {
		levelOneTmp := tmpRewardAllEach / float64(len(levelThree))

		for _, v := range levelThree {

			tmpUsers := v
			stop := false

			tmp := levelOneTmp
			if tmp+tmpUsers.AmountGet >= tmpUsers.Amount*2.5 {
				tmp = math.Abs(tmpUsers.Amount*2.5 - tmpUsers.AmountGet)
				stop = true
			}
			tmp = math.Round(tmp*10000000) / 10000000

			tmp2 := tmp * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmp / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmp || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 推荐人
			var (
				userRecommend *UserRecommend
			)
			if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
				userRecommend = userRecommendsMap[tmpUsers.ID]
			} else {
				fmt.Println("错误分红all，信息缺失：", err, tmpUsers)
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.UpdateUserRewardNewThree(ctx, tmpUsers.ID, tmpBiw, tmp2, tmp, tmpUsers.Amount, 3, stop)
				if err != nil {
					fmt.Println("错误分红all：", err, tmpUsers)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily all", err, tmpUsers)
				continue
			}

			if stop {
				stopUserIds[tmpUsers.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendUserIds[j]) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}
						if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpUsers.Amount)
							if err != nil {
								fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
								return err
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily all 业绩更新", err, tmpUsers)
							continue
						}
					}
				}

			}

		}
	}

	if 0 < len(levelFour) {
		levelOneTmp := tmpRewardAllEach / float64(len(levelFour))

		for _, v := range levelFour {

			tmpUsers := v
			stop := false

			tmp := levelOneTmp
			if tmp+tmpUsers.AmountGet >= tmpUsers.Amount*2.5 {
				tmp = math.Abs(tmpUsers.Amount*2.5 - tmpUsers.AmountGet)
				stop = true
			}
			tmp = math.Round(tmp*10000000) / 10000000

			tmp2 := tmp * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmp / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmp || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 推荐人
			var (
				userRecommend *UserRecommend
			)
			if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
				userRecommend = userRecommendsMap[tmpUsers.ID]
			} else {
				fmt.Println("错误分红all，信息缺失：", err, tmpUsers)
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.UpdateUserRewardNewThree(ctx, tmpUsers.ID, tmpBiw, tmp2, tmp, tmpUsers.Amount, 4, stop)
				if err != nil {
					fmt.Println("错误分红all：", err, tmpUsers)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily all", err, tmpUsers)
				continue
			}

			if stop {
				stopUserIds[tmpUsers.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendUserIds[j]) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}
						if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpUsers.Amount)
							if err != nil {
								fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
								return err
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily all 业绩更新", err, tmpUsers)
							continue
						}
					}
				}

			}

		}
	}

	if 0 < len(levelFive) {
		levelOneTmp := tmpRewardAllEach / float64(len(levelFive))

		for _, v := range levelFive {

			tmpUsers := v
			stop := false

			tmp := levelOneTmp
			if tmp+tmpUsers.AmountGet >= tmpUsers.Amount*2.5 {
				tmp = math.Abs(tmpUsers.Amount*2.5 - tmpUsers.AmountGet)
				stop = true
			}
			tmp = math.Round(tmp*10000000) / 10000000

			tmp2 := tmp * 0.1
			tmp2 = math.Round(tmp2*10000000) / 10000000

			tmpBiw := tmp / uPrice * 0.9
			tmpBiw = math.Round(tmpBiw*10000000) / 10000000

			if 0 >= tmp || 0 >= tmpBiw || 0 >= tmp2 {
				continue
			}

			// 推荐人
			var (
				userRecommend *UserRecommend
			)
			if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
				userRecommend = userRecommendsMap[tmpUsers.ID]
			} else {
				fmt.Println("错误分红all，信息缺失：", err, tmpUsers)
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.UpdateUserRewardNewThree(ctx, tmpUsers.ID, tmpBiw, tmp2, tmp, tmpUsers.Amount, 5, stop)
				if err != nil {
					fmt.Println("错误分红all：", err, tmpUsers)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily all", err, tmpUsers)
				continue
			}

			if stop {
				stopUserIds[tmpUsers.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendUserIds[j]) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}
						if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpUsers.Amount)
							if err != nil {
								fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
								return err
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily all 业绩更新", err, tmpUsers)
							continue
						}
					}
				}

			}

		}
	}

	return &pb.AdminDailyRewardReply{}, nil
}

func (ac *AppUsecase) AdminDaily(ctx context.Context, req *pb.AdminDailyRequest) (*pb.AdminDailyReply, error) {
	var (
		stakeGitRecord  []*StakeGitRecord
		configs         []*Config
		oneRate         float64
		twoRate         float64
		threeRate       float64
		rewardStakeRate float64
		uPrice          float64
		lowRewardU      float64
		err             error
	)
	stakeGitRecord, err = ac.userRepo.GetStakeGitRecords(ctx)
	if nil != err {
		fmt.Println("错误粮仓分红", err)
		return &pb.AdminDailyReply{}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"one_rate", "two_rate", "three_rate", "reward_stake_rate", "u_price", "low_reward_u",
	)
	if nil != err || nil == configs {
		fmt.Println("错误粮仓分红，配置", err)
		return &pb.AdminDailyReply{}, nil
	}

	for _, vConfig := range configs {
		if "one_rate" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "two_rate" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "three_rate" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "reward_stake_rate" == vConfig.KeyName {
			rewardStakeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "low_reward_u" == vConfig.KeyName {
			lowRewardU, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 0 >= uPrice {
		fmt.Println("错误粮仓分红，u和biw价格，配置", uPrice, err)
		return &pb.AdminDailyReply{}, nil
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[uint64]*UserRecommend
	)

	userRecommendsMap = make(map[uint64]*UserRecommend, 0)

	userRecommends, err = ac.userRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr
	}

	// 美区时间16点以后执行的
	lastDay := time.Now().UTC().AddDate(0, 0, -1)
	lastDayStart := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 16, 0, 0, 0, time.UTC)

	for _, v := range stakeGitRecord {
		if v.CreatedAt.After(lastDayStart) {
			continue
		}

		if _, ok := userRecommendsMap[v.UserId]; !ok {
			continue
		}

		if nil == userRecommendsMap[v.UserId] {
			continue
		}

		vUr := userRecommendsMap[v.UserId]
		// 我的直推
		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")

		userIds := make([]uint64, 0)
		tmpIi := 0
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			if 3 <= tmpIi {
				break
			}
			tmpIi++

			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			userIds = append(userIds, tmpUserId)
		}

		usersMap := make(map[uint64]*User, 0)
		if 0 < len(userIds) {
			usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
			if nil != err {
				fmt.Println(err, "每日粮仓，上级查询")
				continue
			}
		}

		tmpAmount := v.Amount * rewardStakeRate

		// 分红，状态变更
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.DailyReward(ctx, v.ID, v.UserId, tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				v.UserId,
				"您在粮仓获得了"+fmt.Sprintf("%.2f", tmpAmount)+"GIT",
				"You've harvest "+fmt.Sprintf("%.2f", tmpAmount)+" GIT from stake",
			)
			if nil != err {
				return err
			}

			// l1-l3，奖励发放
			if tmpAmount > 0 {
				tmpI := 0
				for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
					if 3 <= tmpI {
						break
					}
					tmpI++

					tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
					if 0 >= tmpUserId {
						continue
					}

					if _, ok := usersMap[tmpUserId]; !ok {
						continue
					}

					if lowRewardU > usersMap[tmpUserId].Giw/uPrice {
						continue
					}

					tmpReward := float64(0)

					tmpNum := uint64(6)
					tmpReward = tmpAmount * oneRate
					if 1 == tmpI {

					} else if 2 == tmpI {
						tmpReward = tmpAmount * twoRate
						tmpNum = 9
					} else if 3 == tmpI {
						tmpReward = tmpAmount * threeRate
						tmpNum = 12
					} else {
						break
					}

					// 奖励
					err = ac.userRepo.DailyRewardL(ctx, v.ID, tmpUserId, v.UserId, tmpNum, tmpReward)
					if nil != err {
						return err
					}

					err = ac.userRepo.CreateNotice(
						ctx,
						tmpUserId,
						"您从下级粮仓收获了"+fmt.Sprintf("%.2f", tmpReward)+"GIT",
						"You've harvest "+fmt.Sprintf("%.2f", tmpReward)+" GIT from neighbor stake",
					)
					if nil != err {
						return err
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println(err, "reward daily", v)
			return nil, err
		}
	}

	return nil, nil
}

func (ac *AppUsecase) AdminSetLand(ctx context.Context, req *pb.AdminSetLandRequest) (*pb.AdminSetLandReply, error) {

	var (
		err       error
		landInfos map[uint64]*LandInfo
	)

	landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
	if nil != err {
		return &pb.AdminSetLandReply{
			Status: "信息错误",
		}, nil
	}

	if 0 >= len(landInfos) {
		return &pb.AdminSetLandReply{
			Status: "配置信息错误",
		}, nil
	}

	tmpLevel := req.SendBody.Level
	if _, ok := landInfos[tmpLevel]; !ok {
		return &pb.AdminSetLandReply{
			Status: "级别错误",
		}, nil
	}

	partsAddress := strings.Split(req.SendBody.Address, "&")

	if 0 >= len(partsAddress) {
		return &pb.AdminSetLandReply{
			Status: "地址为空",
		}, nil
	}

	var (
		tmpOne   uint64
		tmpTwo   uint64
		tmpThree uint64
	)
	tmpOne, _ = strconv.ParseUint(req.SendBody.One, 10, 64)
	tmpTwo, _ = strconv.ParseUint(req.SendBody.Two, 10, 64)
	tmpThree, _ = strconv.ParseUint(req.SendBody.Three, 10, 64)

	for _, v := range partsAddress {
		if 20 >= len(v) {
			continue
		}

		if 100 <= len(v) {
			continue
		}

		var (
			user *User
		)

		user, err = ac.userRepo.GetUserByAddress(ctx, v) // 查询用户
		if nil != err || nil == user {
			continue
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		now := time.Now().Unix()
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			_, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     user.ID,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + req.SendBody.Limit,
				Status:     0,
				One:        tmpOne,
				Two:        tmpTwo,
				Three:      tmpThree,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "setLand", user)
			continue
		}
	}

	return &pb.AdminSetLandReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminSetProp(ctx context.Context, req *pb.AdminSetPropRequest) (*pb.AdminSetPropReply, error) {
	var (
		err          error
		propInfos    []*PropInfo
		propInfosMap map[uint64]*PropInfo
	)
	propInfos, err = ac.userRepo.GetAllPropInfo(ctx)
	if nil != err {
		return &pb.AdminSetPropReply{
			Status: "err",
		}, nil
	}

	propInfosMap = make(map[uint64]*PropInfo)
	if 17 == req.SendBody.PropType {

	} else if 16 == req.SendBody.PropType {
		return &pb.AdminSetPropReply{
			Status: "暂不支持手动添加盲盒",
		}, nil
	} else {
		for _, v := range propInfos {
			propInfosMap[v.PropType] = v
		}

		if _, ok := propInfosMap[req.SendBody.PropType]; !ok {
			return &pb.AdminSetPropReply{
				Status: "不存在道具",
			}, nil
		}
	}

	partsAddress := strings.Split(req.SendBody.Address, "&")

	if 0 >= len(partsAddress) {
		return &pb.AdminSetPropReply{
			Status: "地址为空",
		}, nil
	}

	for _, v := range partsAddress {
		if 20 >= len(v) {
			continue
		}

		if 100 <= len(v) {
			continue
		}

		var (
			user *User
		)

		user, err = ac.userRepo.GetUserByAddress(ctx, v) // 查询用户
		if nil != err || nil == user {
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			if 17 == req.SendBody.PropType {
				_, err = ac.userRepo.SetProp(ctx, &Prop{
					UserId:   user.ID,
					PropType: int(req.SendBody.PropType),
				})
				if nil != err {
					return err
				}
			} else if 16 == req.SendBody.PropType {
				return nil
			} else {
				_, err = ac.userRepo.SetProp(ctx, &Prop{
					UserId:   user.ID,
					PropType: int(req.SendBody.PropType),
					OneOne:   int(propInfosMap[req.SendBody.PropType].OneOne),
					OneTwo:   int(propInfosMap[req.SendBody.PropType].OneTwo),
					TwoOne:   int(propInfosMap[req.SendBody.PropType].TwoOne),
					TwoTwo:   propInfosMap[req.SendBody.PropType].TwoTwo,
					ThreeOne: int(propInfosMap[req.SendBody.PropType].ThreeOne),
					FourOne:  int(propInfosMap[req.SendBody.PropType].FourOne),
					FiveOne:  int(propInfosMap[req.SendBody.PropType].FiveOne),
				})
				if nil != err {
					return err
				}
			}

			return nil
		}); nil != err {
			fmt.Println(err, "set prop", user)
			continue
		}

	}

	return &pb.AdminSetPropReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminSetSeed(ctx context.Context, req *pb.AdminSetSeedRequest) (*pb.AdminSetSeedReply, error) {
	var (
		err          error
		seedInfos    []*SeedInfo
		seedInfosMap map[uint64]*SeedInfo
	)
	seedInfos, err = ac.userRepo.GetAllSeedInfo(ctx)
	if nil != err {
		return &pb.AdminSetSeedReply{
			Status: "err",
		}, nil
	}

	seedInfosMap = make(map[uint64]*SeedInfo)
	for _, v := range seedInfos {
		seedInfosMap[v.ID] = v
	}

	if _, ok := seedInfosMap[req.SendBody.SeedId]; !ok {
		return &pb.AdminSetSeedReply{
			Status: "不存在种子",
		}, nil
	}
	partsAddress := strings.Split(req.SendBody.Address, "&")

	if 0 >= len(partsAddress) {
		return &pb.AdminSetSeedReply{
			Status: "地址为空",
		}, nil
	}

	for _, v := range partsAddress {
		if 20 >= len(v) {
			continue
		}

		if 100 <= len(v) {
			continue
		}

		var (
			user *User
		)

		user, err = ac.userRepo.GetUserByAddress(ctx, v) // 查询用户
		if nil != err || nil == user {
			continue
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))

		outMin := int64(seedInfosMap[req.SendBody.SeedId].OutMinAmount)
		outMax := int64(seedInfosMap[req.SendBody.SeedId].OutMaxAmount)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		// 种子
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			_, err = ac.userRepo.SetSeed(ctx, &Seed{
				UserId:       user.ID,
				SeedId:       req.SendBody.SeedId,
				Name:         seedInfosMap[req.SendBody.SeedId].Name,
				OutOverTime:  seedInfosMap[req.SendBody.SeedId].OutOverTime,
				OutMaxAmount: float64(randomNumber),
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "set seed", user)
			continue
		}

	}

	return &pb.AdminSetSeedReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) AdminSetBuyLand(ctx context.Context, req *pb.AdminSetBuyLandRequest) (*pb.AdminSetBuyLandReply, error) {

	var (
		err error
	)

	if 1 <= req.SendBody.Level && 10 >= req.SendBody.Level {

	} else {
		return &pb.AdminSetBuyLandReply{
			Status: "错误的级别",
		}, nil
	}

	var (
		amount    float64
		amountTwo float64
	)
	amount, err = strconv.ParseFloat(req.SendBody.Amount, 10)
	if nil != err {
		return &pb.AdminSetBuyLandReply{
			Status: "金额错误",
		}, nil
	}

	if 1 >= amount {
		return &pb.AdminSetBuyLandReply{
			Status: "金额错误，小于1",
		}, nil
	}

	amountTwo, err = strconv.ParseFloat(req.SendBody.AmountTwo, 10)
	if nil != err {
		return &pb.AdminSetBuyLandReply{
			Status: "一口价金额错误",
		}, nil
	}

	if 1 >= amountTwo {
		return &pb.AdminSetBuyLandReply{
			Status: "一口价金额错误，小于1",
		}, nil
	}

	// 种子
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.SetBuyLand(ctx, &BuyLand{
			Amount:    amount,
			AmountTwo: amountTwo,
			Limit:     uint64(time.Now().Unix()) + req.SendBody.Limit*3600,
			Level:     req.SendBody.Level,
		})
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "set buy land")
		return &pb.AdminSetBuyLandReply{
			Status: "err",
		}, err
	}

	return &pb.AdminSetBuyLandReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) GetEthUserRecordLast(ctx context.Context) (int64, error) {
	return ac.userRepo.GetEthUserRecordLast(ctx)
}

func (ac *AppUsecase) GetEthUserRecordLastTwo(ctx context.Context) (int64, error) {
	return ac.userRepo.GetEthUserRecordLastTwo(ctx)
}

func (ac *AppUsecase) GetUserByAddress(ctx context.Context, Addresses []string) (map[string]*User, error) {
	return ac.userRepo.GetUserByAddresses(ctx, Addresses)
}

func (ac *AppUsecase) DepositNew(ctx context.Context, eth *EthRecord) error {
	// 推荐人
	var (
		err error
	)

	// 推荐
	var (
		userRecommend *UserRecommend
	)
	tmpRecommendUserIds := make([]string, 0)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, eth.UserId)
	if nil == userRecommend || nil != err {
		fmt.Println(err, "deposit err user recommend", eth)
		return nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.AddGiw(ctx, eth.Address, eth.Amount)
		if nil != err {
			return err
		}

		// 加业绩
		dai := uint64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			dai++
			err = ac.userRepo.AddUserTotal(ctx, uint64(tmpUserId), dai, eth.Amount)
			if nil != err {
				return err
			}
		}

		err = ac.userRepo.CreateEth(ctx, eth)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "deposit err", eth)
		return err
	}

	// 奖励
	for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
		tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		var (
			count int64
			user  *User
		)
		count, err = ac.userRepo.GetUserRecommendCount(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(uint64(tmpUserId), 10))
		if nil != err {
			fmt.Println(err, "deposit err reward", eth, count, tmpUserId)
			continue
		}

		if 50 > count {
			continue
		}

		user, err = ac.userRepo.GetUserById(ctx, uint64(tmpUserId))
		if nil != err || nil == user {
			fmt.Println(err, "deposit err user reward", eth, count, tmpUserId)
			continue
		}

		if user.Total <= user.LastRewardTotal {
			continue
		}

		if 100000000 > (user.Total - user.LastRewardTotal) {
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.RewardProp(ctx, 17, uint64(tmpUserId), user.LastRewardTotal+100000000)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "deposit err reward", eth, tmpUserId)
			return err
		}
	}

	return nil
}

func (ac *AppUsecase) DepositNewTwo(ctx context.Context, eth *EthRecord) error {
	// 推荐人
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.AddUsdt(ctx, eth.Address, eth.Amount)
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateEthTwo(ctx, eth)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "deposit err", eth)
		return err
	}

	return nil
}

func (ac *AppUsecase) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error) {
	return ac.userRepo.GetWithdrawPassOrRewardedFirst(ctx)
}

func (ac *AppUsecase) GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*User, error) {
	return ac.userRepo.GetUserByUserIds(ctx, userIds)
}

func (ac *AppUsecase) UpdateWithdrawSuccess(ctx context.Context, id uint64) error {
	return ac.userRepo.UpdateWithdraw(ctx, id, "success")
}

func (ac *AppUsecase) UpdateWithdrawDoing(ctx context.Context, id uint64) error {
	return ac.userRepo.UpdateWithdraw(ctx, id, "doing")
}

func (ac *AppUsecase) AdminRewardListTwo(ctx context.Context, req *pb.AdminRewardListTwoRequest) (*pb.AdminRewardListTwoReply, error) {
	res := make([]*pb.AdminRewardListTwoReply_List, 0)

	var (
		user   *User
		count  int64
		err    error
		userId uint64
		num    uint64
	)

	if 0 < len(req.Address) {
		user, err = ac.userRepo.GetUserByAddress(ctx, req.Address) // 查询用户
		if nil != err || nil == user {
			return &pb.AdminRewardListTwoReply{
				Status: "不存在用户",
			}, nil
		}
		userId = user.ID
	}

	var (
		reward []*RewardTwo
	)
	if 0 < req.RewardType {
		num = req.RewardType
	}

	count, err = ac.userRepo.GetUserRewardTwoPageCount(ctx, userId, num)
	if nil != err {
		return &pb.AdminRewardListTwoReply{
			Status: "不存在数据L，count",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardTwoPage(ctx, userId, num, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.AdminRewardListTwoReply{
			Status: "不存在数据L",
		}, nil
	}

	for _, v := range reward {

		if 1 == v.Reason {
			res = append(res, &pb.AdminRewardListTwoReply_List{
				AmountThree: v.Amount,
				Amount:      v.Five,
				AmountTwo:   v.Three,
				Address:     v.Four,
				Num:         v.One,
				RewardType:  v.Reason,
				CreatedAt:   v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		} else {
			res = append(res, &pb.AdminRewardListTwoReply_List{
				Amount:     v.Three,
				AmountTwo:  v.Amount,
				Address:    v.Four,
				Num:        v.One,
				RewardType: v.Reason,
				CreatedAt:  v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &pb.AdminRewardListTwoReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) AdminRewardList(ctx context.Context, req *pb.AdminRewardListRequest) (*pb.AdminRewardListReply, error) {
	res := make([]*pb.AdminRewardListReply_List, 0)

	var (
		user   *User
		count  int64
		err    error
		userId uint64
		num    uint64
	)

	if 0 < len(req.Address) {
		user, err = ac.userRepo.GetUserByAddress(ctx, req.Address) // 查询用户
		if nil != err || nil == user {
			return &pb.AdminRewardListReply{
				Status: "不存在用户",
			}, nil
		}
		userId = user.ID
	}

	var (
		reward []*Reward
	)
	if 0 < req.RewardType {
		num = req.RewardType
	}

	count, err = ac.userRepo.GetUserRewardAdminPageCount(ctx, userId, num)
	if nil != err {
		return &pb.AdminRewardListReply{
			Status: "不存在数据L，count",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardAdminPage(ctx, userId, num, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.AdminRewardListReply{
			Status: "不存在数据L",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range reward {
		if 0 >= v.One {
			continue
		}

		userIds = append(userIds, v.One)
	}

	usersMap := make(map[uint64]*User, 0)
	if 0 < len(userIds) {
		usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
		if nil != err {
			return &pb.AdminRewardListReply{
				Status: "不存在数据L",
			}, nil
		}
	}

	for _, v := range reward {
		address := ""
		if 0 < v.One {
			if _, ok := usersMap[v.One]; ok {
				address = usersMap[v.One].Address
			}
		}

		res = append(res, &pb.AdminRewardListReply_List{
			Amount:     v.Amount,
			Address:    address,
			RewardType: v.Reason,
			CreatedAt:  v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.AdminRewardListReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) AdminUserBuy(ctx context.Context, req *pb.AdminUserBuyRequest) (*pb.AdminUserBuyReply, error) {
	var (
		address = req.Address
		user    *User
		err     error
	)
	if 0 >= len(address) {
		return &pb.AdminUserBuyReply{
			Status: "不存在用户",
		}, nil
	}

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.AdminUserBuyReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		configs []*Config
		uPrice  float64
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"u_price",
	)
	if nil != err || nil == configs {
		return &pb.AdminUserBuyReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	// 推荐
	var (
		userRecommend   *UserRecommend
		myUserRecommend []*UserRecommend
		team            []*UserRecommend
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.AdminUserBuyReply{
			Status: "推荐错误查询",
		}, nil
	}

	myUserRecommend, err = ac.userRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil == myUserRecommend || nil != err {
		return &pb.AdminUserBuyReply{
			Status: "推荐错误查询",
		}, nil
	}

	team, err = ac.userRepo.GetUserRecommendLikeCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil != err {
		return &pb.AdminUserBuyReply{
			Status: "推荐错误查询",
		}, nil
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = ac.userRepo.GetAllUsers(ctx)
	if nil == users {
		return &pb.AdminUserBuyReply{
			Status: "错误",
		}, nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	// 获取业绩
	tmpAreaMax := float64(0)
	tmpAreaMin := float64(0)
	tmpMaxId := uint64(0)
	tmpTwelve := float64(0)
	tmpRecommendNum := uint64(0)
	for _, vMyLowUser := range myUserRecommend {
		if _, ok := usersMap[vMyLowUser.ID]; !ok {
			continue
		}

		tmpTwelve += usersMap[vMyLowUser.UserId].MyTotalAmount
		tmpRecommendNum += 1
		if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].Amount {
			tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
			tmpMaxId = vMyLowUser.ID
		}
	}

	if 0 < tmpMaxId {
		for _, vMyLowUser := range myUserRecommend {
			if _, ok := usersMap[vMyLowUser.ID]; !ok {
				continue
			}

			if tmpMaxId != vMyLowUser.ID {
				tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
			}
		}
	}

	tmpLevel := uint64(0)
	if 1000000 <= tmpAreaMin {
		tmpLevel = 5
	} else if 500000 <= tmpAreaMin {
		tmpLevel = 4
	} else if 150000 <= tmpAreaMin {
		tmpLevel = 3
	} else if 50000 <= tmpAreaMin {
		tmpLevel = 2
	} else if 10000 <= tmpAreaMin {
		tmpLevel = 1
	}

	tmpBuyNum := uint64(0)
	for _, v := range team {
		if _, ok := usersMap[v.UserId]; !ok {
			continue
		}

		if 0 >= usersMap[v.UserId].OutNum && 0 >= usersMap[v.UserId].Amount {
			continue
		}

		tmpBuyNum++
	}

	tmpFour := float64(0)
	if user.Amount*2.5 <= user.AmountGet {
		tmpFour = 0
	} else {
		tmpFour = user.Amount*2.5 - user.AmountGet
	}

	return &pb.AdminUserBuyReply{
		Status:       "ok",
		One:          user.Amount,
		Two:          2.5,
		Three:        user.AmountGet,
		Four:         tmpFour,
		Five:         user.Location,
		Six:          user.Recommend,
		Seven:        user.RecommendTwo,
		Eight:        user.Area,
		Nine:         user.AreaTwo,
		Ten:          user.All,
		Elven:        user.MyTotalAmount,
		Twelve:       tmpTwelve,
		Thirteen:     tmpAreaMax,
		Fourteen:     tmpAreaMin,
		RecommendNum: tmpRecommendNum,
		Usdt:         user.AmountUsdt,
		Giw:          user.Giw,
		Price:        uPrice,
		TeamNum:      uint64(len(team)),
		BuyNum:       tmpBuyNum,
		Level:        tmpLevel,
	}, nil
}

func (ac *AppUsecase) AdminUserLand(ctx context.Context, req *pb.AdminUserLandRequest) (*pb.AdminUserLandReply, error) {
	res := make([]*pb.AdminUserLandReply_List, 0)
	var (
		address = req.Address
		user    *User
		lands   []*Land
		err     error
	)

	if 0 >= len(address) {
		return &pb.AdminUserLandReply{
			Status: "不存在用户",
		}, nil
	}

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.AdminUserLandReply{
			Status: "不存在用户",
		}, nil
	}
	status := []uint64{0, 1, 2, 3, 4, 5, 8}
	lands, err = ac.userRepo.GetLandByUserID(ctx, user.ID, status, nil)
	if nil != err {
		return &pb.AdminUserLandReply{
			Status: "不存在用户",
		}, nil
	}

	for _, v := range lands {
		statusTmp := v.Status
		if 8 == v.Status {
			statusTmp = 3
		}

		res = append(res, &pb.AdminUserLandReply_List{
			Id:         v.ID,
			Level:      v.Level,
			Health:     v.MaxHealth,
			Status:     statusTmp,
			OutRate:    v.OutPutRate,
			PerHealth:  v.PerHealth,
			RentAmount: v.RentOutPutRate,
			One:        v.One,
			Two:        v.Two,
			Three:      v.Three,
		})
	}

	return &pb.AdminUserLandReply{
		Status: "ok",
		Count:  uint64(len(res)),
		List:   res,
	}, nil
}

func (ac *AppUsecase) AdminUserBackList(ctx context.Context, req *pb.AdminUserBackListRequest) (*pb.AdminUserBackListReply, error) {
	res := make([]*pb.AdminUserBackListReply_List, 0)
	var (
		address = req.Address
		user    *User
		err     error
	)
	if 0 >= len(address) {
		return &pb.AdminUserBackListReply{
			Status: "不存在用户",
		}, nil
	}

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.AdminUserBackListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		seed []*Seed
	)
	seedStatus := []uint64{0, 4}
	seed, err = ac.userRepo.GetSeedByUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.AdminUserBackListReply{
			Status: "查询种子错误",
		}, nil
	}

	for _, vSeed := range seed {
		tmpStatus := uint64(1)
		if 4 == vSeed.Status {
			tmpStatus = 4
		}

		res = append(res, &pb.AdminUserBackListReply_List{
			Id:     vSeed.ID,
			Type:   1,
			Num:    vSeed.SeedId,
			UseNum: 0,
			Status: tmpStatus,
			OutMax: vSeed.OutMaxAmount,
			Time:   vSeed.OutOverTime,
			Amount: vSeed.SellAmount,
		})
	}

	var (
		prop []*Prop
	)
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	propStatus := []uint64{1, 2, 4}
	prop, err = ac.userRepo.GetPropsByUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.AdminUserBackListReply{
			Status: "道具错误",
		}, nil
	}

	for _, vProp := range prop {
		useNum := uint64(0)
		res = append(res, &pb.AdminUserBackListReply_List{
			Id:     vProp.ID,
			Type:   2,
			Num:    uint64(vProp.PropType),
			UseNum: useNum,
			Status: uint64(vProp.Status),
			OutMax: 0,
			Amount: vProp.SellAmount,
		})
	}

	var (
		box []*BoxRecord
	)

	box, err = ac.userRepo.GetUserBoxRecordOpen(ctx, user.ID, 0, false, nil)
	if nil != err {
		return &pb.AdminUserBackListReply{
			Status: "查询盒子错误",
		}, nil
	}

	for _, v := range box {
		res = append(res, &pb.AdminUserBackListReply_List{
			Id:     v.ID,
			Type:   2,
			Num:    16,
			UseNum: 0,
			Status: 0,
			OutMax: 0,
		})
	}

	return &pb.AdminUserBackListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}
