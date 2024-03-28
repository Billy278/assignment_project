package responses

const (
	Unauthorized = "unauthorized request"
	Success      = "Success"
	NotFound     = "NOT FOUND"
	InvalidBody  = "invalid body request"
	InvalidParam = "Invalid Param"
)

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
