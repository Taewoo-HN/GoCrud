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
func (u *User) Update(name string, age int64) error {
	return u.repo.Update(name, age)
}
func (u *User) Get() []*types.User {
	return u.repo.Get()
}
