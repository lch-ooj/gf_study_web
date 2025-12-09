package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	mail "github.com/jordan-wright/email"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"

	"gf-demo-user/internal/consts"
	"gf-demo-user/internal/dao"
	"gf-demo-user/internal/model"
	"gf-demo-user/internal/model/do"
	"gf-demo-user/internal/model/entity"
	"gf-demo-user/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (user *entity.User, err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	if in.Email == "" {
		in.Email = in.Passport
	}
	var (
		available bool
	)
	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	available, err = s.IsEmailAvailable(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Email "%s" is already token by others`, in.Email)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	err = dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, errInsert := dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
			Email:    in.Email,
		}).Insert()
		if errInsert != nil {
			return errInsert
		}
		id, _ := result.LastInsertId()
		user = &entity.User{
			Id:        uint(id),
			Passport:  in.Passport,
			Password:  in.Password,
			Nickname:  in.Nickname,
			Email:     in.Email,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}
		return nil
	})
	return
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (out *model.AuthOutput, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: in.Passport,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New(`Passport or Password not correct`)
	}
	return s.authSuccess(ctx, user)
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsEmailAvailable checks and returns given email is available for signing up.
func (s *sUser) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Email: email}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// SignInWithEmailCode logs a user in via email verification code.
func (s *sUser) SignInWithEmailCode(ctx context.Context, in model.UserEmailSignInInput) (out *model.AuthOutput, err error) {
	if err = s.validateEmailCode(ctx, in.Email, in.Code, consts.EmailPurposeLogin); err != nil {
		return nil, err
	}
	user, err := s.findUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在，请先注册")
	}
	return s.authSuccess(ctx, user)
}

// SignUpWithEmailCode registers user via email verification code.
func (s *sUser) SignUpWithEmailCode(ctx context.Context, in model.UserEmailSignUpInput) (out *model.AuthOutput, err error) {
	if err = s.validateEmailCode(ctx, in.Email, in.Code, consts.EmailPurposeRegister); err != nil {
		return nil, err
	}
	user, err := s.Create(ctx, model.UserCreateInput{
		Passport: in.Email,
		Password: in.Password,
		Nickname: strings.TrimSpace(in.Nickname),
		Email:    in.Email,
	})
	if err != nil {
		return nil, err
	}
	return s.authSuccess(ctx, user)
}

// SendEmailCode generates and emails a verification code.
func (s *sUser) SendEmailCode(ctx context.Context, in model.EmailCodeInput) (err error) {
	in.Email = strings.TrimSpace(in.Email)
	if in.Email == "" {
		return gerror.New("邮箱不能为空")
	}
	if in.Purpose == "" {
		in.Purpose = consts.EmailPurposeRegister
	}
	if in.Purpose != consts.EmailPurposeRegister && in.Purpose != consts.EmailPurposeLogin && in.Purpose != consts.EmailPurposeReset {
		return gerror.New("purpose 只能是 register/login/reset")
	}
	code := s.generateCode()
	expireMinutes := s.codeExpireMinutes(ctx)
	if err = s.storeEmailCode(ctx, in.Email, code, in.Purpose, expireMinutes); err != nil {
		return err
	}
	return s.sendCodeEmail(ctx, in.Email, code, in.Purpose, expireMinutes)
}

// ResetPassword resets password using email code.
func (s *sUser) ResetPassword(ctx context.Context, in model.UserResetPasswordInput) (out *model.AuthOutput, err error) {
	if err = s.validateEmailCode(ctx, in.Email, in.Code, consts.EmailPurposeReset); err != nil {
		return nil, err
	}
	_, err = dao.User.Ctx(ctx).Where(do.User{Email: in.Email}).Data(do.User{Password: in.NewPassword}).Update()
	if err != nil {
		return nil, err
	}
	user, err := s.findUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}
	return s.authSuccess(ctx, user)
}

// GetProfile retrieves and returns current user info from session or JWT context.
func (s *sUser) GetProfile(ctx context.Context) *entity.User {
	if u := service.Session().GetUser(ctx); u != nil {
		return u
	}
	if biz := service.BizCtx().Get(ctx); biz != nil && biz.User != nil {
		var user *entity.User
		if err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, biz.User.Id).Scan(&user); err == nil {
			return user
		}
	}
	return nil
}

// --------------------- helpers ---------------------

func (s *sUser) authSuccess(ctx context.Context, user *entity.User) (*model.AuthOutput, error) {
	if err := service.Session().SetUser(ctx, user); err != nil {
		return nil, err
	}
	ctxUser := s.buildContextUser(user)
	service.BizCtx().SetUser(ctx, ctxUser)
	token, err := service.JWT().Generate(ctx, ctxUser)
	if err != nil {
		return nil, err
	}
	return &model.AuthOutput{Token: token, User: user}, nil
}

func (s *sUser) buildContextUser(user *entity.User) *model.ContextUser {
	return &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
}

func (s *sUser) generateCode() string {
	return grand.Digits(4)
}

func (s *sUser) codeExpireMinutes(ctx context.Context) int {
	minutes := g.Cfg().MustGet(ctx, "email.codeExpireMinutes").Int()
	if minutes == 0 {
		minutes = 10
	}
	return minutes
}

func (s *sUser) storeEmailCode(ctx context.Context, email, code, purpose string, expireMinutes int) error {
	expires := gtime.Now().Add(time.Duration(expireMinutes) * time.Minute)
	_, err := g.Model("email_code").Ctx(ctx).Data(g.Map{
		"email":      email,
		"code":       code,
		"purpose":    purpose,
		"expires_at": expires,
		"used":       0,
	}).Insert()
	return err
}

func (s *sUser) validateEmailCode(ctx context.Context, email, code, purpose string) error {
	var record struct {
		Id        int
		Code      string
		Purpose   string
		Used      int
		ExpiresAt *gtime.Time
	}
	err := g.Model("email_code").Ctx(ctx).
		Where(g.Map{"email": email, "purpose": purpose}).
		OrderDesc("id").
		Limit(1).
		Scan(&record)
	if err != nil {
		if err == sql.ErrNoRows {
			return gerror.New("验证码不存在，请先发送")
		}
		return err
	}
	if record.Id == 0 {
		return gerror.New("验证码不存在，请先发送")
	}
	if record.Code != code {
		return gerror.New("验证码不正确")
	}
	if record.Used == 1 {
		return gerror.New("验证码已使用")
	}
	if record.ExpiresAt != nil && record.ExpiresAt.Before(gtime.Now()) {
		return gerror.New("验证码已过期")
	}
	_, err = g.Model("email_code").Ctx(ctx).Where("id", record.Id).Data(g.Map{"used": 1}).Update()
	return err
}

func (s *sUser) sendCodeEmail(ctx context.Context, to, code, purpose string, expireMinutes int) error {
	var cfg struct {
		From string
		SMTP struct {
			Host     string
			Port     int
			Username string
			Password string
		}
	}
	if err := g.Cfg().MustGet(ctx, "email").Struct(&cfg); err != nil {
		return err
	}
	subject := "您的验证码"
	switch purpose {
	case consts.EmailPurposeRegister:
		subject = "注册验证码"
	case consts.EmailPurposeLogin:
		subject = "登录验证码"
	case consts.EmailPurposeReset:
		subject = "重置密码验证码"
	}
	body := fmt.Sprintf("您正在进行%s操作，验证码：%s，有效期 %d 分钟。", purpose, code, expireMinutes)
	e := mail.NewEmail()
	e.From = cfg.From
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)
	addr := fmt.Sprintf("%s:%d", cfg.SMTP.Host, cfg.SMTP.Port)
	auth := smtp.PlainAuth("", cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host)
	return e.Send(addr, auth)
}

func (s *sUser) findUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(do.User{Email: email}).Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
