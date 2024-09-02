package service

import (
	"gogo/common"
	"gogo/modules/menu/model"
)

type MenuActions interface {
	CreateMenu(data *model.MenuCreation) (*model.MenuCreation, error)
	GetMenus(paging *common.Paging) ([]model.Menu, error)
	GetMenuById(id int) (*model.Menu, error)
	UpdateMenu(id int, data *model.MenuUpdate) (*model.Menu, error)
	DeleteMenu(id int) (int, error)
}

type menuService struct {
	modelActions MenuActions
}

func (service *menuService) CreateMenu(data *model.MenuCreation) (*model.MenuCreation, error) {
	new, err := service.modelActions.CreateMenu(data)
	if err != nil {
		return nil, err
	}

	return new, nil
}

func (service *menuService) GetMenus(paging *common.Paging) ([]model.Menu, error) {
	data, err := service.modelActions.GetMenus(paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *menuService) GetMenuById(id int) (*model.Menu, error) {
	data, err := service.modelActions.GetMenuById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *menuService) UpdateMenu(id int, data *model.MenuUpdate) (*model.Menu, error) {
	newData, err := service.modelActions.UpdateMenu(id, data)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (service *menuService) DeleteMenu(id int) (int, error) {
	if id, err := service.modelActions.DeleteMenu(id); err != nil {
		return id, err
	}
	return id, nil
}

func GetMenuService(modelActions MenuActions) *menuService {
	return &menuService{modelActions: modelActions}
}
