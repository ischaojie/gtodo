/*todo_repo 对接数据库，提供基础的CRUD函数给handler*/
package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shiniao/gtodo/internal/model"
)

type TodoRepo interface {
	CreateUserTodo(db *gorm.DB, todo *model.TodoModel) (id uint64, err error)
	UpdateUserTodo(db *gorm.DB, id uint64, todoMap map[string]interface{}) error

	GetUserTodos(db *gorm.DB, id uint64) (*model.TodoModel, error)

}

type userTodoRepo struct{

}

func (userTodoRepo) CreateUserTodo(db *gorm.DB, todo *model.TodoModel) (id uint64, err error) {

}

func (userTodoRepo) Update(db *gorm.DB, id uint64, todoMap map[string]interface{}) error {
	panic("implement me")
}

func (userTodoRepo) GetTodoByIDs(db *gorm.DB, id uint64) (*model.TodoModel, error) {
	panic("implement me")
}
