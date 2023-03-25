package handler

import (
	"TodoApp/features/activity"
	"TodoApp/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type activityHandler struct {
	srv activity.ActivityService
}

func New(srv activity.ActivityService) activity.ActivityHandler {
	return &activityHandler{
		srv: srv,
	}
}

func (ah *activityHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newAct := CreateActivityRequest{}

		if err := c.Bind(&newAct); err != nil {
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		res, err := ah.srv.Create(CreateRequestToCore(newAct))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("Success", "Success", ToResponse(res)))
	}
}

func (ah *activityHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		activityID, err := strconv.Atoi(c.Param("activity_id"))

		if err != nil || activityID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "invalid activity id",
			})
		}

		updatedAct := UpdateActivityRequest{}
		if err := c.Bind(&updatedAct); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input format")
		}

		res, err := ah.srv.Update(uint(activityID), UpdateRequestToCore(updatedAct))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Success", "Success", ToResponse(res)))
	}
}
func (ah *activityHandler) ListActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ah.srv.ListActivity()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Success", "Success", ToListResponse(res)))
	}
}

func (ah *activityHandler) GetActivity() echo.HandlerFunc {
	return func(c echo.Context) error {
		activityID, err := strconv.Atoi(c.Param("activity_id"))
		if err != nil || activityID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "invalid activity id",
			})
		}

		res, err := ah.srv.GetActivity(uint(activityID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Success", "Success", ToResponse(res)))
	}
}
func (ah *activityHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		activityID, err := strconv.Atoi(c.Param("activity_id"))
		if err != nil || activityID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "invalid activity id",
			})
		}

		err = ah.srv.Delete(uint(activityID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.DeleteReponse("Success", "Success"))
	}
}
