package network

import (
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
)

type Network struct {
	engine  *gin.Engine
	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	h := &Network{
		engine: gin.New(),
	}
	newUserRouter(h, service.User)

	return h
}
func (h Network) ServerStart(port string) error {
	return h.engine.Run(port)
}
