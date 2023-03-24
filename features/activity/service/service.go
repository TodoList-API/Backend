package service

import (
	"TodoApp/features/activity"
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
	res, err := as.qry.Create(newActivity)

	if err != nil {
		return activity.Core{}, err

	}

	return res, nil
}

func (as *activityService) Update(activityID uint, updatedActivity activity.Core) (activity.Core, error) {
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
func (as *activityService) Delete(activityID uint) error {
	err := as.qry.Delete(activityID)

	if err != nil {
		return err

	}
	return nil
}
