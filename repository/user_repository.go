package repository

import (
	"github.com/felixlambertv/online-attendance/model/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	GetUser(userId uint) (entity.User, error)
	WithTrx(db *gorm.DB) userRepository
}

func (u userRepository) Save(user entity.User) (entity.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetAllUsers() (users []entity.User, err error) {
	err = u.DB.Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users, err
}

func (u userRepository) GetUser(userId uint) (user entity.User, err error) {
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
