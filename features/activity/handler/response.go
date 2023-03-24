package handler

import (
	"TodoApp/features/activity"
	"time"
)

type ActivityResponse struct {
	ID        uint      `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

func ToResponse(core activity.Core) ActivityResponse {
	return ActivityResponse{
		ID:        core.ID,
		Title:     core.Title,
		Email:     core.Email,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func ToListResponse(core []activity.Core) []ActivityResponse {
	var listActivity []ActivityResponse
	for _, act := range core {
		listActivity = append(listActivity, ToResponse(act))
	}

	return listActivity
}
