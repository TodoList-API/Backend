package handler

import (
	"TodoApp/features/activity"
	"time"
)

type ActivityResponse struct {
	ID        uint
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
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
