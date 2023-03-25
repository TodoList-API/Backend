package data

import (
	"TodoApp/features/todos"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type todoData struct {
	db *gorm.DB
}

func New(db *gorm.DB) todos.TodoData {
	return &todoData{
		db: db,
	}
}

func (td *todoData) Create(newTodo todos.Core) (todos.Core, error) {

	data := CoreToData(newTodo)
	err := td.db.Create(&data).Error

	if err != nil {
		return todos.Core{}, err
	}

	return DataToCore(data), nil
}

func (td *todoData) Update(todoID uint, updatedTodo todos.Core) (todos.Core, error) {
	data := CoreToData(updatedTodo)

	qry := td.db.Where("id = ?", todoID).Updates(&data).Find(&data)

	if qry.RowsAffected <= 0 {
		return todos.Core{}, errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		return todos.Core{}, err
	}

	return DataToCore(data), nil
}
func (td *todoData) ListTodo(queryParam uint) ([]todos.Core, error) {
	listTodo := []Todos{}

	if queryParam <= 0 {
		err := td.db.Find(&listTodo).Error

		if err != nil {
			return []todos.Core{}, err

		}
	} else {
		err := td.db.Where("activity_group_id = ?", queryParam).Find(&listTodo).Error

		if err != nil {
			return []todos.Core{}, err
		}
	}

	return ListDataToCore(listTodo), nil
}
func (td *todoData) GetTodo(todoID uint) (todos.Core, error) {
	myTodo := Todos{}

	qry := td.db.Where("id = ?", todoID).Find(&myTodo)

	if err := qry.Error; err != nil {
		return todos.Core{}, err
	} else if qry.RowsAffected <= 0 {
		msg := fmt.Sprintf("Todo with ID %d Not Found", todoID)
		return todos.Core{}, errors.New(msg)
	}
	return DataToCore(myTodo), nil
}
func (td *todoData) Delete(todoID uint) (todos.Core, error) {
	record := Todos{}

	err := td.db.Where("id = ?", todoID).Delete(&record).Error

	if err != nil {
		return todos.Core{}, err

	}

	return DataToCore(record), nil
}
