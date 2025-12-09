package service

import (
	"context"
	"gf-demo-user/internal/model"
)

type IJWT interface {
	Generate(ctx context.Context, user *model.ContextUser) (string, error)
	Parse(ctx context.Context, token string) (*model.ContextUser, error)
}

var localJWT IJWT

func JWT() IJWT {
	if localJWT == nil {
		panic("implement not found for interface IJWT, forgot register?")
	}
	return localJWT
}

func RegisterJWT(i IJWT) {
	localJWT = i
}
