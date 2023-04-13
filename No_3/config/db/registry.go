package db

import (
	"golang-crud-clean-architecture/model/entity"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: entity.User{}},
		{Model: entity.Role{}},
	}
}
