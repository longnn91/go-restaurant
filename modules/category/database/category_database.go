package database

import (
	"gogo/common"
	"gogo/modules/category/model"
)

func (s *sqlStore) CreateCategory(data *model.CategoryCreation) (*model.CategoryCreation, error) {
	if err := s.db.Create(&data); err.Error != nil {
		return nil, common.ErrDB(err.Error)
	}

	return data, nil
}

func (s *sqlStore) GetCategories(paging *common.Paging) ([]model.Category, error) {
	var data []model.Category

	db := s.db

	if err := db.Table(model.Category{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *sqlStore) GetCategoryById(id int) (*model.Category, error) {
	var data model.Category

	if err := s.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) UpdateCategory(id int, data *model.CategoryUpdate) (*model.Category, error) {

	var newData model.Category
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	// Query the updated record to get the latest data
	if err := s.db.Where("id = ?", id).First(&newData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (s *sqlStore) DeleteCategory(id int) (int, error) {
	if err := s.db.Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return id, err
	}

	return id, nil
}
