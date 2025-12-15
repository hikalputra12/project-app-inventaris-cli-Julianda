package service

import (
	"context"
	"inventory-system/model"
	"inventory-system/repository"
)

// ServiceCategory provides category-related services.
type ServiceCategory struct {
	repo repository.RepositoryCategoryInterface
}

// ServiceCategoryInterface defines the methods for category services.
type ServiceCategoryInterface interface {
	GetAllCategory() ([]*model.Category, error)
	CreateCategory(category *model.Category) error
	GetCategoryByID(id int) (*model.Category, error)
	UpdateCategory(id int, category *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

// NewServiceCategory creates a new instance of ServiceCategory.
func NewServiceCategory(repo repository.RepositoryCategoryInterface) ServiceCategory {
	return ServiceCategory{repo: repo}
}

// GetAllCategory retrieves all categories using the repository.
func (s *ServiceCategory) GetAllCategory() ([]*model.Category, error) {
	category, err := s.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return category, nil
}

// CreateCategory adds a new category using the repository.
func (s *ServiceCategory) CreateCategory(category *model.Category) error {
	err := s.repo.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceCategory) GetCategoryByID(id int) (*model.Category, error) {
	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *ServiceCategory) UpdateCategory(id int, category *model.Category) (*model.Category, error) {
	updatedCategory, err := s.repo.UpdateCategory(id, category)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}

func (s *ServiceCategory) DeleteCategory(ctx context.Context, id int) error {
	err := s.repo.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
