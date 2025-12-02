package user

import (
	"context"

	"gf-demo-user/api/user/v1"
	"gf-demo-user/internal/model"
	"gf-demo-user/internal/service"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
