package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"wrapError/dao"
	"wrapError/models"
)

type userService struct {
	UserId   string
	UserInfo *models.User
}

func NewUserService(uid string) *userService {
	return &userService{
		UserId: uid,
	}
}

func (u *userService) load() (string, error) {
	if len(u.UserId) == 0 {
		return "", errors.New("id is nil")
	}
	return dao.FindById(u.UserId)
}

func (u *userService) GetUserName() (string, error) {
	name, err := u.load()
	if err != nil {
		return name, errors.WithMessage(err, "GetUserInfo")
	}
	return name, nil
}

func (u *userService) DeleteUser() error {
	_, err := u.load()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.WithMessage(err, "GetUserInfo")
	}
	// TODO 校验是否可以删除
	return nil
}
