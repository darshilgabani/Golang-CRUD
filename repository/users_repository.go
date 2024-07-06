package repository

import (
	"goguru/model"
)

type UsersRepository interface {
	Save(user model.Users)
	Update(user model.Users)
	Delete(userId int)
	FindById(userId int) (user model.Users, err error)
	FindAll() []model.Users
}
