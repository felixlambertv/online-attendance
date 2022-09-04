package response

type BaseResponse struct {
	HttpCode int    `json:"-"`
	Success  bool   `json:"success"`
	Errors   any    `json:"errors,omitempty"`
	Data     any    `json:"data,omitempty"`
	Message  string `json:"message"`
}
