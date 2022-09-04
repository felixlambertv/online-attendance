package controller

import (
	"errors"
	"github.com/felixlambertv/online-attendance/exception"
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/felixlambertv/online-attendance/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type IUserController interface {
	Detail(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{userService: userService}
}

func (u UserController) Detail(ctx *gin.Context) {
	var req request.UserDetail
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(exception.NewAppException("wrong user id"))
		return
	}
	user, err := u.userService.FindUser(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error(exception.NewDataNotFoundException("User"))
		} else {
			ctx.Error(errors.New("something wrong"))
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u UserController) Register(ctx *gin.Context) {
	var req request.UserRegister
	err := ctx.Bind(&req)
	if err != nil {
		logrus.Error("invalid request")
	}
	user, err := u.userService.Register(req)
	if err != nil {
		logrus.Error("fail create user")
	}

	ctx.JSON(http.StatusOK, user)
}
