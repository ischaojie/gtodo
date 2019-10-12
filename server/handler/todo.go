package handler

import (
	"github.com/gin-gonic/gin"
	"mini_todo/errno"
	"mini_todo/model"
	"strconv"
)

type ListTodo struct {
	Total    uint64            `json:"total"`
	TodoList []model.TodoModel `json:"todos"`
}

type createTodo struct {
}

/*todos 路由相关处理函数*/

// * get all todos api.
func FetchAllTodo(c *gin.Context) {
	var todo model.TodoModel
	count, todos, err := todo.GetAll()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, ListTodo{
		Total:    count,
		TodoList: todos,
	})
}

func FetchSingleTodo(c *gin.Context) {
	var todo model.TodoModel
	var err error

	id, _ := strconv.Atoi(c.Param("id"))
	todo.ID = uint(id)

	if todo, err = todo.Get(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, todo)
}

// TODO title为空时错误
func AddTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))

	todo := model.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}

	// * 创建一条记录，错误处理
	if err := todo.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// * 返回结果
	SendResponse(c, nil, "create successful.")
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	todo.ID = uint(id)
	if err := todo.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// * 返回结果
	SendResponse(c, nil, "update successful.")
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := model.TodoModel{}

	todo.ID = uint(id)

	if err := todo.Delete(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, "delete successful.")
}
