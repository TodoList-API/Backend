package main

import (
	"TodoApp/config"
	"log"

	actData "TodoApp/features/activity/data"
	actHandler "TodoApp/features/activity/handler"
	actService "TodoApp/features/activity/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	activityData := actData.New(db)
	activityService := actService.New(activityData)
	activityHandler := actHandler.New(activityService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/activity-groups", activityHandler.Create())
	e.PATCH("/activity-groups/:activity_id", activityHandler.Update())
	e.GET("/activity-groups", activityHandler.ListActivity())
	e.POST("/activity-groups/:activity_id", activityHandler.GetActivity())
	e.DELETE("/activity-groups/:activity_id", activityHandler.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}

}
