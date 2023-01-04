package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday uint64
	Gender   string
	Role     int
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	ListUser(ctx context.Context, page, limit int) ([]*User, int, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	UserByID(ctx context.Context, Id int64) (*User, error)
	Update(ctx context.Context, u *User) (bool, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) (bool, error) {
	return uc.repo.Update(ctx, u)
}

func (uc *UserUseCase) List(ctx context.Context, page, limit int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, page, limit)
}

func (uc *UserUseCase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUseCase) UserByID(ctx context.Context, Id int64) (*User, error) {
	return uc.repo.UserByID(ctx, Id)
}

func (uc *UserUseCase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}
