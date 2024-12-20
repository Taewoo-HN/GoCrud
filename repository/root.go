package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	User *UserRepository
}

func NewService() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})
	return repositoryInstance

}
