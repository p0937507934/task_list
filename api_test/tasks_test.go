package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"tasks_list/config"
	"tasks_list/driver"
	"tasks_list/dto"
	"tasks_list/internal/tasks"
	"tasks_list/models"
	"tasks_list/route"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var orm *gorm.DB
var repo tasks.ITaskRepository
var srv tasks.ITaskService
var once sync.Once
var c *gin.Context
var routes *gin.Engine

func setup() {
	Init()
	gin.SetMode(gin.TestMode)
	routes = route.Serve()
	fmt.Println("Before all tests")
}

func Init() {
	once.Do(func() {
		config.InitConfig()
		godotenv.Load("../../.env")
		orm = driver.InitGorm()
		repo = tasks.NewTasksRepository(orm)
		srv = tasks.NewTaskService(repo)

	})

}

func TestListTasks(t *testing.T) {
	setup()
	req, _ := http.NewRequest(http.MethodGet, "/tasks", bytes.NewBuffer(nil))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	routes.ServeHTTP(response, req)
	assert.Equal(t, 200, response.Code)
}

func TestCreateTasks(t *testing.T) {
	setup()
	data := &models.Tasks{Name: "Gogolook"}
	body, err := json.Marshal(data)
	assert.Equal(t, nil, err)
	req, _ := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	routes.ServeHTTP(response, req)
	assert.Equal(t, 201, response.Code)
}

func TestUpdateTasks(t *testing.T) {
	setup()
	// create task to update
	task := models.Tasks{Id: 1, Name: "test", Status: 0}
	err := repo.Insert(&task)
	assert.Equal(t, nil, err)
	tests := []struct {
		ExpectedCode   int
		ExpectedResult *dto.UpdateTaskResponse
	}{
		{
			ExpectedCode:   200,
			ExpectedResult: &dto.UpdateTaskResponse{Id: 1, Name: "Gogolook good", Status: 1},
		},
		{
			ExpectedCode:   400,
			ExpectedResult: &dto.UpdateTaskResponse{Id: 1, Name: "gogolook", Status: 3},
		},
		{
			ExpectedCode:   404,
			ExpectedResult: &dto.UpdateTaskResponse{Id: 9527, Name: "gogolook", Status: 0},
		},
	}
	for _, test := range tests {
		data := &dto.UpdateTaskRequest{Id: test.ExpectedResult.Id, Name: &test.ExpectedResult.Name, Status: &test.ExpectedResult.Status}
		body, err := json.Marshal(data)
		assert.Equal(t, nil, err)
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/task/%d", test.ExpectedResult.Id), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		routes.ServeHTTP(response, req)
		assert.Equal(t, test.ExpectedCode, response.Code)
	}
	// teardown
	err = repo.Delete(&task)
	assert.Equal(t, nil, err)
}

func TestDeleteTasks(t *testing.T) {
	setup()
	// create task to update
	task := models.Tasks{Id: 1, Name: "test", Status: 0}
	err := repo.Insert(&task)
	assert.Equal(t, nil, err)
	id := 1
	errId := 0
	notFoundId := 9527
	tests := []struct {
		ExpectedCode   int
		ExpectedResult *dto.DeleteTaskRequest
	}{
		{
			ExpectedCode:   200,
			ExpectedResult: &dto.DeleteTaskRequest{Id: &id},
		},
		{
			ExpectedCode:   400,
			ExpectedResult: &dto.DeleteTaskRequest{Id: &errId},
		},
		{
			ExpectedCode:   404,
			ExpectedResult: &dto.DeleteTaskRequest{Id: &notFoundId},
		},
	}
	for _, test := range tests {
		body, err := json.Marshal(test.ExpectedResult)
		assert.Equal(t, nil, err)
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/task/%d", test.ExpectedResult.Id), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		routes.ServeHTTP(response, req)
		assert.Equal(t, test.ExpectedCode, response.Code)
	}
	// teardown
	err = repo.Delete(&task)
	assert.Equal(t, nil, err)

}
