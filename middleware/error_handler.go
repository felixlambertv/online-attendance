package middleware

import (
	"github.com/felixlambertv/online-attendance/exception"
	"github.com/felixlambertv/online-attendance/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	for _, err := range ctx.Errors {
		var res response.BaseResponse
		res.Message = err.Error()
		res.Success = false
		switch err.Err.(type) {
		case *exception.AppException:
			res.HttpCode = http.StatusInternalServerError
		case *exception.DataNotFoundException:
			res.HttpCode = http.StatusNotFound
		case validator.ValidationErrors:
			data := make(map[string]string)
			for _, fieldError := range err.Err.(validator.ValidationErrors) {
				data[fieldError.Field()] = fieldError.ActualTag()
			}
			res.HttpCode = http.StatusBadRequest
			res.Errors = data
			res.Message = "Invalid inputs. Please check your inputs."
		default:
			res.HttpCode = http.StatusInternalServerError
			res.Message = "something wrong"
		}
		ctx.AbortWithStatusJSON(res.HttpCode, res)
	}
}
