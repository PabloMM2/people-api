package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data  any
	Error error
}

func SuccessResponse(w http.ResponseWriter, payload interface{}, statusCode *int) {
	var code int

	if statusCode == nil {
		code = http.StatusOK
	} else {
		code = *statusCode
	}

	respPayload := &ResponseData{
		Data: payload,
	}

	response, _ := json.Marshal(respPayload)

	w.WriteHeader(code)
	w.Write(response)
}
