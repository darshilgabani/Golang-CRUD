package service

import (
	"goguru/data/request"
	"goguru/data/response"
	"goguru/helper"
	"goguru/model"
	"goguru/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Create implements UsersService.
func (u *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := u.Validate.Struct(users)
	helper.ErrorPanic(err)
	userModel := model.Users{
		FirstName: users.FirstName,
	}
	u.UsersRepository.Save(userModel)
}

// Delete implements UsersService.
func (u *UsersServiceImpl) Delete(usersId int) {
	u.UsersRepository.Delete(usersId)
}

// FindAll implements UsersService.
func (u *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := u.UsersRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		tag := response.UsersResponse{
			Id:        value.Id,
			FirstName: value.FirstName,
		}
		users = append(users, tag)
	}

	return users
}

// FindById implements UsersService.
func (u *UsersServiceImpl) FindById(usersId int) response.UsersResponse {
	userData, err := u.UsersRepository.FindById(usersId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		Id:        userData.Id,
		FirstName: userData.FirstName,
	}
	return userResponse
}

// Update implements UsersService.
func (u *UsersServiceImpl) Update(users request.UpdateUsersRequest) {
	userData, err := u.UsersRepository.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.FirstName = users.FirstName
	u.UsersRepository.Update(userData)
}
