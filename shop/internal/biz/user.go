package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/mojocn/base64Captcha"
	pb "shop/api/shop/v1"
	"shop/internal/conf"
	"shop/internal/pkg/middleware/auth"
	"shop/internal/pkg/middleware/captcha"
	"time"

	jwt2 "github.com/golang-jwt/jwt/v4"
)

var (
	ErrPasswordInvalid     = errors.New("password invalid")
	ErrUsernameInvalid     = errors.New("username invalid")
	ErrCaptchaInvalid      = errors.New("captcha code err")
	ErrMobileInvalid       = errors.New("mobile invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrLoginFailed         = errors.New("login failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authorisation failed")
)

// User is a User model.
type User struct {
	ID        int64
	Mobile    string
	Nickname  string
	Birthday  int64
	Gender    string
	Role      int
	CreatedAt time.Time
	Password  string
}

// UserRepo is a User repo.
type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	UserById(ctx context.Context, Id int64) (*User, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

// UserUseCase is a User use case.
type UserUseCase struct {
	repo       UserRepo
	log        *log.Helper
	signingKey string // 这里为了生成 token 时直接取配置
}

// NewUserUseCase new a User usecase.
func NewUserUseCase(repo UserRepo, logger log.Logger, conf *conf.Auth) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger), signingKey: conf.JwtKey}
}

func (uc *UserUseCase) GetCaptcha(ctx context.Context) (*pb.CaptchaReply, error) {

	captchaInfo, err := captcha.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CaptchaReply{
		CaptchaId: captchaInfo.CaptchaId,
		PicPath:   captchaInfo.PicPath,
	}, nil

}

func (uc *UserUseCase) UserDetailByID(ctx context.Context) (*pb.DetailReply, error) {

	var uId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["ID"] == nil {
			return nil, ErrAuthFailed
		}
		uId = int64(c["ID"].(float64))
	}

	user, err := uc.repo.UserById(ctx, uId)
	if err != nil {
		return nil, err
	}
	return &pb.DetailReply{
		Id:       user.ID,
		NickName: user.Nickname,
		Mobile:   user.Mobile,
	}, nil
}

func (uc *UserUseCase) PasswordLogin(ctx context.Context, req *pb.LoginReq) (*pb.RegisterReply, error) {

	if len(req.Mobile) == 0 {
		return nil, ErrMobileInvalid
	}
	if len(req.Password) == 0 {
		return nil, ErrPasswordInvalid
	}
	if base64Captcha.Store.Verify(ctx, req.CaptchaId, req.Captcha, true) {
		return nil, ErrCaptchaInvalid
	}
	user, err := uc.repo.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}

	// 检查密码
	ok, err := uc.repo.CheckPassword(ctx, req.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrPasswordInvalid
	}

	claims := auth.CustomClaims{
		ID:          user.ID,
		NickName:    user.Nickname,
		AuthorityId: user.Role,
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 3600*24*30, // 有效期 30 天
			Issuer:    "Gyl",
		},
	}

	token, err := auth.CreateToken(claims, uc.signingKey)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Id:        user.ID,
		Mobile:    user.Mobile,
		Nickname:  user.Nickname,
		Token:     token,
		ExpiredAt: time.Now().Unix() + 3600*24*30,
	}, nil

}

func NewUser(mobile, nickname, password string) (*User, error) {
	if len(mobile) == 0 {
		return nil, ErrMobileInvalid
	}
	if len(nickname) == 0 {
		return nil, ErrUsernameInvalid
	}
	if len(password) == 0 {
		return nil, ErrPasswordInvalid
	}
	return &User{
		Mobile:   mobile,
		Nickname: nickname,
		Password: password,
	}, nil
}
