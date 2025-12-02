package user

import (
	"context"

	"gf-demo-user/api/user/v1"
	"gf-demo-user/internal/service"
)

func (c *ControllerV1) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}
