package captcha

import (
	"context"
	"github.com/mojocn/base64Captcha"
)

type Info struct {
	CaptchaId string
	PicPath   string
}

func GetCaptcha(ctx context.Context) (*Info, error) {

	dv := base64Captcha.NewDriverDigit(80, 250, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(dv, base64Captcha.DefaultMemStore)
	id, path, err := captcha.Generate()
	if err != nil {
		return nil, err
	}
	return &Info{CaptchaId: id, PicPath: path}, nil
}
