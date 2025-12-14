package handler

import (
	"context"
	"inventory-system/model"
	"inventory-system/service"
)

// HandlerCategory provides category-related handlers.
type HandlerCategory struct {
	Service service.ServiceCategoryInterface
}

// HandlerCategoryInterface defines the methods for category handlers.
type HandlerCategoryInterface interface {
	GetAllCategory() ([]*model.Category, error)
	CreateCategory(model.Category) error
	GetCategoryByID(id int) (*model.Category, error)
	UpdateCategory(id int, category *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

// NewHandlerCategory creates a new instance of HandlerCategory.
func NewHandlerCategory(service service.ServiceCategoryInterface) HandlerCategory {
	return HandlerCategory{Service: service}
}

// GetAllCategory retrieves all categories using the service.
func (h *HandlerCategory) GetAllCategory() ([]*model.Category, error) {
	category, err := h.Service.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return category, nil
}

// CreateCategory adds a new category using the service.
func (h *HandlerCategory) CreateCategory(category *model.Category) error {
	err := h.Service.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

// GetCategoryByID retrieves a category by its ID using the service.
func (h *HandlerCategory) GetCategoryByID(id int) (*model.Category, error) {
	category, err := h.Service.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// UpdateCategory updates an existing category using the service.
func (h *HandlerCategory) UpdateCategory(id int, category *model.Category) (*model.Category, error) {
	updatedCategory, err := h.Service.UpdateCategory(id, category)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}

// DeleteCategory deletes a category by its ID using the service.
func (h *HandlerCategory) DeleteCategory(ctx context.Context, id int) error {
	err := h.Service.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}
	return nil

}
