package service

import (
	"errors"
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type UserService struct {
	baseService
}

func NewUserService(session ...*session.Session) *UserService {
	service := new(UserService)
	service.setSession(session...)
	return service
}

func (s *UserService) Regist(user entity.User) (int, error) {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer s.session.Rollback()

	id, err := dao.NewUserDao(s.session).Insert(user)
	if err != nil {
		log.Println(err)
		return 0, nil
	}

	err = dao.NewUserSeqDao(s.session).Insert(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	s.session.Commit()
	return id, nil
}

var (
	ErrToken    = errors.New("error token")
	ErrPassword = errors.New("error password")
)

// SignIn 登录
func (s *UserService) SignIn(signIn entity.SignIn) error {
	token, err := dao.NewDeviceDao(s.session).GetToken(signIn.DeviceId)
	if err != nil {
		if err=
		log.Println(err)
	}

	if signIn.Token != token {
		return ErrToken
	}

	password, err := dao.NewUserDao(s.session).GetPassword(signIn.UserId)
	if err != nil {
		log.Println(err)
		return err
	}

	if signIn.Password != password {
		return ErrPassword
	}

	err = dao.NewDeviceDao(s.session).UpdateUserIdAndStatus(signIn.DeviceId, signIn.UserId, DeviceOnline)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
