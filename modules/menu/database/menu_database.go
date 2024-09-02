package database

import (
	"gogo/common"
	"gogo/modules/menu/model"
)

func (s *sqlStore) CreateMenu(data *model.MenuCreation) (*model.MenuCreation, error) {
	if err := s.db.Create(&data); err.Error != nil {
		return nil, common.ErrDB(err.Error)
	}

	return data, nil
}

func (s *sqlStore) GetMenus(paging *common.Paging) ([]model.Menu, error) {
	var data []model.Menu

	db := s.db

	if err := db.Table(model.Menu{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *sqlStore) GetMenuById(id int) (*model.Menu, error) {
	var data model.Menu

	if err := s.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) UpdateMenu(id int, data *model.MenuUpdate) (*model.Menu, error) {

	var newData model.Menu
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	// Query the updated record to get the latest data
	if err := s.db.Where("id = ?", id).First(&newData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (s *sqlStore) DeleteMenu(id int) (int, error) {
	if err := s.db.Where("id = ?", id).Delete(&model.Menu{}).Error; err != nil {
		return id, err
	}

	return id, nil
}
