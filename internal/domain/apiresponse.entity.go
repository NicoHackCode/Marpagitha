package domain

type ApiResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
}
