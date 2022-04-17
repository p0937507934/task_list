// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"tasks_list/driver"
	"tasks_list/internal/tasks"
)

// Injectors from wire.go:

func BuildTaskService() tasks.ITaskService {
	db := driver.InitGorm()
	iTaskRepository := tasks.NewTasksRepository(db)
	iTaskService := tasks.NewTaskService(iTaskRepository)
	return iTaskService
}

// wire.go:

var gormSet = wire.NewSet(driver.InitGorm)

var taskServiceSet = wire.NewSet(tasks.NewTaskService, tasks.NewTasksRepository)
