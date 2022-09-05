package controller

import (
	"errors"
	"github.com/felixlambertv/online-attendance/exception"
	"github.com/felixlambertv/online-attendance/model/request"
	"github.com/felixlambertv/online-attendance/model/response"
	"github.com/felixlambertv/online-attendance/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type IUserController interface {
	Detail(ctx *gin.Context)
	Register(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Data:    user,
		Message: "Success get user",
	})
}

func (u UserController) Register(ctx *gin.Context) {
	var req request.UserRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	user, err := u.userService.Register(req)
	if err != nil {
		ctx.Error(errors.New("something wrong"))
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Data:    user,
		Message: "Success register user",
	})
}

func (u UserController) Delete(ctx *gin.Context) {
	var req request.UserDetail
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(exception.NewAppException("wrong user id"))
		return
	}
	err = u.userService.Delete(req.UserId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Message: "Success delete user",
	})
}

func (u UserController) Update(ctx *gin.Context) {
	var req request.UserRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	var userDetail request.UserDetail
	err = ctx.ShouldBindUri(&userDetail)
	if err != nil {
		ctx.Error(exception.NewAppException("wrong user id"))
		return
	}

	user, err := u.userService.Update(userDetail.UserId, req)
	if err != nil {
		ctx.Error(errors.New("something wrong"))
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Data:    user,
		Message: "Success update user",
	})
}
