package tasks

import (
	"tasks_list/models"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFindAll(t *testing.T) {
	Init()
	tasks, err := repo.FindAll()
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(tasks))
}

func TestInsertTask(t *testing.T) {
	Init()
	task := models.Tasks{Name: "吃飯", Status: 0}
	err := repo.Insert(&task)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, task)
	// teardown
	repo.Delete(&task)
}

func TestUpdateTask(t *testing.T) {
	Init()
	// 新增一筆資料
	task := models.Tasks{Name: "吃飯完成", Status: 0}
	_ = repo.Insert(&task)
	// 更新
	task.Name = "吃飽了"
	task.Status = 1
	err := repo.Update(&task)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, task)
	// teardown
	repo.Delete(&task)
}

func TestDeleteTask(t *testing.T) {
	Init()
	// 新增一筆資料
	task := models.Tasks{Name: "吃飯完成", Status: 0}
	_ = repo.Insert(&task)
	// 刪除
	err := repo.Delete(&task)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, task)
}
