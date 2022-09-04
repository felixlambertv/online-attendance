package middleware

import (
	"github.com/felixlambertv/online-attendance/exception"
	"github.com/felixlambertv/online-attendance/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	for _, err := range ctx.Errors {
		var res response.BaseResponse
		res.Message = err.Error()
		switch err.Err.(type) {
		case *exception.AppException:
			res.Status = http.StatusInternalServerError
		case *exception.DataNotFoundException:
			res.Status = http.StatusNotFound
		default:
			res.Status = http.StatusInternalServerError
			res.Message = "something wrong"
		}
		ctx.AbortWithStatusJSON(res.Status, res)
	}
}
