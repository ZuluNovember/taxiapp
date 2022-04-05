package app

import "github.com/gin-gonic/gin"

type Response struct {
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Gin struct {
	C *gin.Context
}

// Creates a gin response and sends as json.
func (g *Gin) Response(status int, data any, err error) {
	res := Response{
		Msg:  "Ok",
		Data: data,
	}
	if err != nil {
		res.Msg = err.Error()
	}

	g.C.JSON(status, res)
}
