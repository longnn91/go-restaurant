package database

import (
	"gogo/common"
	"gogo/modules/food/model"
)

func (s *sqlStore) CreateFood(data *model.FoodCreation) (*model.FoodCreation, error) {
	if err := s.db.Create(&data); err.Error != nil {
		return nil, common.ErrDB(err.Error)
	}

	return data, nil
}

func (s *sqlStore) GetFoods(paging *common.Paging) ([]model.Food, error) {
	var data []model.Food

	db := s.db

	if err := db.Table(model.Food{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *sqlStore) GetFoodById(id int) (*model.Food, error) {
	var data model.Food

	if err := s.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) UpdateFood(id int, data *model.FoodUpdate) (*model.Food, error) {

	var newData model.Food
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	// Query the updated record to get the latest data
	if err := s.db.Where("id = ?", id).First(&newData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (s *sqlStore) DeleteFood(id int) (int, error) {
	if err := s.db.Where("id = ?", id).Delete(&model.Food{}).Error; err != nil {
		return id, err
	}

	return id, nil
}
