package api

import "github.com/gin-gonic/gin"

func ProtectedGroup(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/", AuthMiddleware())
}
