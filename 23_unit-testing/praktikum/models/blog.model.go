package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Id_User string `json:"id_user" form:"id_user"`
	User    User   `gorm:"foreignKey:id_user;" json:"user"`
	Judul   string `json:"judul" form:"judul"`
	Konten  string `json:"konten" form:"konten"`
}
