package repository

import (
	"golang-crud-clean-architecture/model/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll(u []*entity.User) ([]*entity.User, error)
	Save(user *entity.User) (*entity.User, error)
	FindByUserId(id int64) (*entity.User, error)
	UpdateUser(user *entity.User, id int64) (*entity.User, error)
	DeleteUser(id int64) error
	Login(user *entity.User) (*entity.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*entity.User) ([]*entity.User, error) {
	err := ur.Db.Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *userRepository) Save(user *entity.User) (*entity.User, error) {
	err := ur.Db.Where("name =? AND last_name =? AND user_name =?", user.Name, user.LastName,
		user.UserName).First(&user).Error

	err = ur.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (ur *userRepository) FindByUserId(id int64) (*entity.User, error) {
	var user entity.User
	err := ur.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (ur *userRepository) UpdateUser(user *entity.User, id int64) (*entity.User, error) {

	var dbUser entity.User
	var err = ur.Db.Find(&dbUser, id).Error
	dbUser.UserName = user.UserName
	dbUser.Age = user.Age

	err = ur.Db.Save(&dbUser).Error

	if err != nil {
		return nil, err
	}

	return user, err
}
func (ur *userRepository) DeleteUser(id int64) error {
	var user entity.User
	err := ur.Db.Delete(&user, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Login(user *entity.User) (*entity.User, error) {
	err := ur.Db.Where("email = ? OR id = ?", user.Email,
		user.Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
