package service

import (
	"github.com/felixlambertv/online-attendance/exception"
	"github.com/felixlambertv/online-attendance/model/entity"
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/felixlambertv/online-attendance/repository"
)

type IUserService interface {
	FindUser(userId uint) (entity.User, error)
	Register(request request.UserRegister) (entity.User, error)
	Delete(userId uint) error
}

type UserService struct {
	userRepository repository.UserRepository
}

func (u UserService) Delete(userId uint) (err error) {
	user, err := u.userRepository.FindById(userId)
	if err != nil || !user.DeletedAt.Time.IsZero() {
		return exception.NewAppException("User not found or already deleted")
	}

	return u.userRepository.DeleteUser(userId)
}

func NewUserService(userRepository repository.UserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (u UserService) FindUser(userId uint) (entity.User, error) {
	return u.userRepository.FindById(userId)
}

func (u UserService) Register(request request.UserRegister) (entity.User, error) {
	return u.userRepository.Save(entity.User{
		Name: request.Name,
	})
}
