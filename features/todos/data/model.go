package data

import (
	"TodoApp/features/todos"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Title           string
	ActivityGroupId uint
	IsActive        bool   `gorm:"default:true"`
	Priority        string `gorm:"default:very-high"`
	Status          string
}

func DataToCore(data Todos) todos.Core {
	return todos.Core{
		ID:              data.ID,
		Title:           data.Title,
		ActivityGroupId: data.ActivityGroupId,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}

func CoreToData(core todos.Core) Todos {
	return Todos{
		Model:           gorm.Model{ID: core.ID, CreatedAt: core.CreatedAt, UpdatedAt: core.UpdatedAt},
		Title:           core.Title,
		ActivityGroupId: core.ActivityGroupId,
		IsActive:        core.IsActive,
		Priority:        core.Priority,
	}
}

func ListDataToCore(data []Todos) []todos.Core {
	var listTodos []todos.Core
	for _, todo := range data {
		listTodos = append(listTodos, DataToCore(todo))
	}

	return listTodos
}

func ListCoreToData(data []todos.Core) []Todos {
	var listTodos []Todos
	for _, todo := range data {
		listTodos = append(listTodos, CoreToData(todo))
	}

	return listTodos
}
