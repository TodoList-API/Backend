package service

import (
	"TodoApp/features/activity"
	"errors"
)

type activityService struct {
	qry activity.ActivityData
}

func New(ad activity.ActivityData) activity.ActivityService {
	return &activityService{
		qry: ad,
	}
}

func (as *activityService) Create(newActivity activity.Core) (activity.Core, error) {

	if newActivity.Title == "" {
		return activity.Core{}, errors.New("title cannot be null")
	}

	res, err := as.qry.Create(newActivity)

	if err != nil {
		return activity.Core{}, err

	}

	return res, nil
}

func (as *activityService) Update(activityID uint, updatedActivity activity.Core) (activity.Core, error) {

	_, err := as.qry.GetActivity(activityID)

	if err != nil {
		return activity.Core{}, err

	}

	if updatedActivity.Title == "" {
		return activity.Core{}, errors.New("title cannot be null")
	}

	res, err := as.qry.Update(activityID, updatedActivity)

	if err != nil {
		return activity.Core{}, err

	}

	return res, nil
}
func (as *activityService) ListActivity() ([]activity.Core, error) {

	res, err := as.qry.ListActivity()

	if err != nil {
		return []activity.Core{}, err
	}

	return res, nil

}
func (as *activityService) GetActivity(activityID uint) (activity.Core, error) {

	res, err := as.qry.GetActivity(activityID)

	if err != nil {
		return activity.Core{}, err

	}

	return res, nil
}
func (as *activityService) Delete(activityID uint) (activity.Core, error) {

	_, err := as.qry.GetActivity(activityID)

	if err != nil {
		return activity.Core{}, err

	}
	res, err := as.qry.Delete(activityID)

	if err != nil {
		return activity.Core{}, err

	}
	return res, nil
}
