package model

import "gf-demo-user/internal/model/entity"

type UserCreateInput struct {
	Passport string
	Password string
	Nickname string
	Email    string
}

type UserSignInInput struct {
	Passport string
	Password string
}

// Email-based flows
type EmailCodeInput struct {
	Email   string
	Purpose string
}

type UserEmailSignUpInput struct {
	Email    string
	Password string
	Nickname string
	Code     string
}

type UserEmailSignInInput struct {
	Email string
	Code  string
}

type UserResetPasswordInput struct {
	Email       string
	NewPassword string
	Code        string
}

type AuthOutput struct {
	Token string
	User  *entity.User
}
