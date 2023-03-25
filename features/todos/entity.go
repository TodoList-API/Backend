package todos

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	Title           string
	ActivityGroupId uint
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TodoHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	ListTodo() echo.HandlerFunc
	GetTodo() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type TodoService interface {
	Create(newTodo Core) (Core, error)
	Update(todoID uint, updatedTodo Core) (Core, error)
	ListTodo(queryParam uint) ([]Core, error)
	GetTodo(todoID uint) (Core, error)
	Delete(todoID uint) (Core, error)
}

type TodoData interface {
	Create(newTodo Core) (Core, error)
	Update(todoID uint, updatedTodo Core) (Core, error)
	ListTodo(queryParam uint) ([]Core, error)
	GetTodo(todoID uint) (Core, error)
	Delete(todoID uint) (Core, error)
}
