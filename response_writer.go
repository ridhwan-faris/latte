package latte

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	ResponseCode string      `json:"response_code"`
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func (c *Context) ResponseWithJson(w http.ResponseWriter, httpCode int, code string, status string, message string, data interface{}) {
	res := responseWriter{
		ResponseCode: code,
		Status:       status,
		Message:      message,
		Data:         data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(res)
}

func (c *Context) SendSuccess(w http.ResponseWriter, code string, message string, payload interface{}) {
	c.ResponseWithJson(w, 200, code, "SUCCESS", message, payload)
}

func (c *Context) SendBadRequest(w http.ResponseWriter, code string, message string) {
	c.ResponseWithJson(w, 400, code, "ERROR", message, nil)
}

func (c *Context) SendUnauthorized(w http.ResponseWriter, code string, message string) {
	c.ResponseWithJson(w, 401, code, "ERROR.", message, nil)
}

func (c *Context) SendForbidden(w http.ResponseWriter, code string, message string) {
	c.ResponseWithJson(w, 403, code, "ERROR", message, nil)
}

func (c *Context) SendInternalServerError(w http.ResponseWriter, code string, message string) {
	c.ResponseWithJson(w, 500, code, "ERROR", message, nil)
}
