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
	FindAll() ([]entity.User, error)
	FindById(userId uint) (entity.User, error)
	DeleteUser(userId uint) error
}

func (u userRepository) Save(user entity.User) (entity.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) FindAll() (users []entity.User, err error) {
	err = u.DB.Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users, err
}

func (u userRepository) FindById(userId uint) (user entity.User, err error) {
	err = u.DB.First(&user, userId).Error
	return user, err
}

func (u userRepository) DeleteUser(userId uint) (err error) {
	return u.DB.Delete(&entity.User{}, userId).Error
}
