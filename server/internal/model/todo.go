package model

import "time"

// TodoModel 代表todo模型
type TodoModel struct {
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	UserID uint64 `gorm:"column:user_id" json:"user_id"`
	Title     string `gorm:"column:title" json:"title"`
	Completed int    `gorm:"column:completed" json:"completed"`
	CompletedAt time.Time `gorm:"column:completed_at" json:"completed_at"`
	Description string `gorm:"column:description" json:"description"`
	RepeatType string `gorm:"column:repeat_type" json:"repeat_type"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}



// * 设置表名
func (todo TodoModel) TableName() string {
	return "todos"
}

// * 添加todo到数据库
func (todo TodoModel) Create() error {
	return DB.Default.Create(&todo).Error
}

// * 删除某个todo
func (todo TodoModel) Delete() error {
	return DB.Default.Delete(&todo).Error
}

// * Update
func (todo TodoModel) Update() error {
	return DB.Default.Save(&todo).Error
}

// * 获取某一条todo
func (todo TodoModel) Get() (TodoModel, error) {
	return todo, DB.Default.First(&todo, todo.ID).Error
}

func (todo TodoModel) GetAll() (uint64, []TodoModel, error) {

	var todos []TodoModel
	var count uint64

	if err := DB.Default.Table(todo.TableName()).Count(&count).Error; err != nil {
		return count, todos, err
	}
	if err := DB.Default.Find(&todos).Error; err != nil {
		return count, todos, err
	}
	return count, todos, nil

}
