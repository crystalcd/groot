package controller

import (
	"net/http"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/domain/e"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	App *bootstrap.Application
}

func NewTestController() *TestController {
	return &TestController{}
}

func (tc *TestController) Test(c *gin.Context) {
	r := domain.Gin{C: c}
	var t domain.Test
	if err := c.ShouldBindJSON(&t); err != nil {
		r.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	r.Response(http.StatusOK, e.SUCCESS, t)
}
