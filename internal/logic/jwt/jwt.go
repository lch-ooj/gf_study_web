package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"

	"gf-demo-user/internal/model"
	"gf-demo-user/internal/service"
)

type sJWT struct {
	secret string
	expire time.Duration
	initd  bool
}

type claims struct {
	User *model.ContextUser `json:"user"`
	jwt.RegisteredClaims
}

func init() {
	service.RegisterJWT(New())
}

func New() service.IJWT {
	return &sJWT{}
}

func (s *sJWT) init(ctx context.Context) {
	if s.initd {
		return
	}
	cfg := g.Cfg()
	secret := cfg.MustGet(ctx, "jwt.secret").String()
	if secret == "" {
		secret = "change-me"
	}
	hours := cfg.MustGet(ctx, "jwt.expireHours").Int()
	if hours == 0 {
		hours = 24
	}
	s.secret = secret
	s.expire = time.Duration(hours) * time.Hour
	s.initd = true
}

func (s *sJWT) Generate(ctx context.Context, user *model.ContextUser) (string, error) {
	s.init(ctx)
	if user == nil {
		return "", gerror.New("nil user for token generation")
	}
	now := time.Now()
	cl := claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(user.Id),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)
	return token.SignedString([]byte(s.secret))
}

func (s *sJWT) Parse(ctx context.Context, tokenString string) (*model.ContextUser, error) {
	s.init(ctx)
	parsed, err := jwt.ParseWithClaims(tokenString, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if c, ok := parsed.Claims.(*claims); ok && parsed.Valid {
		return c.User, nil
	}
	return nil, gerror.New("invalid token")
}
