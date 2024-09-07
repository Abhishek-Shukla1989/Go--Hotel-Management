package pkg

import (
	constant "code/app/constant"
	"code/app/domain/dto"
)

func Null() interface{} {
	return nil
}
func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ResponseApi[T] {

	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(""), data)
}
func BuildResponse_[T any](status string, message string, data T) dto.ResponseApi[T] {

	return dto.ResponseApi[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}

}
