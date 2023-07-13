package app

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (g *Gin) Response(httpCode, code int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Status: Status{
			Code:    code,
			Message: "123",
		},
		Data: data,
	})
}
