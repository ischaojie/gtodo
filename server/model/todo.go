package model

import "github.com/jinzhu/gorm"

// * bind orm
type TodoModel struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title"`
	Completed int    `json:"completed"`
}



// * 设置表名
func (todo TodoModel) TableName() string {
	return "todos"
}

// * 添加todo到数据库
func (todo TodoModel) Create() error {
	return DB.Self.Create(&todo).Error
}

// * 删除某个todo
func (todo TodoModel) Delete() error {
	return DB.Self.Delete(&todo).Error
}

// * Update
func (todo TodoModel) Update() error {
	return DB.Self.Save(&todo).Error
}

// * 获取某一条todo
func (todo TodoModel) Get() (TodoModel, error) {
	return todo, DB.Self.First(&todo, todo.ID).Error
}

func (todo TodoModel) GetAll() (uint64, []TodoModel, error) {

	var todos []TodoModel
	var count uint64

	if err := DB.Self.Table(todo.TableName()).Count(&count).Error; err != nil {
		return count, todos, err
	}
	if err := DB.Self.Find(&todos).Error; err != nil {
		return count, todos, err
	}
	return count, todos, nil

}
