package tasks

import (
	"tasks_list/dto"
	"tasks_list/models"
	custom_error "tasks_list/pkg/error"

	"gorm.io/gorm"
)

type ITaskService interface {
	CreateTask(*dto.CreateTaskRequest) (*dto.CreateTaskResponse, error)
	ListTasks() ([]*dto.ListTaskResponse, error)
	UpdateTask(*dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error)
	DeleteTask(*dto.DeleteTaskRequest) error
}

type TaskService struct {
	taskRepo ITaskRepository
}

func NewTaskService(taskRepo ITaskRepository) ITaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(req *dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	task := &models.Tasks{
		Name:   *req.Name,
		Status: req.Status,
	}
	err := s.taskRepo.Insert(task)
	res := &dto.CreateTaskResponse{
		Id:     task.Id,
		Name:   task.Name,
		Status: task.Status,
	}
	if err != nil {
		return res, custom_error.NewCustomError(500, err, "Create task error", req)
	}
	return res, nil
}

func (s *TaskService) ListTasks() ([]*dto.ListTaskResponse, error) {
	res := make([]*dto.ListTaskResponse, 0)
	tasks, err := s.taskRepo.FindAll()
	if err != nil {
		return nil, custom_error.NewCustomError(500, err, "Create task error", nil)
	}
	for _, task := range tasks {
		res = append(res, &dto.ListTaskResponse{
			Id:     task.Id,
			Name:   task.Name,
			Status: task.Status,
		})
	}
	return res, nil
}

func (s *TaskService) UpdateTask(req *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error) {
	task := &models.Tasks{
		Id: req.Id,
	}
	// check data exists
	err := s.taskRepo.FindOne(task)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custom_error.NewCustomError(404, err, "Update task error, data not exist", req)
		}
		if err != nil {
			return nil, custom_error.NewCustomError(500, err, "Update task error", req)
		}
	}
	task.Name = *req.Name
	task.Status = *req.Status
	err = s.taskRepo.Update(task)
	res := &dto.UpdateTaskResponse{
		Id:     task.Id,
		Name:   task.Name,
		Status: task.Status,
	}
	if err != nil {
		return nil, custom_error.NewCustomError(500, err, "Update task error", req)
	}
	return res, nil
}

func (s *TaskService) DeleteTask(req *dto.DeleteTaskRequest) error {
	// init db model
	task := &models.Tasks{
		Id: *req.Id,
	}

	// check data exists
	err := s.taskRepo.FindOne(task)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return custom_error.NewCustomError(404, err, "Delete task error, data not exist", req)
		}
		if err != nil {
			return custom_error.NewCustomError(500, err, "Delete task error", req)
		}
	}
	// if exists, do delete
	err = s.taskRepo.Delete(task)
	if err != nil {
		return custom_error.NewCustomError(500, err, "Delete task error", req)
	}
	return nil
}
