package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-demo-user/internal/model/entity"
)

// SendEmailCodeReq 请求发送邮箱验证码
type SendEmailCodeReq struct {
	g.Meta  `path:"/user/email/send-code" method:"post" tags:"UserEmail" summary:"发送邮箱验证码"`
	Email   string `json:"email" v:"required|email"`
	Purpose string `json:"purpose" v:"required|in:register,login,reset"`
}

type SendEmailCodeRes struct {
	ExpiresIn int `json:"expiresIn" dc:"验证码有效期（分钟）"`
}

// EmailSignUpReq 邮箱验证码注册
type EmailSignUpReq struct {
	g.Meta   `path:"/user/email/sign-up" method:"post" tags:"UserEmail" summary:"邮箱验证码注册"`
	Email    string `json:"email" v:"required|email"`
	Password string `json:"password" v:"required|length:6,16"`
	Nickname string `json:"nickname"`
	Code     string `json:"code" v:"required|length:4,6"`
}

type EmailSignUpRes struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}

// EmailSignInReq 邮箱验证码登录
type EmailSignInReq struct {
	g.Meta `path:"/user/email/sign-in" method:"post" tags:"UserEmail" summary:"邮箱验证码登录"`
	Email  string `json:"email" v:"required|email"`
	Code   string `json:"code" v:"required|length:4,6"`
}

type EmailSignInRes struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}

// ResetPasswordReq 邮箱验证码重置密码
type ResetPasswordReq struct {
	g.Meta      `path:"/user/email/reset-password" method:"post" tags:"UserEmail" summary:"邮箱验证码重置密码"`
	Email       string `json:"email" v:"required|email"`
	Code        string `json:"code" v:"required|length:4,6"`
	NewPassword string `json:"newPassword" v:"required|length:6,16"`
}

type ResetPasswordRes struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}
