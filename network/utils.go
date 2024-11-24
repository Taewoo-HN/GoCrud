package network

import (
	"github.com/gin-gonic/gin"
)

func (h Network) RegisterGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return h.engine.GET(path, handler...)
}
func (h Network) RegisterPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return h.engine.POST(path, handler...)
}
func (h Network) RegisterDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return h.engine.DELETE(path, handler...)
}
func (h Network) RegisterUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return h.engine.PUT(path, handler...)
}

func (h Network) okResponse(g *gin.Context, result interface{}) {
	g.JSON(200, result)
}

func (h Network) failedResponse(g *gin.Context, result interface{}) {
	g.JSON(400, result)
}
