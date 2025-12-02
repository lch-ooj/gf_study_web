package user

import (
	"context"

	"gf-demo-user/api/user/v1"
	"gf-demo-user/internal/model"
	"gf-demo-user/internal/service"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
