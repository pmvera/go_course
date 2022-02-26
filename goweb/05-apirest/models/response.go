package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status       int         `json:"status"`
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	content_type string
	resp_write   http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:       http.StatusOK,
		resp_write:   rw,
		content_type: "application/json",
	}
}

func (resp *Response) Send() {
	resp.resp_write.Header().Set("Content-type", resp.content_type)
	resp.resp_write.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.resp_write, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resourse not found"
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity not found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
