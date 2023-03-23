package activity

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ActivityHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	ListActivity() echo.HandlerFunc
	GetActivity() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ActivityService interface {
	Create(newActivity Core) (Core, error)
	Update(activityID uint, updatedActivity Core) (Core, error)
	ListActivity() ([]Core, error)
	GetActivity(activityID uint) (Core, error)
	Delete(activityID uint) error
}

type ActivityData interface {
	Create(newActivity Core) (Core, error)
	Update(activityID uint, updatedActivity Core) (Core, error)
	ListActivity() ([]Core, error)
	GetActivity(activityID uint) (Core, error)
	Delete(activityID uint) error
}
