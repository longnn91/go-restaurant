package database

import (
	"gogo/common"
	"gogo/modules/order/model"
)

func (s *sqlStore) CreateOrder(data *model.OrderCreation) (*model.OrderCreation, error) {
	if err := s.db.Create(&data); err.Error != nil {
		return nil, common.ErrDB(err.Error)
	}

	return data, nil
}

func (s *sqlStore) GetOrders(paging *common.Paging) ([]model.Order, error) {
	var data []model.Order

	db := s.db

	if err := db.Table(model.Order{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *sqlStore) GetOrderById(id int) (*model.Order, error) {
	var data model.Order

	if err := s.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) UpdateOrder(id int, data *model.OrderUpdate) (*model.Order, error) {

	var newData model.Order
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	// Query the updated record to get the latest data
	if err := s.db.Where("id = ?", id).First(&newData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (s *sqlStore) DeleteOrder(id int) (int, error) {
	if err := s.db.Where("id = ?", id).Delete(&model.Order{}).Error; err != nil {
		return id, err
	}

	return id, nil
}
