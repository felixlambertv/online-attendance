package service

import (
	"github.com/felixlambertv/online-attendance/model/entity"
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/felixlambertv/online-attendance/repository"
)

type IUserService interface {
	FindUser(userId uint) (entity.User, error)
	Register(request request.UserRegister) (entity.User, error)
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (u UserService) FindUser(userId uint) (entity.User, error) {
	return u.userRepository.GetUser(userId)
}

func (u UserService) Register(request request.UserRegister) (entity.User, error) {
	return u.userRepository.Save(entity.User{
		Name: request.Name,
	})
}
