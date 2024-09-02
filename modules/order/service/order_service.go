package service

import (
	"gogo/common"
	"gogo/modules/order/model"
)

type OrderActions interface {
	CreateOrder(data *model.OrderCreation) (*model.OrderCreation, error)
	GetOrders(paging *common.Paging) ([]model.Order, error)
	GetOrderById(id int) (*model.Order, error)
	UpdateOrder(id int, data *model.OrderUpdate) (*model.Order, error)
	DeleteOrder(id int) (int, error)
}

type orderService struct {
	modelActions OrderActions
}

func (service *orderService) CreateOrder(data *model.OrderCreation) (*model.OrderCreation, error) {
	new, err := service.modelActions.CreateOrder(data)
	if err != nil {
		return nil, err
	}

	return new, nil
}

func (service *orderService) GetOrders(paging *common.Paging) ([]model.Order, error) {
	data, err := service.modelActions.GetOrders(paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *orderService) GetOrderById(id int) (*model.Order, error) {
	data, err := service.modelActions.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *orderService) UpdateOrder(id int, data *model.OrderUpdate) (*model.Order, error) {
	newData, err := service.modelActions.UpdateOrder(id, data)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (service *orderService) DeleteOrder(id int) (int, error) {
	if id, err := service.modelActions.DeleteOrder(id); err != nil {
		return id, err
	}
	return id, nil
}

func GetOrderService(modelActions OrderActions) *orderService {
	return &orderService{modelActions: modelActions}
}
