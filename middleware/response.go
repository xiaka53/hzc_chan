package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/DeployAndLog/lib"
)

const (
	ERROR   ResponseCode = 500
	SUCCESS ResponseCode = 200
)

type ResponseCode int

//返回信息格式
type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"trace_id"`
}

//返回前端错误信息
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	var (
		trace        interface{}
		traceContext *lib.TraceContext
		traceId      string
		resp         Response
		respone      []byte
	)
	trace, _ = c.Get("trace")
	traceContext, _ = trace.(*lib.TraceContext)
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp = Response{
		ErrorCode: code,
		ErrorMsg:  err.Error(),
		Data:      "",
		TraceId:   traceId,
	}
	c.JSON(200, resp)
	respone, _ = json.Marshal(resp)
	c.Set("response", string(respone))
	_ = c.AbortWithError(200, err)
}

//返回前端信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	var (
		trace        interface{}
		traceContext *lib.TraceContext
		traceId      string
		response     []byte
		resp         Response
	)
	trace, _ = c.Get("trace")
	traceContext, _ = trace.(*lib.TraceContext)
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp = Response{
		ErrorCode: SUCCESS,
		ErrorMsg:  "",
		Data:      data,
		TraceId:   traceId,
	}
	c.JSON(200, resp)
	response, _ = json.Marshal(resp)
	c.Set("response", string(response))
}
