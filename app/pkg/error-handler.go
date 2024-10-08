package pkg

import (
	"code/app/constant"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s:%w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus, resource string) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage(resource))
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		// fmt.Printf("Value if key is %v", key)
		// fmt.Printf("Value if key is %v", msg)

		switch key {
		case constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		case constant.Unauthorised.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			c.Abort()
		case constant.InvalidRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		case constant.ResourceAlreadyExists.GetResponseStatus(): // New case for resource conflict
			c.JSON(http.StatusConflict, BuildResponse_(key, msg, Null())) // Include resource in response
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			c.Abort()
		}
	}
}
