package model

import (
	"github.com/jinzhu/gorm"
	"github.com/thanhngodinh/mysql-simple/app"
)

var db *gorm.DB

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Email string `json:"email"`
}

func init() {
	app.Connect()
	db = app.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) Create() *User {
	db.NewReacord(u)
	db.Create(&u)
	return u
}

func GetAll() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func Delete(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}