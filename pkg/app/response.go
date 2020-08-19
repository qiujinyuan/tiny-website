package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yrjkqq/tiny-website/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, extraErrMsg string, data interface{}) {
	msg := e.GetMsg(errCode)
	if extraErrMsg != "" {
		msg = fmt.Sprintf("%v, detail: [%v]", msg, extraErrMsg)
	}
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})
}
