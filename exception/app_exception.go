package exception

type AppException struct {
	Message string
}

func NewAppException(message string) *AppException {
	return &AppException{Message: message}
}

func (e AppException) Error() string {
	return e.Message
}
