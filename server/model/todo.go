package model

import "github.com/jinzhu/gorm"

// * bind orm
type TodoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// * 展示给用户的
type TransformedTodo struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// * 设置表名
func (t *TodoModel) TableName() string {
	return "todos"
}

// * 添加todo到数据库
func Create(todo *TodoModel) error {
	return DB.Self.Create(&todo).Error
}

// * 删除某个todo
func Delete(id uint) error {
	todo := TodoModel{}
	todo.ID = id
	return DB.Self.Delete(&todo).Error
}

// * Update
func Update(todo *TodoModel) error {
	return DB.Self.Save(&todo).Error
}

// * 获取某一条todo
func Get(id uint) (TodoModel, error) {
	var todo TodoModel
	return todo, DB.Self.First(&todo, id).Error
}


