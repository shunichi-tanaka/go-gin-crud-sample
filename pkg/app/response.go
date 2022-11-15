package app

import (
	"net/http"

	"go-gin-crud-example/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"messagge"`
}

func (g *Gin) SuccessResponse(data interface{}) {
	g.C.JSON(http.StatusOK, data)
	return
}

func (g *Gin) ErrorResponse(httpCode, errCode int) {
	g.C.JSON(httpCode, ErrorResponse{
		ErrorCode: errCode,
		Message:   e.GetMsg(errCode),
	})
	return
}
