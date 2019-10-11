package handler

import (
	"github.com/gin-gonic/gin"
	"mini_todo/errno"
	"mini_todo/model"
	"strconv"
)

/*todos 路由相关处理函数*/

// * get all todos api.
func FetchAllTodo(c *gin.Context) {
	
	count, todo, err := model.GetAll();
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, model.ListTodo{
		Total:    count,
		TodoList: todo,
	})
}

func FetchSingleTodo(c *gin.Context) {
	var todo model.TodoModel
	var err error

	id, _ := strconv.Atoi(c.Param("id"))
	if todo, err = model.Get(uint(id)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, todo)
}

// TODO title为空时错误
func AddTodo(c *gin.Context) {
	//completed, _ := strconv.Atoi(c.PostForm("completed"))
	//
	//todo := model.TodoModel{
	//	Title:     c.PostForm("title"),
	//	Completed: completed,
	//}
	var todo model.TodoModel
	c.Bind(&todo)
	// * 创建一条记录，错误处理
	if err := model.Create(&todo); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// * 返回结果
	SendResponse(c, nil, "create successful.")
}

func UpdateTodo(c *gin.Context) {

	var todo model.TodoModel
	c.Bind(&todo)
	if err := model.Update(&todo); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// * 返回结果
	SendResponse(c, nil, "update successful.")
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := model.Delete(uint(id)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, "delete successful.")
}
