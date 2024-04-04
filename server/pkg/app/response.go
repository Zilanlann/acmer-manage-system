package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"message"`
	Data    interface{} `json:"result"`
}

// SuccessResponse generates a success response with the provided HTTP status code, custom status code, and data.
//
// httpCode: the HTTP status code to return.
// StatusCode: the custom status code.
// data: the data to include in the response.
func (g *Gin) SuccessResponse(httpCode, StatusCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Success: true,
		Code:    StatusCode,
		Msg:     e.GetMsg(StatusCode),
		Data:    data,
	})
}

// ErrorResponse generates a error response with the provided HTTP status code, custom status code, and data.
//
// httpCode: the HTTP status code to return.
// StatusCode: the custom status code.
// data: the data to include in the response.
func (g *Gin) ErrorResponse(httpCode, StatusCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Success: true,
		Code:    StatusCode,
		Msg:     e.GetMsg(StatusCode),
		Data:    data,
	})
}
