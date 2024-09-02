package service

import (
	"gogo/common"
	"gogo/modules/category/model"
)

type CategoryActions interface {
	CreateCategory(data *model.CategoryCreation) (*model.CategoryCreation, error)
	GetCategories(paging *common.Paging) ([]model.Category, error)
	GetCategoryById(id int) (*model.Category, error)
	UpdateCategory(id int, data *model.CategoryUpdate) (*model.Category, error)
	DeleteCategory(id int) (int, error)
}

type categoryService struct {
	modelActions CategoryActions
}

func (service *categoryService) CreateCategory(data *model.CategoryCreation) (*model.CategoryCreation, error) {
	new, err := service.modelActions.CreateCategory(data)
	if err != nil {
		return nil, err
	}

	return new, nil
}

func (service *categoryService) GetCategories(paging *common.Paging) ([]model.Category, error) {
	data, err := service.modelActions.GetCategories(paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *categoryService) GetCategoryById(id int) (*model.Category, error) {
	data, err := service.modelActions.GetCategoryById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *categoryService) UpdateCategory(id int, data *model.CategoryUpdate) (*model.Category, error) {
	newData, err := service.modelActions.UpdateCategory(id, data)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (service *categoryService) DeleteCategory(id int) (int, error) {
	if id, err := service.modelActions.DeleteCategory(id); err != nil {
		return id, err
	}
	return id, nil
}

func GetCategoryService(modelActions CategoryActions) *categoryService {
	return &categoryService{modelActions: modelActions}
}
