package handler

import (
	"TodoApp/features/todos"
	"TodoApp/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	srv todos.TodoService
}

func New(srv todos.TodoService) todos.TodoHandler {
	return &todoHandler{
		srv: srv,
	}
}

func (th *todoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newTodo := TodoRequest{}

		if err := c.Bind(&newTodo); err != nil {
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		res, err := th.srv.Create(RequestToCore(newTodo))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("Success", "Success", ToResponse(res)))
	}
}

func (th *todoHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoID, err := strconv.Atoi(c.Param("todo_id"))

		if err != nil || todoID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "invalid todo id",
			})
		}

		updatedTodo := TodoRequest{}
		if err := c.Bind(&updatedTodo); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input format")
		}

		res, err := th.srv.Update(uint(todoID), RequestToCore(updatedTodo))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Success", "Success", ToResponse(res)))
	}
}

func (th *todoHandler) ListTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		query, _ := strconv.Atoi(c.QueryParam("activity_group_id"))

		res, err := th.srv.ListTodo(uint(query))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Success", "Success", ToListResponse(res)))
	}
}

func (th *todoHandler) GetTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoID, err := strconv.Atoi(c.Param("todo_id"))
		if err != nil || todoID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "invalid todo id",
			})
		}

		res, err := th.srv.GetTodo(uint(todoID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Success", "Success", ToResponse(res)))
	}
}

func (th *todoHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoID, err := strconv.Atoi(c.Param("todo_id"))
		if err != nil || todoID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "invalid todo id",
			})
		}

		err = th.srv.Delete(uint(todoID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.DeleteReponse("Success", "Success"))
	}
}
