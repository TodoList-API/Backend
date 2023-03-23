package handler

import "TodoApp/features/activity"

type CreateActivityRequest struct {
	Title string
	Email string
}

type UpdateActivityRequest struct {
	Title string
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
