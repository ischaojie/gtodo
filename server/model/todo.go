package model

import "github.com/jinzhu/gorm"

// * bind orm
type TodoModel struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title"`
	Completed int    `json:"completed"`
}


type ListTodo struct {
	Total    uint64       `json:"total"`
	TodoList []*TodoModel `json:"todos"`
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

func GetAll() (uint64, []*TodoModel, error) {

	todo := make([]*TodoModel, 0)
	var t TodoModel
	var count uint64

	if err := DB.Self.Table(t.TableName()).Count(&count).Error; err != nil {
		return count, todo, err
	}
	if err := DB.Self.Find(&todo).Error; err != nil {
		return count, todo, err
	}
	return count, todo, nil

}
