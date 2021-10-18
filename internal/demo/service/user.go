package service

import (
	"home/internal/demo/model"
	"home/pkg/code"
	"home/pkg/code/e"
	"home/pkg/constant"
	"home/pkg/utils/jwt"
	"home/pkg/utils/password"
)

type userService struct {
}

// newUserService construct an instance
func newUserService() *userService {
	return &userService{}
}

type ParamLogin struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

// Login service
func (s *userService) Login(p ParamLogin) (token string, err error) {
	user := model.User{Name: p.Name}
	err = db.Where(&user).First(&user).Error
	if err != nil {
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

	return jwt.CreateToken(jwt.User{
		ID:     1,
		RoleID: 1,
		Name:   p.Name,
	})
}
