package service

import (
	"github.com/felixlambertv/online-attendance/exception"
	"github.com/felixlambertv/online-attendance/model/entity"
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/felixlambertv/online-attendance/repository"
)

type IUserService interface {
	FindUser(userId uint) (entity.User, error)
	Register(request request.UserRequest) (entity.User, error)
	Delete(userId uint) error
	Update(userId uint, request request.UserRequest) (entity.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (u UserService) FindUser(userId uint) (entity.User, error) {
	return u.userRepository.FindById(userId)
}

func (u UserService) Register(request request.UserRequest) (entity.User, error) {
	return u.userRepository.Save(entity.User{
		Name: request.Name,
	})
}

func (u UserService) Delete(userId uint) (err error) {
	user, err := u.userRepository.FindById(userId)
	if err != nil || !user.DeletedAt.Time.IsZero() {
		return exception.NewAppException("User not found or already deleted")
	}

	return u.userRepository.Delete(userId)
}

func (u UserService) Update(userId uint, request request.UserRequest) (entity.User, error) {
	update, err := u.userRepository.Update(userId, request)
	if err != nil {
		return entity.User{}, err
	}
	return update, nil
}
