package api

import (
	"strconv"
	"tasks_list/dto"
	"tasks_list/pkg/response"
	"tasks_list/wire"

	"github.com/gin-gonic/gin"
)

// ListTasks
// @Summary List all tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} dto.ListTaskResponse
// @Router /tasks [get]
func ListTasks(c *gin.Context) {
	srv := wire.BuildTaskService()
	tasks, err := srv.ListTasks()
	if err != nil {
		_ = c.Error(err)
	}
	c.JSON(200, response.Response{Result: tasks})
}

// CreateTask
// @Summary Create task
// @Tags Task
// @Accept json
// @Produce json
// @Param body body dto.CreateTaskRequest true "status 不傳為 0"
// @Success 201 {object} dto.CreateTaskResponse
// @Failure 400 string string "參數錯誤"
// @Failure 500 string string  "內部錯誤"
// @Router /task [post]
func CreateTask(c *gin.Context) {
	req := &dto.CreateTaskRequest{}
	if err := c.Bind(req); err != nil {
		_ = c.Error(err)
		return
	}
	srv := wire.BuildTaskService()
	res, err := srv.CreateTask(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(201, response.Response{Result: res})

}

// UpdateTask
// @Summary Update task
// @Tags Task
// @Accept json
// @Produce json
// @Param body body dto.UpdateTaskRequest true "需要傳整個 task 的欄位"
// @Success 200 {object} dto.CreateTaskResponse
// @Failure 400 string string "參數錯誤"
// @Failure 404 string string "資料不存在"
// @Failure 500 string string "內部錯誤"
// @Router /task/{id} [put]
func UpdateTask(c *gin.Context) {
	req := &dto.UpdateTaskRequest{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}
	req.Id = id
	if err := c.Bind(req); err != nil {
		_ = c.Error(err)
		return
	}
	srv := wire.BuildTaskService()
	res, err := srv.UpdateTask(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, response.Response{Result: res})
}

// DeleteTask
// @Summary Delete task
// @Tags Task
// @Accept json
// @Produce json
// @Param body body dto.DeleteTaskRequest true "task id"
// @Success 200 {object} dto.CreateTaskResponse
// @Failure 400 string string "參數錯誤"
// @Failure 404 string string "資料不存在"
// @Failure 500 string string "內部錯誤"
// @Router /task/{id} [delete]
func DeleteTask(c *gin.Context) {
	req := &dto.DeleteTaskRequest{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}
	req.Id = &id
	if err := c.Bind(req); err != nil {
		_ = c.Error(err)
		return
	}
	srv := wire.BuildTaskService()
	if err := srv.DeleteTask(req); err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, nil)
}
