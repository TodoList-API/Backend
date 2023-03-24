package handler

import (
	"TodoApp/features/todos"
	"time"
)

type TodoResponse struct {
	ID              uint      `json:"id" form:"id"`
	Title           string    `json:"title" form:"title"`
	ActivityGroupId uint      `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool      `json:"is_active" form:"is_active"`
	Priority        string    `json:"priority" form:"priority"`
	CreatedAt       time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" form:"updatedAt"`
}

func ToResponse(core todos.Core) TodoResponse {
	return TodoResponse{
		ID:              core.ID,
		Title:           core.Title,
		ActivityGroupId: core.ActivityGroupId,
		IsActive:        core.IsActive,
		Priority:        core.Priority,
		CreatedAt:       core.CreatedAt,
		UpdatedAt:       core.UpdatedAt,
	}
}

func ToListResponse(core []todos.Core) []TodoResponse {
	var listTodo []TodoResponse
	for _, act := range core {
		listTodo = append(listTodo, ToResponse(act))
	}

	return listTodo
}
