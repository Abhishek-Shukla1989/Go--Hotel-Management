package dto

type ResponseApi[T any] struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
}
