package entity

import (
	"time"
)

type User struct {
	Id        int64  `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	UserName  string `gorm:"size:100;not null"`
	Password  string `gorm:"size:200;not null"`
	Email     string `gorm:"size:100;not null"`
	Age       int    `gorm:"size:100;not null"`
	RoleID    int64  `gorm:"size:100;not null"`
	Role      Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
