package user

import (
	"context"

	v1 "gf-demo-user/api/user/v1"
	"gf-demo-user/internal/model"
	"gf-demo-user/internal/service"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	_, err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	out, err := service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.SignUpRes{Token: out.Token}
	return
}
