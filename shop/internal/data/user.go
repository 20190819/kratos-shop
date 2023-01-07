package data

import (
	"context"

	"shop/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func(repo *userRepo)CreateUser(ctx context.Context, u *biz.User) (*biz.User, error){
	repo.
}
func(repo *userRepo)UserByMobile(ctx context.Context, mobile string) (*biz.User, error){

}
func(repo *userRepo)UserById(ctx context.Context, Id int64) (*biz.User, error){

}
func(repo *userRepo)CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error){

}