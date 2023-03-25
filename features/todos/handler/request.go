package handler

import (
	"TodoApp/features/todos"
)

type TodoRequest struct {
	Title           string `json:"title" form:"title"`
	ActivityGroupId uint   `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool   `json:"is_active" form:"is_active"`
	Priority        string `json:"priority" form:"priority"`
	Status          string `json:"status" form:"status"`
}

func RequestToCore(request TodoRequest) todos.Core {

	return todos.Core{
		Title:           request.Title,
		ActivityGroupId: request.ActivityGroupId,
		IsActive:        request.IsActive,
		Priority:        request.Priority,
	}

}
