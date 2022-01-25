package service

import (
	"home/internal/demo/model"
	"home/pkg/code"
	"home/pkg/constant"
	"home/pkg/e"
	"home/pkg/utils/jwt"
	"home/pkg/utils/password"
)

type userService struct {
}

// newUserService construct an instance
func newUserService() *userService {
	return &userService{}
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
func (s *userService) Login(p LoginReq) (res LoginRes, err error) {
	user := model.User{Name: p.Name}
	if err = user.Info(); err != nil {
		err = e.NewError(code.LoginFailed)
		return
	}

	if !password.Verify(p.Password, user.Password) {
		err = e.NewError(code.LoginFailed)
		return
	}

	if user.Status != constant.StatusEnable {
		err = e.NewError(code.UserStatusDisable)
		return
	}

	res.Token, res.Expire, err = jwt.CreateToken(jwt.User{
		ID:     1,
		RoleID: 1,
		Name:   p.Name,
	})

	return
}
