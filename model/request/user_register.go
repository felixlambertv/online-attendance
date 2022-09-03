package request

type UserRegister struct {
	Name string `form:"name" binding:"required"`
}
