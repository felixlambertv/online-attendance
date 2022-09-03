package controller

import (
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/felixlambertv/online-attendance/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	err := ctx.BindUri(&req)
	if err != nil {
		logrus.Error("invalid request")
	}
	user, err := u.userService.FindUser(req.UserId)
	if err != nil {
		logrus.Error("User not found")
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
