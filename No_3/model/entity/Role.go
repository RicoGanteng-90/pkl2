package entity

type Role struct {
	Id   int64  `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name string `gorm:"size:100;not null"`
}
