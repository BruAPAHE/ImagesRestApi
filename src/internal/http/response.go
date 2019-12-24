package http

import (
	"encoding/json"
)

type ResponseBody struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"meta"`
}

func MakeJSONResponse(statusCode int, message string) ([]byte, error) {

	response := &ResponseBody{
		Meta: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{
			Message: message,
			Code:    statusCode,
		},
	}

	return json.Marshal(response)
}

