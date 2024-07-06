package service

import (
	"goguru/data/request"
	"goguru/data/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(users request.UpdateUsersRequest)
	Delete(usersId int)
	FindById(usersId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
