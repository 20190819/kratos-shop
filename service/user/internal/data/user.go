package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}

func (ur *userRepo) ListUser(ctx context.Context, page, limit int) ([]*biz.User, int, error) {
	var users []*biz.User
	results := ur.data.db.Find(&users)
	if results.Error != nil {
		return nil, 0, results.Error
	}
	count := results.RowsAffected
	ur.data.db.Scopes(paginate(page, limit)).Find(&users)

	return users, int(count), nil
}

func paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	if page <= 0 {
		page = 1
	}

	if page > 100 {
		page = 100
	}

	if limit > 100 {
		limit = 100
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(page)
	}
}

func (ur *userRepo) UserByMobile(ctx context.Context, mobile string) (*biz.User, error) {

	var userInfo biz.User

	result := ur.data.db.Where(&biz.User{Mobile: mobile}).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &userInfo, nil
}

func (ur *userRepo) UserByID(ctx context.Context, Id int64) (*biz.User, error) {

	var userInfo biz.User
	result := ur.data.db.Where(&biz.User{ID: Id}).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	return &userInfo, nil
}

func (ur *userRepo) Update(ctx context.Context, u *biz.User) (bool, error) {
	var userInfo biz.User
	result := ur.data.db.Where(&biz.User{ID: u.ID}).First(&userInfo)
	if result.RowsAffected == 0 {
		return false, status.Errorf(codes.NotFound, "用户不存在")
	}

	userInfo.NickName = u.NickName
	userInfo.Birthday = u.Birthday
	userInfo.Gender = u.Gender
	res := ur.data.db.Save(&userInfo)
	if res.Error != nil {
		return false, status.Errorf(codes.Internal, res.Error.Error())
	}
	return true, nil
}

// CheckPassword 校验密码
func (ur *userRepo) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {

	// todo

	return true, nil
}
