package service

import (
	"inventory-system/model"
	"inventory-system/repository"
)

type ServiceInventory struct {
	repository repository.RepositoryInventoryInterface
}

type ServiceInventoryInterface interface {
	ItemsMoreThan100Days() ([]model.TotalUsageDays, error)
	TotalInvesmentValue() ([]model.TotalInvestmentValue, error)
	InvesmentAndDepreciationValueByID(id int) (*model.Depreciation, error)
	FindInventoryByName(name string) (*model.Inventory, error)
}

func NewServiceInventory(repo repository.RepositoryInventoryInterface) ServiceInventory {
	return ServiceInventory{
		repository: repo,
	}
}

func (s *ServiceInventory) ItemsMoreThan100Days() ([]model.TotalUsageDays, error) {
	inventories, err := s.repository.ItemsMoreThan100Days()
	if err != nil {
		return nil, err
	}
	return inventories, err
}

func (s *ServiceInventory) TotalInvesmentValue() ([]model.TotalInvestmentValue, error) {
	inventories, err := s.repository.TotalInvesmentValue()
	if err != nil {
		return nil, err
	}
	return inventories, err
}

func (s *ServiceInventory) InvesmentAndDepreciationValueByID(id int) (*model.Depreciation, error) {
	inventories, err := s.repository.InvesmentAndDepreciationValueByID(id)
	if err != nil {
		return nil, err
	}
	return inventories, err
}

func (s *ServiceInventory) FindInventoryByName(name string) (*model.Inventory, error) {
	inventories, err := s.repository.FindInventoryByName(name)
	if err != nil {
		return nil, err
	}
	return inventories, err
}
