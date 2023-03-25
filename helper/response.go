package helper

import (
	"net/http"
	"strings"
)

type SuccessResponse struct {
	Status  string      `json:"status" form:"status"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

func PrintSuccessReponse(status string, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

}

func DeleteReponse(status string, message string) SuccessResponse {
	return SuccessResponse{
		Status:  status,
		Message: message,
		Data:    struct{}{},
	}

}

func PrintErrorResponse(msg string) (int, ErrorResponse) {
	resp := ErrorResponse{}
	code := http.StatusInternalServerError
	if msg != "" {
		resp.Message = msg
	}
	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "null"):
		code = http.StatusBadRequest
		resp.Status = "Bad Request"
	case strings.Contains(msg, "Not Found"):
		code = http.StatusNotFound
		resp.Status = "Not Found"
	case strings.Contains(msg, "input"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "exist"):
		code = http.StatusBadRequest
	}

	return code, resp
}
