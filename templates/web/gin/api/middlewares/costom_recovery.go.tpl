package middlewares

import (
	"{{ .ProjectName }}/api/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := response.GenResp(false, response.InternalError, err, nil)
		c.AbortWithStatusJSON(500, httpResponse)
		return
	}
	httpResponse := response.GenResp(false, response.InternalError, err, nil)
	c.AbortWithStatusJSON(500, httpResponse)
}
