package repository

import (
	"golang-crud-clean-architecture/model/entity"
	"gorm.io/gorm"
)

type IRoleRepository interface {
	FindAll(role []*entity.Role) ([]*entity.Role, error)
	Save(role *entity.Role) (*entity.Role, error)
	FindByRoleId(id int64) (*entity.Role, error)
	UpdateRole(role *entity.Role, id int64) (*entity.Role, error)
	DeleteRole(id int64) error
}

type roleRepository struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &roleRepository{db}
}

func (r roleRepository) FindAll(role []*entity.Role) ([]*entity.Role, error) {
	err := r.Db.Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r roleRepository) Save(role *entity.Role) (*entity.Role, error) {
	err := r.Db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return role, err
}

func (r roleRepository) FindByRoleId(id int64) (*entity.Role, error) {
	var role entity.Role
	err := r.Db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, err
}

func (r roleRepository) UpdateRole(role *entity.Role, id int64) (*entity.Role, error) {
	var dbRole entity.Role
	var err = r.Db.Find(&dbRole, id).Error

	err = r.Db.Save(&dbRole).Error

	if err != nil {
		return nil, err
	}

	return role, err
}

func (r roleRepository) DeleteRole(id int64) error {
	var role entity.Role
	err := r.Db.Delete(&role, id).Error

	if err != nil {
		return err
	}

	return nil
}
