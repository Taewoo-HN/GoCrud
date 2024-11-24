package cmd

import (
	"awesomeProject/config"
	"awesomeProject/network"
	"awesomeProject/repository"
	"awesomeProject/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}
	c.repository = repository.NewService()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)
	return c
}
