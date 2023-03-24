package service

import (
	"TodoApp/features/todos"
	"errors"
)

type todoService struct {
	qry todos.TodoData
}

func New(td todos.TodoData) todos.TodoService {
	return &todoService{
		qry: td,
	}
}

func (ts *todoService) Create(newTodo todos.Core) (todos.Core, error) {

	if newTodo.Title == "" {
		return todos.Core{}, errors.New("title cannot be null")
	}

	res, err := ts.qry.Create(newTodo)

	if err != nil {
		return todos.Core{}, err

	}

	return res, nil
}

func (ts *todoService) Update(todoID uint, updatedTodo todos.Core) (todos.Core, error) {

	_, err := ts.qry.GetTodo(todoID)

	if err != nil {
		return todos.Core{}, err

	}

	if updatedTodo.Title == "" {
		return todos.Core{}, errors.New("title cannot be null")
	}

	res, err := ts.qry.Update(todoID, updatedTodo)

	if err != nil {
		return todos.Core{}, err

	}

	return res, nil
}
func (ts *todoService) ListTodo() ([]todos.Core, error) {

	res, err := ts.qry.ListTodo()

	if err != nil {
		return []todos.Core{}, err
	}

	return res, nil

}
func (ts *todoService) GetTodo(todoID uint) (todos.Core, error) {

	res, err := ts.qry.GetTodo(todoID)

	if err != nil {
		return todos.Core{}, err

	}

	return res, nil
}
func (ts *todoService) Delete(todoID uint) error {

	_, err := ts.qry.GetTodo(todoID)

	if err != nil {
		return err

	}
	err = ts.qry.Delete(todoID)

	if err != nil {
		return err

	}
	return nil
}
