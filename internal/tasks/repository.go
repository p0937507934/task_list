package tasks

import (
	"tasks_list/models"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	Insert(*models.Tasks) error
	FindAll() ([]*models.Tasks, error)
	Update(*models.Tasks) error
	Delete(*models.Tasks) error
	FindOne(*models.Tasks) error
}

type TasksRepository struct {
	orm *gorm.DB
}

func NewTasksRepository(orm *gorm.DB) ITaskRepository {
	return &TasksRepository{orm: orm}
}

func (r *TasksRepository) Insert(task *models.Tasks) error {
	if result := r.orm.Create(task); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TasksRepository) FindAll() ([]*models.Tasks, error) {
	tasks := make([]*models.Tasks, 0)
	if result := r.orm.Find(&tasks); result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *TasksRepository) Update(task *models.Tasks) error {
	if result := r.orm.Updates(&task); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TasksRepository) Delete(task *models.Tasks) error {
	if result := r.orm.Delete(&task); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TasksRepository) FindOne(task *models.Tasks) error {
	result := r.orm.First(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
