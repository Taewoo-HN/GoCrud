package repository

import (
	"awesomeProject/types"
	"awesomeProject/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		//userMap
		userMap: []*types.User{},
	}
}

func (repo *UserRepository) Create(newUser *types.User) error {
	repo.userMap = append(repo.userMap, newUser)
	return nil
}
func (repo *UserRepository) Delete(userName string) error {
	isExist := false
	for index, user := range repo.userMap {
		if userName == user.Name {
			repo.userMap = append(repo.userMap[:index], repo.userMap[index+1:]...)
			isExist = true
			continue
		}
	}
	if !isExist {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
	return nil
}
func (repo *UserRepository) Update(name string, newAge int64) error {
	isExist := false
	for _, user := range repo.userMap {
		// 포인트 타입에 대한 슬라이스, 특정값을 바꾸기 위해서는 포인트타입을 사용해야함
		if user.Name == name {
			user.Age = newAge
			isExist = true
			continue
		}
		if !isExist {
			return errors.Errorf(errors.NotFoundUser, nil)
		} else {
			return nil
		}
	}
	return nil
}

func (repo *UserRepository) Get() []*types.User {
	return repo.userMap
}
