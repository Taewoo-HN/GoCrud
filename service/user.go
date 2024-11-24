package service

import (
	"awesomeProject/repository"
	"awesomeProject/types"
)

type User struct {
	repo *repository.UserRepository
}

func newUserService(repo *repository.UserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.repo.Create(newUser)
}
func (u *User) Delete(newUser *types.User) error {
	return u.repo.Delete(newUser.Name)
}
func (u *User) Update(beforeUser *types.User, afterUser *types.User) error {
	return u.repo.Update(beforeUser, afterUser)
}
func (u *User) Get() []*types.User {
	return u.repo.Get()
}
