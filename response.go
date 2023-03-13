package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, &ResponseBody{
		Msg:  "success",
		Data: data,
	})
}

func HandleError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	if e, ok := err.(Error); ok {
		WithError(ctx, e)
		return
	}
	WithError(ctx, ErrorBadRequest(40000, err.Error()))
}

func WithError(ctx *gin.Context, err Error) {
	ctx.JSON(err.StatusCode(), &ResponseBody{
		Code: err.ErrorCode(),
		Msg:  err.Msg(),
		Data: nil,
	})
}

func Middleware(c *gin.Context) {
	c.Next()
	if c.Errors != nil {
		for _, err := range c.Errors {
			if respErr, ok := err.Err.(Error); ok {
				WithError(c, respErr)
				return
			}
		}
	}
}
