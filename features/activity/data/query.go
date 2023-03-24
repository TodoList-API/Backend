package data

import (
	"TodoApp/features/activity"
	"errors"

	"gorm.io/gorm"
)

type activityData struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.ActivityData {
	return &activityData{
		db: db,
	}
}

func (ad *activityData) Create(newActivity activity.Core) (activity.Core, error) {

	data := CoreToData(newActivity)
	err := ad.db.Create(&data).Error

	if err != nil {
		return activity.Core{}, err
	}

	return DataToCore(data), nil
}

func (ad *activityData) Update(activityID uint, updatedActivity activity.Core) (activity.Core, error) {
	data := CoreToData(updatedActivity)

	qry := ad.db.Where("id = ?", activityID).Updates(&data)

	if qry.RowsAffected <= 0 {
		return activity.Core{}, errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		return activity.Core{}, err
	}

	return DataToCore(data), nil
}
func (ad *activityData) ListActivity() ([]activity.Core, error) {
	listActivity := []Activity{}

	err := ad.db.Find(&listActivity).Error

	if err != nil {
		return []activity.Core{}, err

	}

	return ListDataToCore(listActivity), nil
}
func (ad *activityData) GetActivity(activityID uint) (activity.Core, error) {
	myActivity := Activity{}

	err := ad.db.Where("id = ?", activityID).Find(&myActivity).Error

	if err != nil {
		return activity.Core{}, err
	}

	return DataToCore(myActivity), nil
}
func (ad *activityData) Delete(activityID uint) error {
	record := Activity{}

	err := ad.db.Where("id = ?", activityID).Delete(&record).Error

	if err != nil {
		return err

	}

	return nil
}
