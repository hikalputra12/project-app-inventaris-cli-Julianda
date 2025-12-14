package handler

import (
	"inventory-system/model"
	"inventory-system/service"
)

// HandlerInventory provides inventory-related handlers.
type HandlerInventory struct {
	service service.ServiceInventoryInterface
}

// HandlerInventoryInterface defines the methods for inventory handlers.
type HandlerInventoryInterface interface {
	ItemsMoreThan100Days() ([]model.TotalUsageDays, error)
	TotalInvesmentValue() ([]model.TotalInvestmentValue, error)
	InvesmentAndDepreciationValueByID(id int) (*model.Depreciation, error)
	FindInventoryByName(name string) (*model.Inventory, error)
}

// NewHandlerInventory creates a new instance of HandlerInventory.
func NewHandlerInventory(service service.ServiceInventoryInterface) HandlerInventory {
	return HandlerInventory{
		service: service,
	}
}

// ItemsMoreThan100Days retrieves inventory items used for more than 100 days using the service.
func (h *HandlerInventory) ItemsMoreThan100Days() ([]model.TotalUsageDays, error) {
	inventories, err := h.service.ItemsMoreThan100Days()
	if err != nil {
		return nil, err
	}
	return inventories, err
}

// TotalInvesmentValue retrieves the total investment value of inventory items using the service.
func (h *HandlerInventory) TotalInvesmentValue() ([]model.TotalInvestmentValue, error) {
	inventories, err := h.service.TotalInvesmentValue()
	if err != nil {
		return nil, err
	}
	return inventories, err
}

// InvesmentAndDepreciationValueByID retrieves investment and depreciation value by inventory ID using the service.
func (h *HandlerInventory) InvesmentAndDepreciationValueByID(id int) (*model.Depreciation, error) {
	inventories, err := h.service.InvesmentAndDepreciationValueByID(id)
	if err != nil {
		return nil, err
	}
	return inventories, err
}

// FindInventoryByName retrieves an inventory item by its name using the service.
func (h *HandlerInventory) FindInventoryByName(name string) (*model.Inventory, error) {
	inventories, err := h.service.FindInventoryByName(name)
	if err != nil {
		return nil, err
	}
	return inventories, err
}
