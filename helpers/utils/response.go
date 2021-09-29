package util

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}
	return jsonResponse
}
