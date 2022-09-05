package repository

import (
	"github.com/felixlambertv/online-attendance/model/entity"
	"github.com/felixlambertv/online-attendance/model/request"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindById(userId uint) (entity.User, error)
	Delete(userId uint) error
	Update(userId uint, request request.UserRequest) (entity.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Update(userId uint, request request.UserRequest) (entity.User, error) {
	var user entity.User
	err := u.DB.First(&user, userId).Error
	if err != nil {
		return user, err
	}
	user.Name = request.Name
	u.DB.Save(&user)
	return user, nil
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{DB: db}
}

func (u UserRepository) Save(user entity.User) (entity.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

func (u UserRepository) FindAll() (users []entity.User, err error) {
	err = u.DB.Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users, err
}

func (u UserRepository) FindById(userId uint) (user entity.User, err error) {
	err = u.DB.First(&user, userId).Error
	return user, err
}

func (u UserRepository) Delete(userId uint) (err error) {
	return u.DB.Delete(&entity.User{}, userId).Error
}
