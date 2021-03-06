package middleware

import (
	h "github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//SetRequestID set request id header
func SetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewV4()
		setHeader(&c.Request.Header, h.RequestIDXHeader, id.String())
	}
}

//SetRequestName set request name header
func SetRequestName(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		setHeader(&c.Request.Header, h.RequestNameXHeader, name)
	}
}
