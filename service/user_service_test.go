package service

import (
	"fmt"
	"github.com/felixlambertv/online-attendance/mocks"
	"github.com/felixlambertv/online-attendance/model/entity"
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

var userRepoMock = new(mocks.UserRepository)
var service = NewUserService(userRepoMock)
var dummyUser = entity.User{
	ID:        1,
	Name:      "Dummy",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
}

func TestMain(m *testing.M) {
	fmt.Println("before")
	userRepoMock.On("FindById", uint(1)).Return(dummyUser, nil)
	m.Run()
	fmt.Println("after")
}

func TestUserService_FindUser(t *testing.T) {
	user, err := service.FindUser(1)
	assert.Equal(t, user, entity.User{
		ID:        1,
		Name:      "Dummy",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	})
	assert.Equal(t, err, nil)
}

func TestUserService_Register(t *testing.T) {
	register, err := service.Register(request.UserRegister{Name: "Test"})
	if err != nil {
		return
	}
	fmt.Println(register)
}
