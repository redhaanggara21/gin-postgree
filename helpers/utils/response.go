package util

type Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}
type ResponseEror struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, code int, status bool, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}
	return jsonResponse
}
func APIResponseFailed(message error, code int, status bool, data interface{}) ResponseEror {
	jsonResponse := ResponseEror{
		Message: message.Error(),
		Status:  status,
		Data:    data,
	}
	return jsonResponse
}
