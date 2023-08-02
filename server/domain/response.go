package domain

import (
	"github.com/crystal/groot/domain/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpcode, errcode int, data interface{}) {
	g.C.JSON(httpcode, Response{
		Code: errcode,
		Msg:  e.GetMsg(errcode),
		Data: data,
	})
}
