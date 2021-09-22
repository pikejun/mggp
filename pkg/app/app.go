package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应处理
type Response struct {
	C *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{
		C: c,
	}
}

func (r *Response) ToResponse(data interface{}) {
	data = gin.H{"status":"1","message":"success","success": true, "result":data}
	r.C.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseFailed(status int,msg string) {
	data := gin.H{"status":status,"message":msg,"success": false, "result":nil}
	r.C.JSON(http.StatusOK, data)
}
