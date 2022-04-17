package dto

// apireq
type CreateTaskRequest struct {
	Name   *string `json:"name" binding:"required"`
	Status int     `json:"status" binding:"oneof=0 1"`
}

type UpdateTaskRequest struct {
	Id     int     `json:"id" binding:"required"`
	Name   *string `json:"name" binding:"required"`
	Status *int    `json:"status" binding:"required,oneof=0 1"`
}

type DeleteTaskRequest struct {
	Id *int `json:"id" binding:"required,min=1"`
}

// apires

type ListTaskResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type CreateTaskResponse struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
	Id     int    `json:"id"`
}

type UpdateTaskResponse struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
	Id     int    `json:"id"`
}
