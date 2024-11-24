package service

import (
	"awesomeProject/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	// repository 선언
	repository *repository.Repository
	User       *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
		serviceInstance.User = newUserService(rep.User)
	})
	return serviceInstance
}
