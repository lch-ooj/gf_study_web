package user

import (
	"context"

	"gf-demo-user/api/user/v1"
	"gf-demo-user/internal/service"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}
