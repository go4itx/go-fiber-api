package service

import (
	conf "home/config"
	"home/pkg/code"
	"home/pkg/e"
	"home/pkg/utils/jwt"
)

type adminService struct {
	account Account
}

type Account struct {
	Name     string
	Password string
}

// newAdminService construct an instance
func newAdminService() *adminService {
	return &adminService{
		account: Account{
			Name:     conf.Viper.GetString("account.admin.name"),
			Password: conf.Viper.GetString("account.admin.password"),
		},
	}
}

type LoginReq struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

type LoginRes struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// Login service
func (s *adminService) Login(p LoginReq) (res LoginRes, err error) {
	if !(p.Name == s.account.Name && p.Password == s.account.Password) {
		err = e.NewError(code.LoginFailed)
		return
	}

	res.Token, res.Expire, err = jwt.CreateToken(jwt.User{
		ID:     1,
		RoleID: 1,
		Name:   p.Name,
	})

	return
}
