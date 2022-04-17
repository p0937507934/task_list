package tasks

import (
	"sync"
	"tasks_list/config"
	"tasks_list/driver"
	"tasks_list/dto"
	"tasks_list/models"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var orm *gorm.DB
var repo ITaskRepository
var srv ITaskService
var once sync.Once

func Init() {
	once.Do(func() {
		config.InitConfig()
		godotenv.Load("../../.env")
		orm = driver.InitGorm()
		repo = NewTasksRepository(orm)
		srv = NewTaskService(repo)
	})

}

func TestCreateTaskSrv(t *testing.T) {
	Init()
	name := "Gogolook"
	status := 0
	req := &dto.CreateTaskRequest{Name: &name, Status: status}
	res, err := srv.CreateTask(req)
	assert.Equal(t, nil, err)
	// teardown
	task := &models.Tasks{Id: res.Id}
	repo.Delete(task)
}

func TestUpdateTaskSrv(t *testing.T) {
	Init()
	// create a task first
	task := &models.Tasks{Name: "回家"}
	err := repo.Insert(task)
	assert.Equal(t, nil, err)
	// update a record success
	name := "Gogolook"
	status := 1
	req := &dto.UpdateTaskRequest{Id: task.Id, Name: &name, Status: &status}
	res, err := srv.UpdateTask(req)
	assert.Equal(t, nil, err)
	// update a record fail, data not exist
	id := 9527
	req = &dto.UpdateTaskRequest{Id: id, Name: &name, Status: &status}
	res, err = srv.UpdateTask(req)
	assert.Equal(t, res, nil)
	assert.NotEqual(t, nil, err)
	// teardown
	task = &models.Tasks{Id: task.Id}
	repo.Delete(task)
}
func TestDeleteTaskSrv(t *testing.T) {
	Init()
	// create a task first
	task := &models.Tasks{Name: "回家"}
	err := repo.Insert(task)
	assert.Equal(t, nil, err)
	// delete a task failed
	id := 9527
	req := &dto.DeleteTaskRequest{Id: &id}
	err = srv.DeleteTask(req)
	assert.NotEqual(t, nil, err)
	// delete a task success
	req.Id = &task.Id
	err = srv.DeleteTask(req)
	assert.Equal(t, nil, err)
}
