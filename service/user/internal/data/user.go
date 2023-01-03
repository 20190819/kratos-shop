package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"user/internal/biz"
)

type User struct {
	ID        int64     `gorm:"primarykey"`
	Mobile    string    `gorm:"index:idx_mobile;unique;type:varchar(11);not null;comment:'手机号码，用户唯一标识'"`
	Password  string    `gorm:"type:varchar(100);not null"`
	NickName  string    `gorm:"type:varchar(25) comment '用户昵称'"`
	Birthday  time.Time `gorm:"type:date comment '出生年月日'"`
	Gender    string    `gorm:"default:male;type:varchar(16) comment 'female--女 male--男'"`
	Role      int       `gorm:"default:1;type:int comment '1--普通用户 2--管理员'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {

	return nil, nil
}

// NewUserRepo 飘红咋回事？
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}





