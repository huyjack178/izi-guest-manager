package xhttp
import (
	"net/http"
	"encoding/json"
)

type Response struct {
	StatusCode int `json:"statusCode"`
	Message interface{} `json:"message"`
}

func NewResponse(statusCode int, message interface{}) *Response  {
	return &Response{
		StatusCode:statusCode,
		Message:message,
	}
}

func ResponseJson(w http.ResponseWriter, response *Response)  {
	responseJS, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(responseJS)
}
