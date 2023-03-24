package handler

import "TodoApp/features/activity"

type CreateActivityRequest struct {
	Title string `json:"title" form:"title"`
	Email string `json:"email" form:"email"`
}

type UpdateActivityRequest struct {
	Title string `json:"title" form:"title"`
}

func CreateRequestToCore(request CreateActivityRequest) activity.Core {

	return activity.Core{
		Title: request.Title,
		Email: request.Email,
	}

}

func UpdateRequestToCore(request UpdateActivityRequest) activity.Core {

	return activity.Core{
		Title: request.Title,
	}

}
