package data

import (
	"TodoApp/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string
	Email string
}

func DataToCore(data Activity) activity.Core {
	return activity.Core{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func CoreToData(core activity.Core) Activity {
	return Activity{
		Model: gorm.Model{ID: core.ID, CreatedAt: core.CreatedAt, UpdatedAt: core.UpdatedAt},
		Title: core.Title,
		Email: core.Email,
	}
}

func ListDataToCore(data []Activity) []activity.Core {
	var listActivity []activity.Core
	for _, act := range data {
		listActivity = append(listActivity, DataToCore(act))
	}

	return listActivity
}

func ListCoreToData(data []activity.Core) []Activity {
	var listActivity []Activity
	for _, act := range data {
		listActivity = append(listActivity, CoreToData(act))
	}

	return listActivity
}
