package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/shiniao/gtodo/pkg/auth"
	"time"
)

// UserModel 用户模型
type UserModel struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Username  string    `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password  string    `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Phone     int       `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Sex       int       `gorm:"column:sex" json:"sex"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// UserInfo 对外暴露的结构体
type UserInfo struct {
	ID       uint64 `json:"id" example:"1"`
	Username string `json:"username" example:"张三"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
}

// TableName 表名
func (u *UserModel) TableName() string {
	return "users"
}

// Token: JWT
type Token struct {
	Token string `json:"token"`
}

// Compare 比较加密前后密码是否一致
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt 加密用户密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
