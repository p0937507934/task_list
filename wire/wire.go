//go:build wireinject
// +build wireinject

package wire

import (
	"tasks_list/driver"
	"tasks_list/internal/tasks"

	"github.com/google/wire"
)

var gormSet = wire.NewSet(driver.InitGorm)

var taskServiceSet = wire.NewSet(tasks.NewTaskService, tasks.NewTasksRepository)

func BuildTaskService() tasks.ITaskService {
	wire.Build(taskServiceSet, gormSet)
	return &tasks.TaskService{}
}
