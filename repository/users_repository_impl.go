package repository

import (
	"errors"
	"goguru/data/request"
	"goguru/helper"
	"goguru/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements UsersRepository.
func (u *UsersRepositoryImpl) Delete(userId int) {
	var user model.Users
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository.
func (u *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository.
func (u *UsersRepositoryImpl) FindById(userId int) (users model.Users, err error) {
	var user model.Users
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements UsersRepository.
func (u *UsersRepositoryImpl) Save(user model.Users) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository.
func (u *UsersRepositoryImpl) Update(user model.Users) {
	var updateUser = request.UpdateUsersRequest{
		Id:        user.Id,
		FirstName: user.FirstName,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{
		Db: Db,
	}
}
