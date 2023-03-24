package helper

import (
	"net/http"
	"strings"
)

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	}

	return code, resp

}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
		code = http.StatusBadRequest
		code = http.StatusBadRequest
	case strings.Contains(msg, "tidak ditemukan"):
		code = http.StatusNotFound
	case strings.Contains(msg, "not found"):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "input invalid"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "input"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "denied"):
		code = http.StatusUnauthorized
	case strings.Contains(msg, "exist"):
		code = http.StatusBadRequest
	}

	return code, resp
}
