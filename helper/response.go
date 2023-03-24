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
	resp := map[string]interface{}{
		"status":  "",
		"message": "",
	}
	code := http.StatusInternalServerError
	if msg != "" {
		resp["message"] = msg
	}
	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "null"):
		code = http.StatusBadRequest
		resp["status"] = "Bad Request"
	case strings.Contains(msg, "Not Found"):
		code = http.StatusNotFound
		resp["status"] = "Not Found"
	case strings.Contains(msg, "input"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "exist"):
		code = http.StatusBadRequest
	}

	return code, resp
}
