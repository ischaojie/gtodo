package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shiniao/gtodo/errno"
	"github.com/shiniao/gtodo/model"
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
// @Summary Get all todos
// @Description Get all todos from database
// @Tags todo
// @Accept  json
// @Produce  json
// @Success 200 {object} model.TodoModel "{"code":0,"message":"OK","data":{"total":233, "todos":[{"ID":91,"title": "烫头", "completed": 1,"CreatedAt": "2019-10-12T10:10:05+08:00","UpdatedAt": "2019-10-12T10:16:24+08:00","DeletedAt": null}]}}"
// @Router /v1/todos [get]
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

// @Summary Get a todo
// @Description Get a todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Success 200 {object} model.TodoModel "{"code":0,"message":"OK","data":{"ID":91,"title": "烫头", "completed": 1,"CreatedAt": "2019-10-12T10:10:05+08:00","UpdatedAt": "2019-10-12T10:16:24+08:00","DeletedAt": null}}"
// @Router /v1/todos/{id} [get]
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
// @Summary Add new todos to the database
// @Description Add a new todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Param todo body model.TodoModel true "The todo info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":"create successful."}"
// @Router /v1/todos/ [post]
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

// @Summary Update a todo
// @Description Update a todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The todo's database id index num"
// @Param todo body model.TodoModel true "The todo info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":"update successful."}"
// @Router /v1/todos/{id} [put]
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

// @Summary Delete a todo
// @Description Delete a todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":"delete successful."}"
// @Router /v1/todos/{id} [delete]
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
