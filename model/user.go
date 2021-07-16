package model

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;size:20;index"`
	Password string `gorm:"not null;size:16"`
	Head string  `gorm:"size:255"`
	Sign string  `gorm:"default:这个人很懒，什么也欸有些"`
}

//注册（创建新的user）
func (u *User)Signup() (user *User,err error) {

	if err:=db.Table("users").Create(u).Error;err!=nil{
		return &User{},err
	}
	return &User{
		Model :u.Model,
		Username: u.Username,
	},nil
}

//登录
func (u *User)Login() (user *User,err error) {
	if err:=db.Table("users").Where("username = ?",u.Username).First(user).Error;err!=nil{
		return &User{},err
	}
	if u.Password == user.Password {
		return user,nil
	}else {
		return &User{},errors.New("密码错误")
	}
	
}

//通过用户id查找用户信息
func GetUserInfo(uid interface{}) (user *User,err error) {
	if err=db.Table("users").Where("id = ?",uid).First(user).Error;err!=nil{
		return &User{}, err
	}
	return user,nil
}

//修改用户信息
func (u *User)ModifyUserInfo() (user *User,err error) {
	tx:=db.Begin()
	err=tx.Table("users").Model(u).Updates(User{Username: u.Username,Password: u.Password,Head: u.Head,Sign: u.Sign}).Error
	if err!=nil {
		tx.Rollback()
		return &User{}, err
	}
	if err=tx.Where("id = ?",u.ID).First(user).Error;err!=nil{
		tx.Rollback()
		return &User{}, err
	}
	return user,nil
}
