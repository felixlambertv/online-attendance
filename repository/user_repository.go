package repository

import (
	"github.com/felixlambertv/online-attendance/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

type UserRepository interface {
	Save(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUser(userId int) (model.User, error)
	WithTrx(db *gorm.DB) userRepository
}

func (u userRepository) Save(user model.User) (model.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetAllUsers() (users []model.User, err error) {
	err = u.DB.Find(&users).Error
	return users, err
}

func (u userRepository) GetUser(userId int) (user model.User, err error) {
	err = u.DB.First(&user, userId).Error
	return user, err
}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		return u
	}
	u.DB = trxHandle
	return u
}
