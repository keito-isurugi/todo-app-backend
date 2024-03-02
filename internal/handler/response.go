package handler

import "encoding/json"

type emptyResponse struct{}

type createdResponse struct {
	ID any `json:"id" example:"1" swaggertype:"string"`
}

type createdUUIDResponse struct {
	ID string `json:"id" example:"cc293e0a-7342-4aac-b49b-a851e8af9dfc" swaggertype:"string"`
}

func newCreatedResponse(id any) createdResponse {
	return createdResponse{
		ID: id,
	}
}

func createCommonCreatedResponse(responseID any) string {
	resp := newCreatedResponse(responseID)

	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}

func createCommonEmptyResponse() string {
	resp := emptyResponse{}

	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}
