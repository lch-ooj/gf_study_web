package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf-demo-user/api/user/v1"
	"gf-demo-user/internal/model"
	"gf-demo-user/internal/service"
)

// SendEmailCode 发送邮箱验证码
func (c *ControllerV1) SendEmailCode(ctx context.Context, req *v1.SendEmailCodeReq) (res *v1.SendEmailCodeRes, err error) {
	if err = service.User().SendEmailCode(ctx, model.EmailCodeInput{Email: req.Email, Purpose: req.Purpose}); err != nil {
		return nil, err
	}
	expires := g.Cfg().MustGet(ctx, "email.codeExpireMinutes").Int()
	if expires == 0 {
		expires = 10
	}
	res = &v1.SendEmailCodeRes{ExpiresIn: expires}
	return
}

// EmailSignUp 邮箱验证码注册
func (c *ControllerV1) EmailSignUp(ctx context.Context, req *v1.EmailSignUpReq) (res *v1.EmailSignUpRes, err error) {
	out, err := service.User().SignUpWithEmailCode(ctx, model.UserEmailSignUpInput{Email: req.Email, Password: req.Password, Nickname: req.Nickname, Code: req.Code})
	if err != nil {
		return nil, err
	}
	res = &v1.EmailSignUpRes{Token: out.Token, User: out.User}
	return
}

// EmailSignIn 邮箱验证码登录
func (c *ControllerV1) EmailSignIn(ctx context.Context, req *v1.EmailSignInReq) (res *v1.EmailSignInRes, err error) {
	out, err := service.User().SignInWithEmailCode(ctx, model.UserEmailSignInInput{Email: req.Email, Code: req.Code})
	if err != nil {
		return nil, err
	}
	res = &v1.EmailSignInRes{Token: out.Token, User: out.User}
	return
}

// ResetPassword 邮箱验证码重置密码
func (c *ControllerV1) ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error) {
	out, err := service.User().ResetPassword(ctx, model.UserResetPasswordInput{Email: req.Email, NewPassword: req.NewPassword, Code: req.Code})
	if err != nil {
		return nil, err
	}
	res = &v1.ResetPasswordRes{Token: out.Token, User: out.User}
	return
}
