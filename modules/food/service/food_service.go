package service

import (
	"gogo/common"
	"gogo/modules/food/model"
)

type FoodActions interface {
	CreateFood(data *model.FoodCreation) (*model.FoodCreation, error)
	GetFoods(paging *common.Paging) ([]model.Food, error)
	GetFoodById(id int) (*model.Food, error)
	UpdateFood(id int, data *model.FoodUpdate) (*model.Food, error)
	DeleteFood(id int) (int, error)
}

type foodService struct {
	modelActions FoodActions
}

func (service *foodService) CreateFood(data *model.FoodCreation) (*model.FoodCreation, error) {
	new, err := service.modelActions.CreateFood(data)
	if err != nil {
		return nil, err
	}

	return new, nil
}

func (service *foodService) GetFoods(paging *common.Paging) ([]model.Food, error) {
	data, err := service.modelActions.GetFoods(paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *foodService) GetFoodById(id int) (*model.Food, error) {
	data, err := service.modelActions.GetFoodById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *foodService) UpdateFood(id int, data *model.FoodUpdate) (*model.Food, error) {
	newData, err := service.modelActions.UpdateFood(id, data)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (service *foodService) DeleteFood(id int) (int, error) {
	if id, err := service.modelActions.DeleteFood(id); err != nil {
		return id, err
	}
	return id, nil
}

func GetFoodService(modelActions FoodActions) *foodService {
	return &foodService{modelActions: modelActions}
}
