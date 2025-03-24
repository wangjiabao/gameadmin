package data

import (
	"context"
	"game/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID               uint64    `gorm:"primarykey;type:int"`
	Address          string    `gorm:"type:varchar(100);not null"`
	Level            uint64    `gorm:"type:int;not null"`
	Giw              float64   `gorm:"type:decimal(65,20);not null"`
	GiwAdd           float64   `gorm:"type:decimal(65,20);not null"`
	Git              float64   `gorm:"type:decimal(65,20);not null"`
	Total            float64   `gorm:"type:decimal(65,20);not null"`
	TotalOne         float64   `gorm:"type:decimal(65,20);not null"`
	TotalTwo         float64   `gorm:"type:decimal(65,20);not null"`
	TotalThree       float64   `gorm:"type:decimal(65,20);not null"`
	RewardOne        float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwo        float64   `gorm:"type:decimal(65,20);not null"`
	RewardThree      float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwoOne     float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwoTwo     float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwoThree   float64   `gorm:"type:decimal(65,20);not null"`
	RewardThreeOne   float64   `gorm:"type:decimal(65,20);not null"`
	RewardThreeTwo   float64   `gorm:"type:decimal(65,20);not null"`
	RewardThreeThree float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type UserRecommend struct {
	ID            uint64    `gorm:"primarykey;type:int"`
	UserId        uint64    `gorm:"type:int;not null"`
	RecommendCode string    `gorm:"type:varchar(1000);not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type Config struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(45);not null"`
	KeyName   string    `gorm:"type:varchar(45);not null"`
	Value     string    `gorm:"type:varchar(1000);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type BoxRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Num       uint64    `gorm:"type:int;not null"`
	GoodId    uint64    `gorm:"type:int;not null"`
	GoodType  uint64    `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Reward struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Reason    uint64    `gorm:"type:int;not null"`
	One       uint64    `gorm:"type:int;not null"`
	Two       uint64    `gorm:"type:int;not null"`
	Three     float64   `gorm:"type:decimal(65,20);not null"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Seed struct {
	ID           uint64    `gorm:"primarykey;type:int"`
	UserId       uint64    `gorm:"type:int;not null;comment:用户id"`
	SeedId       uint64    `gorm:"type:int;not null;comment:种子信息id"`
	Name         string    `gorm:"type:varchar(45);not null;default:'1';comment:名字"`
	OutAmount    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:产出数量"`
	OutOverTime  uint64    `gorm:"type:int;not null;default:0;comment:成熟时间"`
	OutMaxAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最大产出数量"`
	OutMinAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最小产出数量"`
	Status       uint64    `gorm:"type:int;not null;default:0;comment:状态：0未使用，1使用，出售中4，已售出5，不可用6"`
	SellAmount   float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}

type Land struct {
	ID             uint64    `gorm:"primarykey;type:int"`
	UserId         uint64    `gorm:"type:int;not null;default:0;comment:用户id"`
	Level          uint64    `gorm:"type:int;not null;default:1;comment:级别"`
	OutPutRate     float64   `gorm:"type:decimal(65,20);not null;default:100.00000000000000000000;comment:增产率"`
	RentOutPutRate float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:出租产出比率"`
	MaxHealth      uint64    `gorm:"type:int;not null;default:100;comment:最大可消耗肥沃度"`
	PerHealth      uint64    `gorm:"type:int;not null;default:10;comment:每次消耗肥沃度"`
	LimitDate      uint64    `gorm:"type:int;not null;default:0;comment:使用期限"`
	Status         uint64    `gorm:"type:int;not null;default:0;comment:状态"`
	LocationNum    uint64    `gorm:"type:int;not null;default:0;comment:首页位置"`
	One            uint64    `gorm:"type:int;not null;default:1;comment:可出租"`
	Two            uint64    `gorm:"type:int;not null;default:1;comment:可合成"`
	Three          uint64    `gorm:"type:int;not null;default:1;comment:可出售"`
	SellAmount     float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;"`
	CreatedAt      time.Time `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null"`
}

type LandUserUse struct {
	ID           uint64    `gorm:"primarykey;type:int"`
	LandId       uint64    `gorm:"type:int;not null;comment:用户土地id"`
	Level        uint64    `gorm:"type:int;not null;comment:级别"`
	UserId       uint64    `gorm:"type:int;not null;comment:使用用户id"`
	OwnerUserId  uint64    `gorm:"type:int;not null;comment:拥有者用户id"`
	SeedId       uint64    `gorm:"type:int;not null;comment:种子id"`
	SeedTypeId   uint64    `gorm:"type:int;not null;"`
	Status       uint64    `gorm:"type:int;not null;default:1;comment:状态"`
	BeginTime    uint64    `gorm:"type:int;not null;default:0;comment:开始时间戳"`
	TotalTime    uint64    `gorm:"type:int;not null;default:0;comment:成熟总时长"`
	OverTime     uint64    `gorm:"type:int;not null;default:0;comment:结束时间戳"`
	OutMaxNum    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最大产出"`
	OutNum       float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:已产出"`
	InsectStatus uint64    `gorm:"type:int;not null;default:1;comment:虫子状态"`
	OutSubNum    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:减产数"`
	StealNum     float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:被偷走总数"`
	StopStatus   uint64    `gorm:"type:int;not null;default:1;comment:生长状态"`
	StopTime     uint64    `gorm:"type:int;not null;default:0;comment:暂停时间"`
	SubTime      uint64    `gorm:"type:int;not null;default:0;comment:加速总时长"`
	UseChan      uint64    `gorm:"type:int;not null;default:0;comment:使用铲子次数"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
	One          uint64    `gorm:"type:int;not null;"`
	Two          uint64    `gorm:"type:int;not null;"`
}

type BuyLand struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	Amount    float64   `gorm:"type:decimal(65,20);not null;default:0.0"`
	Status    uint64    `gorm:"type:int;not null;default:1;"`
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"type:datetime;default:null"`
	AmountTwo float64   `gorm:"type:decimal(65,20);not null;default:0.0"`
	Limit     uint64    `gorm:"not null;default:0"`
}

type BuyLandRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	BuyLandID uint64    `gorm:"type:int;not null;"`
	Amount    float64   `gorm:"type:decimal(65,20);not null;default:0.0"`
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"type:datetime;default:null"`
	Status    uint64    `gorm:"type:int;not null;default:1;"`
	UserID    uint64    `gorm:"type:int;not null;"`
}

type ExchangeRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null;comment:用户ID"`
	Git       float64   `gorm:"type:decimal(65,20);not null;comment:git数量"`
	Giw       float64   `gorm:"type:decimal(65,20);not null;comment:giw数量"`
	Fee       float64   `gorm:"type:decimal(65,20);not null;comment:手续费"`
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"type:datetime;default:null"`
}

type Market struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:售卖人用户id"`
	GoodId    uint64    `gorm:"type:int;not null;comment:上级的商品各个表中的id"`
	GoodType  int       `gorm:"type:int;not null;comment:商品类型：1土地，2种子，3道具"`
	Amount    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:售价"`
	Status    int       `gorm:"type:int;not null;default:0;comment:状态：0下架，1上架，2已出售"`
	GetUserId uint64    `gorm:"type:int;not null;default:0;comment:购买人id"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Notice struct {
	ID            uint64    `gorm:"primarykey;type:int"`
	UserId        uint64    `gorm:"type:int;not null;comment:用户id"`
	NoticeContent string    `gorm:"type:varchar(500);not null;comment:消息内容"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type Prop struct {
	ID         uint64    `gorm:"primarykey;type:int"`
	UserId     uint64    `gorm:"type:int;not null;comment:用户id"`
	Status     int       `gorm:"type:int;not null;default:1;comment:状态：未使用1，使用中2，已使用3，出售中4，已出售5，不可用11"`
	PropType   int       `gorm:"type:int;not null;comment:道具类型：1化肥，2铲子，3水，4除虫剂，5手套，6盲盒"`
	OneOne     int       `gorm:"type:int;not null;default:20;comment:化肥土地肥沃度增加量"`
	OneTwo     int       `gorm:"type:int;not null;default:14400;comment:化肥植物加速成熟时间戳"`
	TwoOne     int       `gorm:"type:int;not null;default:7;comment:铲子最大试用次数"`
	TwoTwo     float64   `gorm:"type:decimal(65,20);not null;default:20.00000000000000000000;comment:铲子使用之后获得百分比"`
	ThreeOne   int       `gorm:"type:int;not null;default:1;comment:水最大试用次数"`
	FourOne    int       `gorm:"type:int;not null;default:10;comment:除虫剂最大试用次数"`
	FiveOne    int       `gorm:"type:int;not null;default:20;comment:手套最大可偷植物次数"`
	SellAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

type StakeGet struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Status    int       `gorm:"type:int;not null;default:1;comment:状态：1质押，2已解压"`
	StakeRate float64   `gorm:"type:decimal(65,18);not null;comment:质押比率"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGetPlayRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:玩金额"`
	Reward    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:收益"`
	Status    int       `gorm:"type:int;not null;default:0;comment:状态：1赢了，2输"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGetRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:操作金额"`
	StakeRate float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:比率"`
	Total     float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:操作后金额"`
	StakeType int       `gorm:"type:int;not null;default:0;comment:操作类型：1质押，2解压"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGetTotal struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:总数"`
	Balance   float64   `gorm:"type:decimal(65,18);not null;default:0.0"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGit struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:质押金额"`
	Reward    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:收益总数"`
	Status    uint64    `gorm:"type:int;not null;default:0;comment:状态：1质押，0解压"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGitRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:金额"`
	StakeType int       `gorm:"type:int;not null;default:0;comment:操作类型：1质押，2解压"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Withdraw struct {
	ID        uint64    `gorm:"primarykey;type:int;comment:主键"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    uint64    `gorm:"type:bigint(20);not null;comment:金额"`
	RelAmount uint64    `gorm:"type:bigint(20);not null;comment:实际提现金额"`
	Status    string    `gorm:"type:varchar(45);not null;default:'default';comment:状态"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SeedInfo struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(45);not null;default:'1'" json:"name"`
	OutMinAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000" json:"out_min_amount"`
	OutMaxAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000" json:"out_max_amount"`
	GetRate      float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000" json:"get_rate"`
	OutOverTime  uint64    `gorm:"type:int(10) unsigned;not null;default:0" json:"out_over_time"`
	CreatedAt    time.Time `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null" json:"updated_at"`
}

type LandInfo struct {
	ID                uint64    `gorm:"primaryKey;autoIncrement"`
	Level             uint64    `gorm:"not null;default:1;comment:级别"`
	OutPutRateMax     float64   `gorm:"type:decimal(65,20);not null;default:100.00000000000000000000;comment:最大增产量"`
	OutPutRateMin     float64   `gorm:"type:decimal(65,20);not null;default:100.00000000000000000000;comment:最小增产量"`
	RentOutPutRateMax float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最大出租产出比率"`
	MaxHealth         uint64    `gorm:"not null;default:100;comment:最大可消耗肥沃度"`
	PerHealth         uint64    `gorm:"not null;default:10;comment:每次消耗肥沃度"`
	LimitDateMax      uint64    `gorm:"not null;default:0;comment:最大使用期限"`
	CreatedAt         time.Time `gorm:"type:datetime;not null"`
	UpdatedAt         time.Time `gorm:"type:datetime;not null"`
}

type PropInfo struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	PropType uint64 `gorm:"not null" json:"prop_type"` // 1:化肥, 2:铲子, 3:水, 4:除虫剂, 5:手套

	// 化肥相关字段
	OneOne uint64 `gorm:"not null;default:20" json:"one_one"`    // 化肥土地肥沃度增加量
	OneTwo uint64 `gorm:"not null;default:14400" json:"one_two"` // 化肥植物加速成熟时间

	// 铲子相关字段
	TwoOne uint64  `gorm:"not null;default:7" json:"two_one"`                                           // 铲子最大试用次数
	TwoTwo float64 `gorm:"type:decimal(65,20);not null;default:20.00000000000000000000" json:"two_two"` // 铲子使用获得的百分比

	// 水相关字段
	ThreeOne uint64 `gorm:"not null;default:7" json:"three_one"` // 水最大试用次数

	// 除虫剂相关字段
	FourOne uint64 `gorm:"not null;default:7" json:"four_one"` // 除虫剂最大试用次数

	// 手套相关字段
	FiveOne uint64  `gorm:"not null;default:20" json:"five_one"`                                         // 手套最大可偷次数
	GetRate float64 `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000" json:"get_rate"` // 获取概率

	// 时间字段
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type RandomSeed struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Scene     uint64    `gorm:"not null;unique"` // 业务场景 (如 1=blind_box, 2=plant)
	SeedValue uint64    `gorm:"not null"`        // 随机种子
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetAllUsers .
func (u *UserRepo) GetAllUsers(ctx context.Context) ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	if err := u.data.DB(ctx).Table("user").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Level:            user.Level,
			Giw:              user.Giw,
			GiwAdd:           user.GiwAdd,
			Git:              user.Git,
			Total:            user.Total,
			TotalOne:         user.TotalOne,
			TotalTwo:         user.TotalTwo,
			TotalThree:       user.TotalThree,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			RewardOne:        user.RewardOne,
			RewardTwo:        user.RewardTwo,
			RewardThree:      user.RewardThree,
			RewardTwoOne:     user.RewardTwoOne,
			RewardTwoTwo:     user.RewardTwoTwo,
			RewardTwoThree:   user.RewardTwoThree,
			RewardThreeOne:   user.RewardThreeOne,
			RewardThreeTwo:   user.RewardThreeTwo,
			RewardThreeThree: user.RewardThreeThree,
		})
	}

	return res, nil
}

// GetUserByUserIds .
func (u *UserRepo) GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.User, error) {
	var users []*User

	res := make(map[uint64]*biz.User, 0)
	if err := u.data.DB(ctx).Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res[user.ID] = &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Level:            user.Level,
			Giw:              user.Giw,
			GiwAdd:           user.GiwAdd,
			Git:              user.Git,
			Total:            user.Total,
			TotalOne:         user.TotalOne,
			TotalTwo:         user.TotalTwo,
			TotalThree:       user.TotalThree,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			RewardOne:        user.RewardOne,
			RewardTwo:        user.RewardTwo,
			RewardThree:      user.RewardThree,
			RewardTwoOne:     user.RewardTwoOne,
			RewardTwoTwo:     user.RewardTwoTwo,
			RewardTwoThree:   user.RewardTwoThree,
			RewardThreeOne:   user.RewardThreeOne,
			RewardThreeTwo:   user.RewardThreeTwo,
			RewardThreeThree: user.RewardThreeThree,
		}
	}
	return res, nil
}

// GetUserByAddress .
func (u *UserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user *User
	if err := u.data.DB(ctx).Where("address=?", address).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:               user.ID,
		Address:          user.Address,
		Level:            user.Level,
		Giw:              user.Giw,
		GiwAdd:           user.GiwAdd,
		Git:              user.Git,
		Total:            user.Total,
		TotalOne:         user.TotalOne,
		TotalTwo:         user.TotalTwo,
		TotalThree:       user.TotalThree,
		RewardOne:        user.RewardOne,
		RewardTwo:        user.RewardTwo,
		RewardThree:      user.RewardThree,
		RewardTwoOne:     user.RewardTwoOne,
		RewardTwoTwo:     user.RewardTwoTwo,
		RewardTwoThree:   user.RewardTwoThree,
		RewardThreeOne:   user.RewardThreeOne,
		RewardThreeTwo:   user.RewardThreeTwo,
		RewardThreeThree: user.RewardThreeThree,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}, nil
}

// GetUserRecommendByUserId .
func (u *UserRepo) GetUserRecommendByUserId(ctx context.Context, userId uint64) (*biz.UserRecommend, error) {
	var userRecommend *UserRecommend
	if err := u.data.DB(ctx).Where("user_id=?", userId).Table("user_recommend").First(&userRecommend).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
		CreatedAt:     userRecommend.CreatedAt,
		UpdatedAt:     userRecommend.UpdatedAt,
	}, nil
}

// GetUserRecommendByCode .
func (u *UserRepo) GetUserRecommendByCode(ctx context.Context, code string) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)

	res := make([]*biz.UserRecommend, 0)
	instance := u.data.DB(ctx).Table("user_recommend").Where("recommend_code=?", code)
	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
			UpdatedAt:     userRecommend.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserRecommends .
func (u *UserRepo) GetUserRecommends(ctx context.Context) ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := u.data.DB(ctx).Table("user_recommend").Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
			UpdatedAt:     userRecommend.UpdatedAt,
		})
	}

	return res, nil
}

// CreateUser .
func (u *UserRepo) CreateUser(ctx context.Context, uc *biz.User) (*biz.User, error) {
	var user User
	user.Address = uc.Address

	res := u.data.DB(ctx).Table("user").Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "用户创建失败")
	}

	return &biz.User{
		ID:               user.ID,
		Address:          user.Address,
		Level:            user.Level,
		Giw:              user.Giw,
		GiwAdd:           user.GiwAdd,
		Git:              user.Git,
		Total:            user.Total,
		TotalOne:         user.TotalOne,
		TotalTwo:         user.TotalTwo,
		TotalThree:       user.TotalThree,
		RewardOne:        user.RewardOne,
		RewardTwo:        user.RewardTwo,
		RewardThree:      user.RewardThree,
		RewardTwoOne:     user.RewardTwoOne,
		RewardTwoTwo:     user.RewardTwoTwo,
		RewardTwoThree:   user.RewardTwoThree,
		RewardThreeOne:   user.RewardThreeOne,
		RewardThreeTwo:   user.RewardThreeTwo,
		RewardThreeThree: user.RewardThreeThree,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}, nil
}

// CreateStakeGit 创建 StakeGit 记录
func (u *UserRepo) CreateStakeGit(ctx context.Context, sg *biz.StakeGit) error {
	var stakeGit StakeGit
	stakeGit.UserId = sg.UserId

	res := u.data.DB(ctx).Table("stake_git").Create(&stakeGit)
	if res.Error != nil {
		return errors.New(500, "CREATE_SKATEGIT_ERROR", "StakeGit 记录创建失败")
	}

	return nil
}

// CreateUserRecommend .
func (u *UserRepo) CreateUserRecommend(ctx context.Context, user *biz.User, recommendUser *biz.UserRecommend) (*biz.UserRecommend, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatUint(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var userRecommend UserRecommend
	userRecommend.UserId = user.ID
	userRecommend.RecommendCode = tmpRecommendCode

	res := u.data.DB(ctx).Table("user_recommend").Create(&userRecommend)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_RECOMMEND_ERROR", "用户推荐关系创建失败")
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
		CreatedAt:     userRecommend.CreatedAt,
		UpdatedAt:     userRecommend.UpdatedAt,
	}, nil
}

// GetConfigByKeys .
func (u *UserRepo) GetConfigByKeys(ctx context.Context, keys ...string) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := u.data.DB(ctx).Where("key_name IN (?)", keys).Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// GetStakeGitByUserId .
func (u *UserRepo) GetStakeGitByUserId(ctx context.Context, userId uint64) (*biz.StakeGit, error) {
	var stakeGit *StakeGit
	if err := u.data.DB(ctx).Where("user_id=?", userId).Table("stake_git").First(&stakeGit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "stake GIT ERROR", err.Error())
	}

	return &biz.StakeGit{
		ID:        stakeGit.ID,
		UserId:    stakeGit.UserId,
		Status:    stakeGit.Status,
		Amount:    stakeGit.Amount,
		Reward:    stakeGit.Reward,
		CreatedAt: stakeGit.CreatedAt,
		UpdatedAt: stakeGit.UpdatedAt,
	}, nil
}

// GetBoxRecord .
func (u *UserRepo) GetBoxRecord(ctx context.Context, num uint64) ([]*biz.BoxRecord, error) {
	var boxRecord []*BoxRecord
	res := make([]*biz.BoxRecord, 0)
	if err := u.data.DB(ctx).Where("num=?", num).Table("box_record").Find(&boxRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	for _, v := range boxRecord {
		res = append(res, &biz.BoxRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Num:       v.Num,
			GoodId:    v.GoodId,
			GoodType:  v.GoodType,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}

	return res, nil
}

// GetUserBoxRecord .
func (u *UserRepo) GetUserBoxRecord(ctx context.Context, userId, num uint64, b *biz.Pagination) ([]*biz.BoxRecord, error) {
	var boxRecord []*BoxRecord
	res := make([]*biz.BoxRecord, 0)
	instance := u.data.DB(ctx).Where("user_id=?", userId).Table("box_record")

	if 0 < num {
		instance = instance.Where("num=?", num)
	}

	if err := instance.Order("id desc").Scopes(Paginate(b.PageNum, b.PageSize)).Find(&boxRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	for _, v := range boxRecord {
		res = append(res, &biz.BoxRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Num:       v.Num,
			GoodId:    v.GoodId,
			GoodType:  v.GoodType,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}

	return res, nil
}

// GetUserBoxRecordOpen .
func (u *UserRepo) GetUserBoxRecordOpen(ctx context.Context, userId, num uint64, open bool, b *biz.Pagination) ([]*biz.BoxRecord, error) {
	var boxRecord []*BoxRecord
	res := make([]*biz.BoxRecord, 0)
	instance := u.data.DB(ctx).Where("user_id=?", userId).Table("box_record")

	if 0 < num {
		instance = instance.Where("num=?", num)
	}

	if open {
		instance = instance.Where("good_id > ?", 0)
	} else {
		instance = instance.Where("good_id=?", 0)
	}

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Order("id desc").Find(&boxRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	for _, v := range boxRecord {
		res = append(res, &biz.BoxRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Num:       v.Num,
			GoodId:    v.GoodId,
			GoodType:  v.GoodType,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}

	return res, nil
}

func (u *UserRepo) GetBoxRecordCount(ctx context.Context, num uint64) (int64, error) {
	var count int64

	if err := u.data.DB(ctx).Table("box_record").
		Where("num = ?", num).
		Count(&count).Error; err != nil {
		return 0, errors.New(500, "GetBoxRecordCount ERROR", err.Error())
	}

	return count, nil
}

// GetTotalStakeRate 获取所有用户的质押比率总和
func (u *UserRepo) GetTotalStakeRate(ctx context.Context) (float64, error) {
	var totalStakeRate float64
	if err := u.data.DB(ctx).Table("stake_get").Select("SUM(stake_rate)").Scan(&totalStakeRate).Error; err != nil {
		return 0, errors.New(500, "STAKE_RATE_SUM_ERROR", err.Error())
	}

	return totalStakeRate, nil
}

// GetUserStakeGet .
func (u *UserRepo) GetUserStakeGet(ctx context.Context, userId uint64) (*biz.StakeGet, error) {
	var stakeGet *StakeGet
	if err := u.data.DB(ctx).Table("stake_get").Where("user_id=?", userId).First(&stakeGet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "stake GET ERROR", err.Error())
	}

	return &biz.StakeGet{
		ID:        stakeGet.ID,
		Status:    stakeGet.Status,
		UserId:    stakeGet.UserId,
		StakeRate: stakeGet.StakeRate,
		CreatedAt: stakeGet.CreatedAt,
		UpdatedAt: stakeGet.UpdatedAt,
	}, nil
}

// GetUserRecommendCount .
func (u *UserRepo) GetUserRecommendCount(ctx context.Context, code string) (int64, error) {
	var count int64
	if err := u.data.DB(ctx).Table("user_recommend").Where("recommend_code=?", code).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return count, nil
		}

		return count, errors.New(500, "USER ERROR", err.Error())
	}

	return count, nil
}

// GetUserRecommendByCodePage .
func (u *UserRepo) GetUserRecommendByCodePage(ctx context.Context, code string, b *biz.Pagination) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)

	res := make([]*biz.UserRecommend, 0)
	instance := u.data.DB(ctx).Table("user_recommend").
		Where("recommend_code=?", code).
		Order("id asc").
		Scopes(Paginate(b.PageNum, b.PageSize))
	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
			UpdatedAt:     userRecommend.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserRewardByCodePage .
func (u *UserRepo) GetUserRewardPage(ctx context.Context, userId uint64, reason []uint64, b *biz.Pagination) ([]*biz.Reward, error) {
	var (
		rewards []*Reward
	)

	res := make([]*biz.Reward, 0)
	instance := u.data.DB(ctx).Table("reward").
		Where("user_id=?", userId).
		Where("reason in (?)", reason).
		Order("id desc").
		Scopes(Paginate(b.PageNum, b.PageSize))
	if err := instance.Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "GetUserRewardByCodePage ERROR", err.Error())
	}

	for _, v := range rewards {
		res = append(res, &biz.Reward{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			One:       v.One,
			Two:       v.Two,
			Three:     v.Three,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserRewardPageCount(ctx context.Context, userId uint64, reason []uint64) (int64, error) {
	var count int64

	instance := u.data.DB(ctx).Table("reward").
		Where("user_id = ?", userId).
		Where("reason IN (?)", reason)

	if err := instance.Count(&count).Error; err != nil {
		return 0, errors.New(500, "GetUserRewardPageCount ERROR", err.Error())
	}

	return count, nil
}

// GetSeedByUserID 查询用户的种子数据
func (u *UserRepo) GetSeedByUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Seed, error) {
	var (
		seeds []*Seed
	)

	res := make([]*biz.Seed, 0)
	instance := u.data.DB(ctx).Table("seed")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&seeds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	for _, seed := range seeds {
		res = append(res, &biz.Seed{
			ID:           seed.ID,
			UserId:       seed.UserId,
			SeedId:       seed.SeedId,
			Name:         seed.Name,
			OutAmount:    seed.OutAmount,
			OutOverTime:  seed.OutOverTime,
			OutMaxAmount: seed.OutMaxAmount,
			OutMinAmount: seed.OutMinAmount,
			Status:       seed.Status,
			CreatedAt:    seed.CreatedAt,
			UpdatedAt:    seed.UpdatedAt,
			SellAmount:   seed.SellAmount,
		})
	}

	return res, nil
}

// GetSeedByExUserID 查询用户的种子数据
func (u *UserRepo) GetSeedByExUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Seed, error) {
	var (
		seeds []*Seed
	)

	res := make([]*biz.Seed, 0)
	instance := u.data.DB(ctx).Table("seed").Where("user_id != ?", userID).Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&seeds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	for _, seed := range seeds {
		res = append(res, &biz.Seed{
			ID:           seed.ID,
			UserId:       seed.UserId,
			SeedId:       seed.SeedId,
			Name:         seed.Name,
			OutAmount:    seed.OutAmount,
			OutOverTime:  seed.OutOverTime,
			OutMaxAmount: seed.OutMaxAmount,
			OutMinAmount: seed.OutMinAmount,
			Status:       seed.Status,
			CreatedAt:    seed.CreatedAt,
			UpdatedAt:    seed.UpdatedAt,
			SellAmount:   seed.SellAmount,
		})
	}

	return res, nil
}

// GetLandByUserID getLandByUserID
func (u *UserRepo) GetLandByUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Land, error) {
	var (
		lands []*Land
	)

	res := make([]*biz.Land, 0)
	instance := u.data.DB(ctx).Table("land")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("limit_date>=?", time.Now().Unix()).
		Where("status in (?)", status).
		Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&lands).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	for _, land := range lands {
		res = append(res, &biz.Land{
			ID:             land.ID,
			UserId:         land.UserId,
			Level:          land.Level,
			OutPutRate:     land.OutPutRate,
			RentOutPutRate: land.RentOutPutRate,
			MaxHealth:      land.MaxHealth,
			PerHealth:      land.PerHealth,
			LimitDate:      land.LimitDate,
			Status:         land.Status,
			LocationNum:    land.LocationNum,
			CreatedAt:      land.CreatedAt,
			UpdatedAt:      land.UpdatedAt,
			One:            land.One,
			Two:            land.Two,
			Three:          land.Three,
			SellAmount:     land.SellAmount,
		})
	}

	return res, nil
}

// GetSeedByID 根据种子 ID 查询单条种子数据
func (u *UserRepo) GetSeedByID(ctx context.Context, seedID, userId, status uint64) (*biz.Seed, error) {
	var seed Seed

	if err := u.data.DB(ctx).Table("seed").Where("id = ?", seedID).Where("user_id = ?", userId).Where("status=?", status).First(&seed).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 代表未找到
		}
		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	return &biz.Seed{
		ID:           seed.ID,
		UserId:       seed.UserId,
		SeedId:       seed.SeedId,
		Name:         seed.Name,
		OutAmount:    seed.OutAmount,
		OutOverTime:  seed.OutOverTime,
		OutMaxAmount: seed.OutMaxAmount,
		OutMinAmount: seed.OutMinAmount,
		Status:       seed.Status,
		CreatedAt:    seed.CreatedAt,
		UpdatedAt:    seed.UpdatedAt,
		SellAmount:   seed.SellAmount,
	}, nil
}

// GetSeedBuyByID 根据种子 ID 查询单条种子数据
func (u *UserRepo) GetSeedBuyByID(ctx context.Context, seedID, status uint64) (*biz.Seed, error) {
	var seed Seed

	if err := u.data.DB(ctx).Table("seed").Where("id = ?", seedID).Where("status=?", status).First(&seed).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 代表未找到
		}
		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	return &biz.Seed{
		ID:           seed.ID,
		UserId:       seed.UserId,
		SeedId:       seed.SeedId,
		Name:         seed.Name,
		OutAmount:    seed.OutAmount,
		OutOverTime:  seed.OutOverTime,
		OutMaxAmount: seed.OutMaxAmount,
		OutMinAmount: seed.OutMinAmount,
		Status:       seed.Status,
		CreatedAt:    seed.CreatedAt,
		UpdatedAt:    seed.UpdatedAt,
		SellAmount:   seed.SellAmount,
	}, nil
}

// GetLandByUserIDUsing getLandByUserIDUsing
func (u *UserRepo) GetLandByUserIDUsing(ctx context.Context, userID uint64, status []uint64) ([]*biz.Land, error) {
	var (
		lands []*Land
	)

	res := make([]*biz.Land, 0)
	instance := u.data.DB(ctx).Table("land").
		Where("user_id = ?", userID).
		Where("limit_date>=?", time.Now().Unix()).
		Where("status in (?)", status).
		Where("location_num >?", 0).
		Order("id asc")

	if err := instance.Find(&lands).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	for _, land := range lands {
		res = append(res, &biz.Land{
			ID:             land.ID,
			UserId:         land.UserId,
			Level:          land.Level,
			OutPutRate:     land.OutPutRate,
			RentOutPutRate: land.RentOutPutRate,
			MaxHealth:      land.MaxHealth,
			PerHealth:      land.PerHealth,
			LimitDate:      land.LimitDate,
			Status:         land.Status,
			LocationNum:    land.LocationNum,
			CreatedAt:      land.CreatedAt,
			UpdatedAt:      land.UpdatedAt,
			One:            land.One,
			Two:            land.Two,
			Three:          land.Three,
			SellAmount:     land.SellAmount,
		})
	}

	return res, nil
}

// GetLandByExUserID getLandByExUserID
func (u *UserRepo) GetLandByExUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Land, error) {
	var (
		lands []*Land
	)

	res := make([]*biz.Land, 0)
	instance := u.data.DB(ctx).Table("land").Where("user_id != ?", userID).Where("limit_date>=?", time.Now().Unix()).
		Where("status in (?)", status).
		Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&lands).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	for _, land := range lands {
		res = append(res, &biz.Land{
			ID:             land.ID,
			UserId:         land.UserId,
			Level:          land.Level,
			OutPutRate:     land.OutPutRate,
			RentOutPutRate: land.RentOutPutRate,
			MaxHealth:      land.MaxHealth,
			PerHealth:      land.PerHealth,
			LimitDate:      land.LimitDate,
			Status:         land.Status,
			LocationNum:    land.LocationNum,
			CreatedAt:      land.CreatedAt,
			UpdatedAt:      land.UpdatedAt,
			One:            land.One,
			Two:            land.Two,
			Three:          land.Three,
			SellAmount:     land.SellAmount,
		})
	}

	return res, nil
}

// GetLandUserUseByUserIDUseing getLandUserUseByUserIDUseing
func (u *UserRepo) GetLandUserUseByUserIDUseing(ctx context.Context, userID uint64, status uint64, b *biz.Pagination) ([]*biz.LandUserUse, error) {
	var (
		landUserUses []*LandUserUse
	)

	res := make([]*biz.LandUserUse, 0)
	instance := u.data.DB(ctx).Table("land_user_use").
		Where("user_id = ?", userID).
		Where("status = ?", status).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&landUserUses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND USER USE ERROR", err.Error())
	}

	for _, landUserUse := range landUserUses {
		res = append(res, &biz.LandUserUse{
			ID:           landUserUse.ID,
			LandId:       landUserUse.LandId,
			Level:        landUserUse.Level,
			UserId:       landUserUse.UserId,
			OwnerUserId:  landUserUse.OwnerUserId,
			SeedId:       landUserUse.SeedId,
			SeedTypeId:   landUserUse.SeedTypeId,
			Status:       landUserUse.Status,
			BeginTime:    landUserUse.BeginTime,
			TotalTime:    landUserUse.TotalTime,
			OverTime:     landUserUse.OverTime,
			OutMaxNum:    landUserUse.OutMaxNum,
			OutNum:       landUserUse.OutNum,
			InsectStatus: landUserUse.InsectStatus,
			OutSubNum:    landUserUse.OutSubNum,
			StealNum:     landUserUse.StealNum,
			StopStatus:   landUserUse.StopStatus,
			StopTime:     landUserUse.StopTime,
			SubTime:      landUserUse.SubTime,
			UseChan:      landUserUse.UseChan,
			CreatedAt:    landUserUse.CreatedAt,
			UpdatedAt:    landUserUse.UpdatedAt,
			One:          landUserUse.One,
			Two:          landUserUse.Two,
		})
	}

	return res, nil
}

func (u *UserRepo) GetExchangeRecordsByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.ExchangeRecord, error) {
	var (
		exchangeRecords []*ExchangeRecord
	)

	res := make([]*biz.ExchangeRecord, 0)
	instance := u.data.DB(ctx).Table("exchange_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&exchangeRecords).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "EXCHANGE RECORD ERROR", err.Error())
	}

	for _, exchangeRecord := range exchangeRecords {
		res = append(res, &biz.ExchangeRecord{
			ID:        exchangeRecord.ID,
			UserId:    exchangeRecord.UserId,
			Git:       exchangeRecord.Git,
			Giw:       exchangeRecord.Giw,
			Fee:       exchangeRecord.Fee,
			CreatedAt: exchangeRecord.CreatedAt,
			UpdatedAt: exchangeRecord.UpdatedAt,
		})
	}

	return res, nil
}

// GetLandUserUseByID 根据 ID 获取一条 LandUserUse 记录
func (u *UserRepo) GetLandUserUseByID(ctx context.Context, id uint64) (*biz.LandUserUse, error) {
	var landUserUse LandUserUse

	res := u.data.DB(ctx).Table("land_user_use").Where("id = ?", id).Where("status=?", 1).First(&landUserUse)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "GET_LAND_USER_USE_ERROR", res.Error.Error())
	}

	return &biz.LandUserUse{
		ID:           landUserUse.ID,
		LandId:       landUserUse.LandId,
		Level:        landUserUse.Level,
		UserId:       landUserUse.UserId,
		OwnerUserId:  landUserUse.OwnerUserId,
		SeedId:       landUserUse.SeedId,
		SeedTypeId:   landUserUse.SeedTypeId,
		Status:       landUserUse.Status,
		BeginTime:    landUserUse.BeginTime,
		TotalTime:    landUserUse.TotalTime,
		OverTime:     landUserUse.OverTime,
		OutMaxNum:    landUserUse.OutMaxNum,
		OutNum:       landUserUse.OutNum,
		InsectStatus: landUserUse.InsectStatus,
		OutSubNum:    landUserUse.OutSubNum,
		StealNum:     landUserUse.StealNum,
		StopStatus:   landUserUse.StopStatus,
		StopTime:     landUserUse.StopTime,
		SubTime:      landUserUse.SubTime,
		UseChan:      landUserUse.UseChan,
		CreatedAt:    landUserUse.CreatedAt,
		UpdatedAt:    landUserUse.UpdatedAt,
		One:          landUserUse.One,
		Two:          landUserUse.Two,
	}, nil
}

func (u *UserRepo) GetMarketRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *biz.Pagination) ([]*biz.Market, error) {
	var (
		marketRecords []*Market
	)

	res := make([]*biz.Market, 0)
	instance := u.data.DB(ctx).Table("market").
		Where("user_id = ?", userID).
		Where("status in (?)", status).
		Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&marketRecords).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "MARKET RECORD ERROR", err.Error())
	}

	for _, marketRecord := range marketRecords {
		res = append(res, &biz.Market{
			ID:        marketRecord.ID,
			UserId:    marketRecord.UserId,
			GoodId:    marketRecord.GoodId,
			GoodType:  marketRecord.GoodType,
			Amount:    marketRecord.Amount,
			Status:    marketRecord.Status,
			GetUserId: marketRecord.GetUserId,
			CreatedAt: marketRecord.CreatedAt,
			UpdatedAt: marketRecord.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNoticesByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.Notice, error) {
	var (
		noticeRecords []*Notice
	)

	res := make([]*biz.Notice, 0)
	instance := u.data.DB(ctx).Table("notice").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&noticeRecords).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "NOTICE RECORD ERROR", err.Error())
	}

	for _, noticeRecord := range noticeRecords {
		res = append(res, &biz.Notice{
			ID:            noticeRecord.ID,
			UserId:        noticeRecord.UserId,
			NoticeContent: noticeRecord.NoticeContent,
			CreatedAt:     noticeRecord.CreatedAt,
			UpdatedAt:     noticeRecord.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNoticesCountByUserID(ctx context.Context, userID uint64) (int64, error) {
	var count int64
	err := u.data.DB(ctx).Table("notice").
		Where("user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return 0, errors.New(500, "NOTICE COUNT ERROR", err.Error())
	}
	return count, nil
}

func (u *UserRepo) GetPropsByUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Prop, error) {
	var (
		props []*Prop
	)

	res := make([]*biz.Prop, 0)
	instance := u.data.DB(ctx).Table("prop")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&props).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PROP RECORD ERROR", err.Error())
	}

	for _, prop := range props {
		res = append(res, &biz.Prop{
			ID:         prop.ID,
			UserId:     prop.UserId,
			Status:     prop.Status,
			PropType:   prop.PropType,
			OneOne:     prop.OneOne,
			OneTwo:     prop.OneTwo,
			TwoOne:     prop.TwoOne,
			TwoTwo:     prop.TwoTwo,
			ThreeOne:   prop.ThreeOne,
			FourOne:    prop.FourOne,
			FiveOne:    prop.FiveOne,
			CreatedAt:  prop.CreatedAt,
			UpdatedAt:  prop.UpdatedAt,
			SellAmount: prop.SellAmount,
		})
	}

	return res, nil
}

func (u *UserRepo) GetPropsByUserIDPropType(ctx context.Context, userID uint64, propType []uint64) ([]*biz.Prop, error) {
	var (
		props []*Prop
	)

	res := make([]*biz.Prop, 0)
	instance := u.data.DB(ctx).Table("prop")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("status=?", 1).Where("prop_type in (?)", propType).Order("id asc")

	if err := instance.Find(&props).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PROP RECORD ERROR", err.Error())
	}

	for _, prop := range props {
		res = append(res, &biz.Prop{
			ID:         prop.ID,
			UserId:     prop.UserId,
			Status:     prop.Status,
			PropType:   prop.PropType,
			OneOne:     prop.OneOne,
			OneTwo:     prop.OneTwo,
			TwoOne:     prop.TwoOne,
			TwoTwo:     prop.TwoTwo,
			ThreeOne:   prop.ThreeOne,
			FourOne:    prop.FourOne,
			FiveOne:    prop.FiveOne,
			CreatedAt:  prop.CreatedAt,
			UpdatedAt:  prop.UpdatedAt,
			SellAmount: prop.SellAmount,
		})
	}

	return res, nil
}

func (u *UserRepo) GetPropsByExUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Prop, error) {
	var (
		props []*Prop
	)

	res := make([]*biz.Prop, 0)
	instance := u.data.DB(ctx).Table("prop").Where("user_id != ?", userID).Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&props).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PROP RECORD ERROR", err.Error())
	}

	for _, prop := range props {
		res = append(res, &biz.Prop{
			ID:         prop.ID,
			UserId:     prop.UserId,
			Status:     prop.Status,
			PropType:   prop.PropType,
			OneOne:     prop.OneOne,
			OneTwo:     prop.OneTwo,
			TwoOne:     prop.TwoOne,
			TwoTwo:     prop.TwoTwo,
			ThreeOne:   prop.ThreeOne,
			FourOne:    prop.FourOne,
			FiveOne:    prop.FiveOne,
			CreatedAt:  prop.CreatedAt,
			UpdatedAt:  prop.UpdatedAt,
			SellAmount: prop.SellAmount,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetsByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.StakeGet, error) {
	var (
		stakeGets []*StakeGet
	)

	res := make([]*biz.StakeGet, 0)
	instance := u.data.DB(ctx).Table("stake_get").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&stakeGets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "STAKE GET RECORD ERROR", err.Error())
	}

	for _, stakeGet := range stakeGets {
		res = append(res, &biz.StakeGet{
			ID:        stakeGet.ID,
			UserId:    stakeGet.UserId,
			Status:    stakeGet.Status,
			StakeRate: stakeGet.StakeRate,
			CreatedAt: stakeGet.CreatedAt,
			UpdatedAt: stakeGet.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetPlayRecordCount(ctx context.Context, userID uint64, status uint64) (int64, error) {
	var count int64
	instance := u.data.DB(ctx).Table("stake_get_play_record").
		Where("user_id = ?", userID)

	if status > 0 {
		instance = instance.Where("status = ?", status)
	}

	err := instance.Count(&count).Error
	if err != nil {
		return 0, errors.New(500, "STAKE GET PLAY RECORD COUNT ERROR", err.Error())
	}
	return count, nil
}

func (u *UserRepo) GetStakeGetPlayRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *biz.Pagination) ([]*biz.StakeGetPlayRecord, error) {
	var (
		records []*StakeGetPlayRecord
	)

	res := make([]*biz.StakeGetPlayRecord, 0)
	instance := u.data.DB(ctx).Table("stake_get_play_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if 0 < status {
		instance = instance.Where("status=?", status)
	}

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "STAKE GET PLAY RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.StakeGetPlayRecord{
			ID:        record.ID,
			UserId:    record.UserId,
			Amount:    record.Amount,
			Reward:    record.Reward,
			Status:    record.Status,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetRecordsByUserID(ctx context.Context, userID int64, b *biz.Pagination) ([]*biz.StakeGetRecord, error) {
	var (
		records []*StakeGetRecord
	)

	res := make([]*biz.StakeGetRecord, 0)
	instance := u.data.DB(ctx).Table("stake_get_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "STAKE GET RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.StakeGetRecord{
			ID:        record.ID,
			UserId:    record.UserId,
			Amount:    record.Amount,
			StakeRate: record.StakeRate,
			Total:     record.Total,
			StakeType: record.StakeType,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetTotal(ctx context.Context) (*biz.StakeGetTotal, error) {
	var total StakeGetTotal

	err := u.data.DB(ctx).Table("stake_get_total").First(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKE GET TOTAL ERROR", err.Error())
	}

	return &biz.StakeGetTotal{
		ID:        total.ID,
		Amount:    total.Amount,
		Balance:   total.Balance,
		CreatedAt: total.CreatedAt,
		UpdatedAt: total.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetStakeGitByUserID(ctx context.Context, userID int64) (*biz.StakeGit, error) {
	var stake StakeGit

	err := u.data.DB(ctx).Table("stake_git").
		Where("user_id = ?", userID).
		First(&stake).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKE GIT ERROR", err.Error())
	}

	return &biz.StakeGit{
		ID:        stake.ID,
		UserId:    stake.UserId,
		Amount:    stake.Amount,
		Reward:    stake.Reward,
		Status:    stake.Status,
		CreatedAt: stake.CreatedAt,
		UpdatedAt: stake.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetStakeGitRecordsByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.StakeGitRecord, error) {
	var (
		records []*StakeGitRecord
	)

	res := make([]*biz.StakeGitRecord, 0)
	instance := u.data.DB(ctx).Table("stake_git_record").
		Where("user_id = ?", userID).
		Where("stake_type=?", 1).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "STAKE GIT RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.StakeGitRecord{
			ID:        record.ID,
			UserId:    record.UserId,
			Amount:    record.Amount,
			StakeType: record.StakeType,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGitRecordsByID(ctx context.Context, id, userId uint64) (*biz.StakeGitRecord, error) {
	var (
		record StakeGitRecord
	)

	instance := u.data.DB(ctx).Table("stake_git_record").
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Where("stake_type=?", 1).
		Order("id desc")

	if err := instance.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKE GIT RECORD ERROR", err.Error())
	}

	return &biz.StakeGitRecord{
		ID:        record.ID,
		UserId:    record.UserId,
		Amount:    record.Amount,
		StakeType: record.StakeType,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetWithdrawRecordsByUserID(ctx context.Context, userID int64, b *biz.Pagination) ([]*biz.Withdraw, error) {
	var (
		records []*Withdraw
	)

	res := make([]*biz.Withdraw, 0)
	instance := u.data.DB(ctx).Table("withdraw").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "WITHDRAW RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.Withdraw{
			ID:        record.ID,
			UserID:    record.UserId,
			Amount:    record.Amount,
			RelAmount: record.RelAmount,
			Status:    record.Status,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserOrderCount(ctx context.Context) (int64, error) {
	var count int64
	err := u.data.DB(ctx).Table("user").
		Where("git > ?", 0).
		Count(&count).Error

	if err != nil {
		return 0, errors.New(500, "USER COUNT ERROR", err.Error())
	}

	return count, nil
}

func (u *UserRepo) GetUserOrder(ctx context.Context, b *biz.Pagination) ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	instance := u.data.DB(ctx).Table("user").Where("git>?", 0).Order("git desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Level:            user.Level,
			Giw:              user.Giw,
			GiwAdd:           user.GiwAdd,
			Git:              user.Git,
			Total:            user.Total,
			TotalOne:         user.TotalOne,
			TotalTwo:         user.TotalTwo,
			TotalThree:       user.TotalThree,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			RewardOne:        user.RewardOne,
			RewardTwo:        user.RewardTwo,
			RewardThree:      user.RewardThree,
			RewardTwoOne:     user.RewardTwoOne,
			RewardTwoTwo:     user.RewardTwoTwo,
			RewardTwoThree:   user.RewardTwoThree,
			RewardThreeOne:   user.RewardThreeOne,
			RewardThreeTwo:   user.RewardThreeTwo,
			RewardThreeThree: user.RewardThreeThree,
		})
	}

	return res, nil
}

func (u *UserRepo) GetLandUserUseByLandIDsMapUsing(ctx context.Context, userId uint64, landIDs []uint64) (map[uint64]*biz.LandUserUse, error) {
	var landUserUses []*LandUserUse
	res := make(map[uint64]*biz.LandUserUse)

	instance := u.data.DB(ctx).Table("land_user_use").
		Where("owner_user_id = ?", userId).
		Where("land_id IN (?)", landIDs).
		Where("status = ?", 1).
		Order("id desc")

	if err := instance.Find(&landUserUses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "LAND USER USE ERROR", err.Error())
	}

	// 归类数据到 map
	for _, landUserUse := range landUserUses {
		res[landUserUse.LandId] = &biz.LandUserUse{
			ID:           landUserUse.ID,
			LandId:       landUserUse.LandId,
			Level:        landUserUse.Level,
			UserId:       landUserUse.UserId,
			OwnerUserId:  landUserUse.OwnerUserId,
			SeedId:       landUserUse.SeedId,
			SeedTypeId:   landUserUse.SeedTypeId,
			Status:       landUserUse.Status,
			BeginTime:    landUserUse.BeginTime,
			TotalTime:    landUserUse.TotalTime,
			OverTime:     landUserUse.OverTime,
			OutMaxNum:    landUserUse.OutMaxNum,
			OutNum:       landUserUse.OutNum,
			InsectStatus: landUserUse.InsectStatus,
			OutSubNum:    landUserUse.OutSubNum,
			StealNum:     landUserUse.StealNum,
			StopStatus:   landUserUse.StopStatus,
			StopTime:     landUserUse.StopTime,
			SubTime:      landUserUse.SubTime,
			UseChan:      landUserUse.UseChan,
			CreatedAt:    landUserUse.CreatedAt,
			UpdatedAt:    landUserUse.UpdatedAt,
			One:          landUserUse.One,
			Two:          landUserUse.Two,
		}
	}

	return res, nil
}

// BuyBox .
func (u *UserRepo) BuyBox(ctx context.Context, giw float64, originValue, value string, uc *biz.BoxRecord) (uint64, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", uc.UserId).Where("giw>=?", giw).
		Updates(map[string]interface{}{"giw": gorm.Expr("giw - ?", giw), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("config").Where("id=?", 21).Where("value=?", originValue).
		Updates(map[string]interface{}{"value": value, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "config")
	}

	var box BoxRecord

	box.UserId = uc.UserId
	box.Num = uc.Num

	res = u.data.DB(ctx).Table("box_record").Create(&box)
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "创建失败")
	}

	return box.ID, nil
}

// GetUserBoxRecordById .
func (u *UserRepo) GetUserBoxRecordById(ctx context.Context, id uint64) (*biz.BoxRecord, error) {
	var v *BoxRecord
	instance := u.data.DB(ctx).Where("id=?", id).Table("box_record")

	if err := instance.First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	return &biz.BoxRecord{
		ID:        v.ID,
		UserId:    v.UserId,
		Num:       v.Num,
		GoodId:    v.GoodId,
		GoodType:  v.GoodType,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Content:   v.Content,
	}, nil
}

// OpenBoxSeed .
func (u *UserRepo) OpenBoxSeed(ctx context.Context, id uint64, content string, seedInfo *biz.Seed) (uint64, error) {
	res := u.data.DB(ctx).Table("box_record").Where("id=?", id).Where("good_id=?", 0).
		Updates(map[string]interface{}{"good_id": seedInfo.SeedId, "good_type": 1, "content": content, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "config")
	}

	var seed Seed
	seed.SeedId = seedInfo.SeedId
	seed.UserId = seedInfo.UserId
	seed.OutMaxAmount = seedInfo.OutMaxAmount
	seed.OutOverTime = seedInfo.OutOverTime
	res = u.data.DB(ctx).Table("seed").Create(&seed)
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "创建失败")
	}

	return seed.ID, nil
}

// OpenBoxProp .
func (u *UserRepo) OpenBoxProp(ctx context.Context, id uint64, content string, propInfo *biz.Prop) (uint64, error) {
	res := u.data.DB(ctx).Table("box_record").Where("id=?", id).Where("good_id=?", 0).
		Updates(map[string]interface{}{"good_id": propInfo.PropType, "good_type": 2, "content": content, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "config")
	}

	var prop Prop
	prop.PropType = propInfo.PropType
	prop.UserId = propInfo.UserId
	prop.OneOne = propInfo.OneOne
	prop.OneTwo = propInfo.OneTwo
	prop.TwoOne = propInfo.TwoOne
	prop.TwoTwo = propInfo.TwoTwo
	prop.ThreeOne = propInfo.ThreeOne
	prop.FourOne = propInfo.FourOne
	prop.FiveOne = propInfo.FiveOne
	res = u.data.DB(ctx).Table("prop").Create(&prop)
	if res.Error != nil {
		return 0, errors.New(500, "BuyBox", "创建失败")
	}

	return prop.ID, nil
}

// GetAllSeedInfo 获取所有种子信息
func (u *UserRepo) GetAllSeedInfo(ctx context.Context) ([]*biz.SeedInfo, error) {
	var seeds []*SeedInfo

	res := make([]*biz.SeedInfo, 0)
	if err := u.data.DB(ctx).Table("seed_info").Order("id asc").Find(&seeds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	for _, seed := range seeds {
		res = append(res, &biz.SeedInfo{
			ID:           seed.ID,
			Name:         seed.Name,
			OutMinAmount: seed.OutMinAmount,
			OutMaxAmount: seed.OutMaxAmount,
			GetRate:      seed.GetRate,
			OutOverTime:  seed.OutOverTime,
			CreatedAt:    seed.CreatedAt,
			UpdatedAt:    seed.UpdatedAt,
		})
	}

	return res, nil
}

// GetAllPropInfo 获取所有道具信息
func (u *UserRepo) GetAllPropInfo(ctx context.Context) ([]*biz.PropInfo, error) {
	var props []*PropInfo

	res := make([]*biz.PropInfo, 0)
	if err := u.data.DB(ctx).Table("prop_info").Order("id asc").Find(&props).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PROP ERROR", err.Error())
	}

	for _, prop := range props {
		res = append(res, &biz.PropInfo{
			ID:        prop.ID,
			PropType:  prop.PropType,
			OneOne:    prop.OneOne,
			OneTwo:    prop.OneTwo,
			TwoOne:    prop.TwoOne,
			TwoTwo:    prop.TwoTwo,
			ThreeOne:  prop.ThreeOne,
			FourOne:   prop.FourOne,
			FiveOne:   prop.FiveOne,
			GetRate:   prop.GetRate,
			CreatedAt: prop.CreatedAt,
			UpdatedAt: prop.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetAllRandomSeeds(ctx context.Context) ([]*biz.RandomSeed, error) {
	var seeds []*RandomSeed

	res := make([]*biz.RandomSeed, 0)
	if err := u.data.DB(ctx).Table("random_seeds").Order("id asc").Find(&seeds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	for _, seed := range seeds {
		res = append(res, &biz.RandomSeed{
			ID:        seed.ID,
			Scene:     seed.Scene,
			SeedValue: seed.SeedValue,
			CreatedAt: seed.CreatedAt,
			UpdatedAt: seed.UpdatedAt,
		})
	}

	return res, nil
}

// UpdateSeedValue 更新种子值
func (u *UserRepo) UpdateSeedValue(ctx context.Context, scene uint64, newSeed uint64) error {
	err := u.data.DB(ctx).Table("random_seeds").
		Where("scene = ?", scene).
		Update("seed_value", newSeed).
		Update("updated_at", time.Now().Format("2006-01-02 15:04:05")).Error

	if err != nil {
		return errors.New(500, "SEED UPDATE ERROR", err.Error())
	}

	return nil
}

// GetLandByUserIdLocationNum
func (u *UserRepo) GetLandByUserIdLocationNum(ctx context.Context, userId uint64, locationNum uint64) (*biz.Land, error) {
	var land Land

	if err := u.data.DB(ctx).Table("land").Where("user_id = ?", userId).Where("location_num = ?", locationNum).First(&land).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 没有找到返回 nil
		}
		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	return &biz.Land{
		ID:             land.ID,
		UserId:         land.UserId,
		Level:          land.Level,
		OutPutRate:     land.OutPutRate,
		RentOutPutRate: land.RentOutPutRate,
		MaxHealth:      land.MaxHealth,
		PerHealth:      land.PerHealth,
		LimitDate:      land.LimitDate,
		Status:         land.Status,
		LocationNum:    land.LocationNum,
		CreatedAt:      land.CreatedAt,
		UpdatedAt:      land.UpdatedAt,
		One:            land.One,
		Two:            land.Two,
		Three:          land.Three,
		SellAmount:     land.SellAmount,
	}, nil
}

// GetLandByIDTwo
func (u *UserRepo) GetLandByIDTwo(ctx context.Context, landID uint64) (*biz.Land, error) {
	var land Land

	if err := u.data.DB(ctx).Table("land").Where("id = ?", landID).First(&land).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 没有找到返回 nil
		}
		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	return &biz.Land{
		ID:             land.ID,
		UserId:         land.UserId,
		Level:          land.Level,
		OutPutRate:     land.OutPutRate,
		RentOutPutRate: land.RentOutPutRate,
		MaxHealth:      land.MaxHealth,
		PerHealth:      land.PerHealth,
		LimitDate:      land.LimitDate,
		Status:         land.Status,
		LocationNum:    land.LocationNum,
		CreatedAt:      land.CreatedAt,
		UpdatedAt:      land.UpdatedAt,
		One:            land.One,
		Two:            land.Two,
		Three:          land.Three,
		SellAmount:     land.SellAmount,
	}, nil
}

// GetLandByID 根据 Land ID 查询单条 Land 记录
func (u *UserRepo) GetLandByID(ctx context.Context, landID uint64) (*biz.Land, error) {
	var land Land

	if err := u.data.DB(ctx).Table("land").Where("limit_date>=?", time.Now().Unix()).Where("id = ?", landID).First(&land).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 没有找到返回 nil
		}
		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	return &biz.Land{
		ID:             land.ID,
		UserId:         land.UserId,
		Level:          land.Level,
		OutPutRate:     land.OutPutRate,
		RentOutPutRate: land.RentOutPutRate,
		MaxHealth:      land.MaxHealth,
		PerHealth:      land.PerHealth,
		LimitDate:      land.LimitDate,
		Status:         land.Status,
		LocationNum:    land.LocationNum,
		CreatedAt:      land.CreatedAt,
		UpdatedAt:      land.UpdatedAt,
		One:            land.One,
		Two:            land.Two,
		Three:          land.Three,
		SellAmount:     land.SellAmount,
	}, nil
}

// GetPropByIDSell 根据道具 ID 查询单条道具数据
func (u *UserRepo) GetPropByIDSell(ctx context.Context, propID, status uint64) (*biz.Prop, error) {
	var prop Prop

	if err := u.data.DB(ctx).Table("prop").
		Where("id = ?", propID).
		Where("status <=?", status).
		First(&prop).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 代表未找到
		}
		return nil, errors.New(500, "PROP ERROR", err.Error())
	}

	return &biz.Prop{
		ID:         prop.ID,
		UserId:     prop.UserId,
		Status:     prop.Status,
		PropType:   prop.PropType,
		OneOne:     prop.OneOne,
		OneTwo:     prop.OneTwo,
		TwoOne:     prop.TwoOne,
		TwoTwo:     prop.TwoTwo,
		ThreeOne:   prop.ThreeOne,
		FourOne:    prop.FourOne,
		FiveOne:    prop.FiveOne,
		SellAmount: prop.SellAmount,
		CreatedAt:  prop.CreatedAt,
		UpdatedAt:  prop.UpdatedAt,
	}, nil
}

// GetPropByID 根据道具 ID 查询单条道具数据
func (u *UserRepo) GetPropByID(ctx context.Context, propID, status uint64) (*biz.Prop, error) {
	var prop Prop

	if err := u.data.DB(ctx).Table("prop").
		Where("id = ?", propID).
		Where("status = ?", status).
		First(&prop).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 代表未找到
		}
		return nil, errors.New(500, "PROP ERROR", err.Error())
	}

	return &biz.Prop{
		ID:         prop.ID,
		UserId:     prop.UserId,
		Status:     prop.Status,
		PropType:   prop.PropType,
		OneOne:     prop.OneOne,
		OneTwo:     prop.OneTwo,
		TwoOne:     prop.TwoOne,
		TwoTwo:     prop.TwoTwo,
		ThreeOne:   prop.ThreeOne,
		FourOne:    prop.FourOne,
		FiveOne:    prop.FiveOne,
		SellAmount: prop.SellAmount,
		CreatedAt:  prop.CreatedAt,
		UpdatedAt:  prop.UpdatedAt,
	}, nil
}

// GetPropByIDTwo 根据道具 ID 查询单条道具数据
func (u *UserRepo) GetPropByIDTwo(ctx context.Context, propID uint64) (*biz.Prop, error) {
	var prop Prop

	if err := u.data.DB(ctx).Table("prop").
		Where("id = ?", propID).
		First(&prop).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 代表未找到
		}
		return nil, errors.New(500, "PROP ERROR", err.Error())
	}

	return &biz.Prop{
		ID:         prop.ID,
		UserId:     prop.UserId,
		Status:     prop.Status,
		PropType:   prop.PropType,
		OneOne:     prop.OneOne,
		OneTwo:     prop.OneTwo,
		TwoOne:     prop.TwoOne,
		TwoTwo:     prop.TwoTwo,
		ThreeOne:   prop.ThreeOne,
		FourOne:    prop.FourOne,
		FiveOne:    prop.FiveOne,
		SellAmount: prop.SellAmount,
		CreatedAt:  prop.CreatedAt,
		UpdatedAt:  prop.UpdatedAt,
	}, nil
}

// BuySeed .
func (u *UserRepo) BuySeed(ctx context.Context, git, getGit float64, userId, userIdGet, seedId uint64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userIdGet).Where("git>=?", git).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", git), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", getGit), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("seed").Where("id=?", seedId).Where("user_id=?", userId).Where("status=?", 4).
		Updates(map[string]interface{}{"user_id": userIdGet, "status": 0, "sell_amount": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuySeed", "用户信息修改失败")
	}
	return nil
}

// SellSeed .
func (u *UserRepo) SellSeed(ctx context.Context, seedId, userId uint64, sellAmount float64) error {
	res := u.data.DB(ctx).Table("seed").Where("id=?", seedId).Where("user_id=?", userId).Where("status=?", 0).
		Updates(map[string]interface{}{"status": 4, "sell_amount": sellAmount, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellSeed", "用户信息修改失败")
	}
	return nil
}

// UnSellSeed .
func (u *UserRepo) UnSellSeed(ctx context.Context, seedId, userId uint64) error {
	res := u.data.DB(ctx).Table("seed").Where("id=?", seedId).Where("user_id=?", userId).Where("status=?", 4).
		Updates(map[string]interface{}{"status": 0, "sell_amount": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellSeed", "用户信息修改失败")
	}
	return nil
}

// BuyProp .
func (u *UserRepo) BuyProp(ctx context.Context, git, getGit float64, userId, userIdGet, propId uint64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userIdGet).Where("git>=?", git).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", git), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", getGit), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", propId).Where("user_id=?", userId).Where("status=?", 4).
		Updates(map[string]interface{}{"user_id": userIdGet, "status": 1, "sell_amount": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyProp", "用户信息修改失败")
	}
	return nil
}

// UnSellProp .
func (u *UserRepo) UnSellProp(ctx context.Context, propId uint64, userId uint64) error {
	res := u.data.DB(ctx).Table("prop").Where("id=?", propId).Where("user_id=?", userId).Where("status=?", 4).
		Updates(map[string]interface{}{"status": 1, "sell_amount": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellProp", "用户信息修改失败")
	}
	return nil
}

// SellProp .
func (u *UserRepo) SellProp(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error {
	res := u.data.DB(ctx).Table("prop").Where("id=?", propId).Where("user_id=?", userId).Where("status<=?", 2).
		Updates(map[string]interface{}{"status": 4, "sell_amount": sellAmount, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellProp", "用户信息修改失败")
	}
	return nil
}

// BuyLand .
func (u *UserRepo) BuyLand(ctx context.Context, git, getGit float64, userId, userIdGet, landId uint64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userIdGet).Where("git>=?", git).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", git), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", getGit), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("land").Where("id=?", landId).Where("user_id=?", userId).Where("status=?", 4).Where("three=?", 1).Where("limit_date>=?", time.Now().Unix()).
		Updates(map[string]interface{}{"user_id": userIdGet, "status": 0, "sell_amount": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyLand", "用户信息修改失败")
	}

	return nil
}

// UnSellLand .
func (u *UserRepo) UnSellLand(ctx context.Context, propId uint64, userId uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", propId).Where("user_id=?", userId).Where("status=?", 4).
		Updates(map[string]interface{}{"status": 0, "sell_amount": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellLand", "用户信息修改失败")
	}
	return nil
}

// SellLand .
func (u *UserRepo) SellLand(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", propId).Where("user_id=?", userId).Where("status=?", 0).
		Updates(map[string]interface{}{"status": 4, "sell_amount": sellAmount, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellLand", "用户信息修改失败")
	}
	return nil
}

// RentLand .
func (u *UserRepo) RentLand(ctx context.Context, landId uint64, userId uint64, rentRate float64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("user_id=?", userId).Where("status=?", 1).
		Updates(map[string]interface{}{"status": 3, "rent_out_put_rate": rentRate, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "RentLand", "用户信息修改失败")
	}
	return nil
}

// UnRentLand .
func (u *UserRepo) UnRentLand(ctx context.Context, landId uint64, userId uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("user_id=?", userId).Where("status=?", 3).
		Updates(map[string]interface{}{"status": 1, "rent_out_put_rate": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "RentLand", "用户信息修改失败")
	}
	return nil
}

// LandPull .
func (u *UserRepo) LandPull(ctx context.Context, landId uint64, userId uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("user_id=?", userId).Where("status=?", 1).
		Updates(map[string]interface{}{"status": 0, "location_num": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "LandPull", "用户信息修改失败")
	}
	return nil
}

// LandPush .
func (u *UserRepo) LandPush(ctx context.Context, landId uint64, userId, locationNum uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("user_id=?", userId).Where("status=?", 0).Where("location_num=?", 0).
		Updates(map[string]interface{}{"status": 1, "location_num": locationNum, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "LandPush", "用户信息修改失败")
	}
	return nil
}

// stakeGit .
func (u *UserRepo) stakeGit(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error {
	res := u.data.DB(ctx).Table("stake_git").Where("user_id=?", userId).Where("status=?", 1).
		Updates(map[string]interface{}{"status": 4, "sell_amount": sellAmount, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellLand", "用户信息修改失败")
	}

	return nil
}

// Plant .
func (u *UserRepo) Plant(ctx context.Context, status, originStatus, perHealth uint64, landUserUse *biz.LandUserUse) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landUserUse.LandId).Where("status=?", originStatus).Where("limit_date>=?", time.Now().Unix()).Where("max_health>=?", perHealth).
		Updates(map[string]interface{}{
			"status":     status,
			"max_health": gorm.Expr("max_health - ?", perHealth),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "plant", "更新记录失败")
	}

	res = u.data.DB(ctx).Table("seed").Where("id=?", landUserUse.SeedId).Where("status=?", 0).
		Updates(map[string]interface{}{
			"status":     1,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "plant", "更新失败")
	}

	res = u.data.DB(ctx).Table("land_user_use").Create(&LandUserUse{
		LandId:       landUserUse.LandId,
		Level:        landUserUse.Level,
		UserId:       landUserUse.UserId,
		OwnerUserId:  landUserUse.OwnerUserId,
		SeedId:       landUserUse.SeedId,
		SeedTypeId:   landUserUse.SeedTypeId,
		Status:       landUserUse.Status,
		BeginTime:    landUserUse.BeginTime,
		TotalTime:    landUserUse.TotalTime,
		OverTime:     landUserUse.OverTime,
		OutMaxNum:    landUserUse.OutMaxNum,
		OutNum:       landUserUse.OutNum,
		InsectStatus: landUserUse.InsectStatus,
		OutSubNum:    landUserUse.OutSubNum,
		StealNum:     landUserUse.StealNum,
		StopStatus:   landUserUse.StopStatus,
		StopTime:     landUserUse.StopTime,
		SubTime:      landUserUse.SubTime,
		UseChan:      landUserUse.UseChan,
		One:          landUserUse.One,
		Two:          landUserUse.Two,
	})
	if res.Error != nil {
		return errors.New(500, "OpenBox", "创建土地使用记录失败")
	}

	return nil
}

// PlantPlatTwo .
func (u *UserRepo) PlantPlatTwo(ctx context.Context, id, landId uint64, rent bool) error {
	if rent {
		res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("status=?", 8).
			Updates(map[string]interface{}{"status": 3, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
		if res.Error != nil {
			return errors.New(500, "sellLand", "用户信息修改失败")
		}

	} else {
		res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("status=?", 2).
			Updates(map[string]interface{}{"status": 1, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
		if res.Error != nil {
			return errors.New(500, "sellLand", "用户信息修改失败")
		}
	}

	res := u.data.DB(ctx).Table("land_user_use").Where("id=?", id).Where("status=?", 1).
		Updates(map[string]interface{}{"status": 2, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellLand", "用户信息修改失败")
	}

	return nil
}

// PlantPlatThree .
func (u *UserRepo) PlantPlatThree(ctx context.Context, id, overTime, propId uint64, one, two bool) error {
	updateColums := map[string]interface{}{
		"over_time":  overTime,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	if one {
		updateColums["one"] = uint64(0)
	}

	if two {
		updateColums["two"] = uint64(0)
	}

	res := u.data.DB(ctx).Table("land_user_use").Where("id=?", id).Where("status=?", 1).
		Updates(updateColums)
	if res.Error != nil {
		return errors.New(500, "PlantPlatThree", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", propId).Where("status=?", 1).
		Updates(map[string]interface{}{"status": 3, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "PlantPlatThree", "用户信息修改失败")
	}

	return nil
}

// PlantPlatFour .
func (u *UserRepo) PlantPlatFour(ctx context.Context, outMax float64, id, propId, propStatus, propNum uint64) error {
	updateColums := map[string]interface{}{
		"out_max_num": outMax,
		"two":         0,
		"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
	}

	res := u.data.DB(ctx).Table("land_user_use").Where("id=?", id).Where("status=?", 1).Where("two>?", 0).
		Updates(updateColums)
	if res.Error != nil {
		return errors.New(500, "PlantPlatFour", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", propId).
		Updates(map[string]interface{}{"status": propStatus, "four_one": propNum, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "PlantPlatFour", "用户信息修改失败")
	}

	return nil
}

// PlantPlatFive .
func (u *UserRepo) PlantPlatFive(ctx context.Context, overTime, id, propId, propStatus, propNum uint64) error {
	updateColums := map[string]interface{}{
		"over_time":  overTime,
		"one":        0,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	res := u.data.DB(ctx).Table("land_user_use").Where("id=?", id).Where("status=?", 1).Where("one>?", 0).
		Updates(updateColums)
	if res.Error != nil {
		return errors.New(500, "PlantPlatFive", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", propId).
		Updates(map[string]interface{}{"status": propStatus, "three_one": propNum, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "PlantPlatFive", "用户信息修改失败")
	}

	return nil
}

// PlantPlatSix .
func (u *UserRepo) PlantPlatSix(ctx context.Context, id, propId, propStatus, propNum, landId uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("status=?", 8).
		Updates(map[string]interface{}{"status": 3, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "sellLand", "用户信息修改失败")
	}

	updateColums := map[string]interface{}{
		"status":     2,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	res = u.data.DB(ctx).Table("land_user_use").Where("id=?", id).Where("status=?", 1).
		Updates(updateColums)
	if res.Error != nil {
		return errors.New(500, "PlantPlatSix", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", propId).
		Updates(map[string]interface{}{"status": propStatus, "two_one": propNum, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "PlantPlatSix", "用户信息修改失败")
	}

	return nil
}

// PlantPlatSeven .
func (u *UserRepo) PlantPlatSeven(ctx context.Context, outMax, amount float64, subTime, lastTime, id, propId, propStatus, propNum, userId uint64) error {
	updateColums := map[string]interface{}{
		"sub_time":    subTime,
		"out_max_num": outMax,
		"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
	}

	res := u.data.DB(ctx).Table("land_user_use").Where("id=?", id).Where("status=?", 1).Where("sub_time=?", lastTime).
		Updates(updateColums)
	if res.Error != nil {
		return errors.New(500, "PlantPlatFive", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", propId).
		Updates(map[string]interface{}{"status": propStatus, "five_one": propNum, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "PlantPlatFive", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "PlantPlatFive", "用户信息修改失败")
	}

	var reward Reward

	reward.Reason = 13
	reward.UserId = userId
	reward.Amount = amount
	reward.Two = id
	res = u.data.DB(ctx).Table("reward").Create(&reward)
	if res.Error != nil {
		return errors.New(500, "PlantPlatTwoTwo", "用户信息修改失败")
	}

	return nil
}

// PlantPlatTwoTwo .
func (u *UserRepo) PlantPlatTwoTwo(ctx context.Context, id, userId, rentUserId uint64, amount, rentAmount float64) error {
	if amount > 0 {
		res := u.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
		if res.Error != nil {
			return errors.New(500, "BuyBox", "用户信息修改失败")
		}

		var reward Reward

		reward.Reason = 1
		reward.UserId = userId
		reward.Amount = amount
		reward.Two = id
		res = u.data.DB(ctx).Table("reward").Create(&reward)
		if res.Error != nil {
			return errors.New(500, "PlantPlatTwoTwo", "用户信息修改失败")
		}

	}

	if rentUserId > 0 && rentAmount > 0 {
		res := u.data.DB(ctx).Table("user").Where("id=?", rentUserId).
			Updates(map[string]interface{}{"git": gorm.Expr("git + ?", rentAmount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
		if res.Error != nil {
			return errors.New(500, "PlantPlatTwoTwo", "用户信息修改失败")
		}

		var reward Reward

		reward.Reason = 2
		reward.UserId = rentUserId
		reward.Amount = rentAmount
		reward.Two = id
		res = u.data.DB(ctx).Table("reward").Create(&reward)
		if res.Error != nil {
			return errors.New(500, "PlantPlatTwoTwo", "用户信息修改失败")
		}
	}

	return nil
}

// PlantPlatTwoTwoL .
func (u *UserRepo) PlantPlatTwoTwoL(ctx context.Context, id, userId, lowUserId, num uint64, amount float64) error {
	if amount > 0 {
		if 4 == num {
			res := u.data.DB(ctx).Table("user").Where("id=?", userId).
				Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "reward_one": gorm.Expr("reward_one + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
			if res.Error != nil {
				return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
			}
		} else if 5 == num {
			res := u.data.DB(ctx).Table("user").Where("id=?", userId).
				Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "reward_two": gorm.Expr("reward_two + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
			if res.Error != nil {
				return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
			}
		} else if 7 == num {
			res := u.data.DB(ctx).Table("user").Where("id=?", userId).
				Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "reward_two_one": gorm.Expr("reward_two_one + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
			if res.Error != nil {
				return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
			}
		} else if 8 == num {
			res := u.data.DB(ctx).Table("user").Where("id=?", userId).
				Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "reward_two_two": gorm.Expr("reward_two_two + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
			if res.Error != nil {
				return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
			}
		} else if 10 == num {
			res := u.data.DB(ctx).Table("user").Where("id=?", userId).
				Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "reward_three_one": gorm.Expr("reward_three_one + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
			if res.Error != nil {
				return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
			}
		} else if 11 == num {
			res := u.data.DB(ctx).Table("user").Where("id=?", userId).
				Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "reward_three_two": gorm.Expr("reward_three_two + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
			if res.Error != nil {
				return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
			}
		}

		var reward Reward

		reward.Reason = num
		reward.UserId = userId
		reward.Amount = amount
		reward.One = lowUserId
		reward.Two = id
		res := u.data.DB(ctx).Table("reward").Create(&reward)
		if res.Error != nil {
			return errors.New(500, "PlantPlatTwoTwoL", "用户信息修改失败")
		}
	}

	return nil
}

// LandAddOutRate .
func (u *UserRepo) LandAddOutRate(ctx context.Context, id, landId, userId uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", landId).Where("user_id=?", userId).
		Updates(map[string]interface{}{"max_health": gorm.Expr("max_health + ?", 20), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "LandAddOutRate", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("prop").Where("id=?", id).Where("user_id=?", userId).
		Updates(map[string]interface{}{"status": 3, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "LandAddOutRate", "用户信息修改失败")
	}

	return nil
}

// GetLand .
func (u *UserRepo) GetLand(ctx context.Context, id, id2, userId uint64) error {
	res := u.data.DB(ctx).Table("land").Where("id=?", id).Where("user_id=?", userId).Where("status=?", 0).
		Updates(map[string]interface{}{"status": 10, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "GetLand", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("land").Where("id=?", id2).Where("user_id=?", userId).Where("status=?", 0).
		Updates(map[string]interface{}{"status": 10, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "GetLand", "用户信息修改失败")
	}

	return nil
}

// CreateLand 创建一条 Land 记录
func (u *UserRepo) CreateLand(ctx context.Context, lc *biz.Land) (*biz.Land, error) {
	var land Land
	land.UserId = lc.UserId
	land.Level = lc.Level
	land.OutPutRate = lc.OutPutRate
	land.RentOutPutRate = lc.RentOutPutRate
	land.MaxHealth = lc.MaxHealth
	land.PerHealth = lc.PerHealth
	land.LimitDate = lc.LimitDate
	land.Status = lc.Status
	land.LocationNum = lc.LocationNum
	land.One = lc.One
	land.Two = lc.Two
	land.Three = lc.Three
	land.SellAmount = lc.SellAmount

	res := u.data.DB(ctx).Table("land").Create(&land)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_LAND_ERROR", "土地创建失败")
	}

	return &biz.Land{
		ID:             land.ID,
		UserId:         land.UserId,
		Level:          land.Level,
		OutPutRate:     land.OutPutRate,
		RentOutPutRate: land.RentOutPutRate,
		MaxHealth:      land.MaxHealth,
		PerHealth:      land.PerHealth,
		LimitDate:      land.LimitDate,
		Status:         land.Status,
		LocationNum:    land.LocationNum,
		One:            land.One,
		Two:            land.Two,
		Three:          land.Three,
		SellAmount:     land.SellAmount,
		CreatedAt:      land.CreatedAt,
		UpdatedAt:      land.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetLandInfoByLevels(ctx context.Context) (map[uint64]*biz.LandInfo, error) {
	var landInfos []*LandInfo

	res := make(map[uint64]*biz.LandInfo, 0)
	if err := u.data.DB(ctx).Table("land_info").Find(&landInfos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "LAND_INFO_ERROR", err.Error())
	}

	for _, landInfo := range landInfos {
		res[landInfo.Level] = &biz.LandInfo{
			ID:                landInfo.ID,
			Level:             landInfo.Level,
			OutPutRateMax:     landInfo.OutPutRateMax,
			OutPutRateMin:     landInfo.OutPutRateMin,
			RentOutPutRateMax: landInfo.RentOutPutRateMax,
			MaxHealth:         landInfo.MaxHealth,
			PerHealth:         landInfo.PerHealth,
			LimitDateMax:      landInfo.LimitDateMax,
			CreatedAt:         landInfo.CreatedAt,
			UpdatedAt:         landInfo.UpdatedAt,
		}
	}
	return res, nil
}

// SetGiw .
func (u *UserRepo) SetGiw(ctx context.Context, address string, giw uint64) error {
	res := u.data.DB(ctx).Table("user").Where("address=?", address).
		Updates(map[string]interface{}{"giw": gorm.Expr("giw + ?", float64(giw)), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	return nil
}

// SetGit .
func (u *UserRepo) SetGit(ctx context.Context, address string, git uint64) error {
	res := u.data.DB(ctx).Table("user").Where("address=?", address).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", float64(git)), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "BuyBox", "用户信息修改失败")
	}

	return nil
}

// SetStakeGetTotal .
func (u *UserRepo) SetStakeGetTotal(ctx context.Context, amount, balance float64) error {
	res := u.data.DB(ctx).Table("stake_get_total").Where("id=?", 1).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount + ?", amount),
			"balance":    gorm.Expr("balance + ?", balance),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetTotal", "用户信息修改失败")
	}

	return nil
}

// SetStakeGetTotalSub .
func (u *UserRepo) SetStakeGetTotalSub(ctx context.Context, amount, balance float64) error {
	res := u.data.DB(ctx).Table("stake_get_total").Where("id=?", 1).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount - ?", amount),
			"balance":    gorm.Expr("balance- ?", balance),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetTotal", "用户信息修改失败")
	}

	return nil
}

// SetStakeGit .
func (u *UserRepo) SetStakeGit(ctx context.Context, userId uint64, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("git>=?", amount).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	var stakeRecord StakeGitRecord

	stakeRecord.Amount = amount
	stakeRecord.UserId = userId
	stakeRecord.StakeType = 1

	res = u.data.DB(ctx).Table("stake_git_record").Create(&stakeRecord)
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "创建质押记录失败")
	}

	return nil
}

// SetUnStakeGit .
func (u *UserRepo) SetUnStakeGit(ctx context.Context, id, userId uint64, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetUnStakeGet", "用户信息修改失败")
	}
	res = u.data.DB(ctx).Table("stake_git_record").Where("id=?", id).
		Updates(map[string]interface{}{"stake_type": 2, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetUnStakeGet", "用户信息修改失败")
	}

	return nil
}

// SetStakeGet .
func (u *UserRepo) SetStakeGet(ctx context.Context, userId uint64, git, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("git>=?", amount).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", git), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("stake_get").Where("user_id=?", userId).
		Updates(map[string]interface{}{
			"stake_rate": gorm.Expr("stake_rate + ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	return nil
}

// SetStakeGetSub .
func (u *UserRepo) SetStakeGetSub(ctx context.Context, userId uint64, git, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", git), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("stake_get").Where("user_id=?", userId).
		Updates(map[string]interface{}{
			"stake_rate": gorm.Expr("stake_rate - ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	return nil
}

// SetStakeGetPlaySub .
func (u *UserRepo) SetStakeGetPlaySub(ctx context.Context, userId uint64, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("git>=?", amount).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", amount), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("stake_get_total").Where("id=?", 1).
		Updates(map[string]interface{}{
			"balance":    gorm.Expr("balance + ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "用户信息修改失败")
	}

	var stakeRecord StakeGetPlayRecord

	stakeRecord.Amount = amount
	stakeRecord.Reward = 0
	stakeRecord.UserId = userId
	stakeRecord.Status = 2

	res = u.data.DB(ctx).Table("stake_get_play_record").Create(&stakeRecord)
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "创建质押记录失败")
	}
	return nil
}

// SetStakeGetPlay .
func (u *UserRepo) SetStakeGetPlay(ctx context.Context, userId uint64, git, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"git": gorm.Expr("git + ?", git), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "用户信息修改失败")
	}

	res = u.data.DB(ctx).Table("stake_get_total").Where("id=?", 1).
		Updates(map[string]interface{}{
			"balance":    gorm.Expr("balance - ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "用户信息修改失败")
	}

	var stakeRecord StakeGetPlayRecord

	stakeRecord.Amount = amount
	stakeRecord.Reward = git
	stakeRecord.UserId = userId
	stakeRecord.Status = 1

	res = u.data.DB(ctx).Table("stake_get_play_record").Create(&stakeRecord)
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "创建质押记录失败")
	}
	return nil
}

func (u *UserRepo) CreateStakeGet(ctx context.Context, sg *biz.StakeGet) error {
	var stakeGet StakeGet
	stakeGet.UserId = sg.UserId
	stakeGet.Status = 2
	stakeGet.StakeRate = 0

	res := u.data.DB(ctx).Table("stake_get").Create(&stakeGet)
	if res.Error != nil {
		return errors.New(500, "CREATE_STAKE_GET_ERROR", "创建质押记录失败")
	}

	return nil
}

// Exchange .
func (u *UserRepo) Exchange(ctx context.Context, userId uint64, git, giw float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("git>=?", git).
		Updates(map[string]interface{}{"git": gorm.Expr("git - ?", git), "giw": gorm.Expr("giw + ?", giw), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	return nil
}

// Withdraw .
func (u *UserRepo) Withdraw(ctx context.Context, userId uint64, giw float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("giw>=?", giw).
		Updates(map[string]interface{}{"giw": gorm.Expr("giw - ?", giw), "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	if res.Error != nil {
		return errors.New(500, "SetStakeGet", "用户信息修改失败")
	}

	var withdraw Withdraw

	withdraw.UserId = userId
	withdraw.Amount = uint64(giw)
	withdraw.RelAmount = uint64(giw)
	withdraw.Status = "rewarded"
	res = u.data.DB(ctx).Table("withdraw").Create(&withdraw)
	if res.Error != nil {
		return errors.New(500, "CREATE_STAKE_GET_ERROR", "创建提现记录失败")
	}

	return nil
}

// GetBuyLandById 获取指定 ID 的 BuyLand 记录
func (u *UserRepo) GetBuyLandById(ctx context.Context) (*biz.BuyLand, error) {
	var buyLand *BuyLand
	if err := u.data.DB(ctx).Where("status=?", 1).Order("id desc").Table("buy_land").First(&buyLand).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "BUY LAND ERROR", err.Error())
	}

	return &biz.BuyLand{
		ID:        buyLand.ID,
		Amount:    buyLand.Amount,
		Status:    buyLand.Status,
		CreatedAt: buyLand.CreatedAt,
		UpdatedAt: buyLand.UpdatedAt,
		AmountTwo: buyLand.AmountTwo,
		Limit:     buyLand.Limit,
	}, nil
}

func (u *UserRepo) CreateBuyLandRecord(ctx context.Context, limit uint64, bl *biz.BuyLandRecord) error {
	res := u.data.DB(ctx).Table("buy_land").Where("id=?", bl.BuyLandID).
		Updates(map[string]interface{}{
			"limit":      limit,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "SetStakeGetPlaySub", "用户信息修改失败")
	}

	var buyLandRecord BuyLandRecord
	buyLandRecord.BuyLandID = bl.BuyLandID
	buyLandRecord.UserID = bl.UserID
	buyLandRecord.Amount = bl.Amount
	buyLandRecord.Status = bl.Status

	res = u.data.DB(ctx).Table("buy_land_record").Create(&buyLandRecord)
	if res.Error != nil {
		return errors.New(500, "CREATE BUY LAND RECORD ERROR", "创建拍卖记录失败")
	}

	return nil
}

// GetAllBuyLandRecords 获取所有买地记录
func (u *UserRepo) GetAllBuyLandRecords(ctx context.Context, id uint64) ([]*biz.BuyLandRecord, error) {
	var records []*BuyLandRecord

	res := make([]*biz.BuyLandRecord, 0)
	if err := u.data.DB(ctx).Table("buy_land_record").Where("buy_land_id", id).Order("amount desc").Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "BUY LAND RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.BuyLandRecord{
			ID:        record.ID,
			BuyLandID: record.BuyLandID,
			Amount:    record.Amount,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
			Status:    record.Status,
			UserID:    record.UserID,
		})
	}

	return res, nil
}
