package service

import (
	"errors"
	"inventory-system/model"
	"inventory-system/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategory(t *testing.T) {
	//buat dua skenario
	//jika category ada
	t.Run("get_category", func(t *testing.T) {
		mockData := []*model.Category{
			{
				Model: model.Model{
					ID: 1,
				},
				Name:        "ATK",
				Description: "Atk is the best",
			},
		}
		//buat mock repository
		mockRepo := new(repository.RepositoryCategoryMock)
		//dan kemudian aktifkan fungsinya dengan mengambil data nya
		mockRepo.
			On("GetAllCategory").
			Return(mockData, nil)

		categoryService := NewServiceCategory(mockRepo)
		result, err := categoryService.GetAllCategory()

		assert.NoError(t, err)
		assert.Equal(t, mockData, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("not_get_all_category", func(t *testing.T) {

		//buat mock repository
		mockRepo := new(repository.RepositoryCategoryMock)
		mockRepo.
			On("GetAllCategory").
			Return(nil, errors.New("db error"))
		categoryService := NewServiceCategory(mockRepo)
		result, err := categoryService.GetAllCategory()

		assert.Error(t, err)  //pastikan err tidak nil
		assert.Nil(t, result) //pastikan hasilnya kosong
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateCategory(t *testing.T) {
	//buat dua skenario
	//jika category ada
	t.Run("success_create_new_category", func(t *testing.T) {
		newCategory := &model.Category{
			Name:        "ATK",
			Description: "Atk is the best",
		}
		//buat mock repository
		mockRepo := new(repository.RepositoryCategoryMock)
		//dan kemudian aktifkan fungsinya dengan mengambil data nya
		mockRepo.
			On("CreateCategory", newCategory).
			Return(nil)

		categoryService := NewServiceCategory(mockRepo)
		err := categoryService.CreateCategory(newCategory)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	t.Run("no_success_create_new_category", func(t *testing.T) {

		newCategory := &model.Category{
			Name:        "",
			Description: "",
		}
		//buat mock repository
		mockRepo := new(repository.RepositoryCategoryMock)
		mockRepo.
			On("CreateCategory", newCategory).
			Return(errors.New("name and description required"))
		categoryService := NewServiceCategory(mockRepo)
		err := categoryService.CreateCategory(newCategory)

		assert.Error(t, err) //pastikan err tidak nil
		mockRepo.AssertExpectations(t)
	})
}
func TestGetCategoryByID(t *testing.T) {
	//buat dua skenario
	t.Run("success_to_get_category", func(t *testing.T) {
		getCategory := &model.Category{
			Model: model.Model{
				ID: 1,
			},
			Name:        "ATK",
			Description: "Atk is the best",
		}
		//buat mock repository
		mockRepo := new(repository.RepositoryCategoryMock)
		//dan kemudian aktifkan fungsinya dengan mengambil data nya
		mockRepo.
			On("GetCategoryByID", 1).
			Return(getCategory, nil)

		categoryService := NewServiceCategory(mockRepo)
		result, err := categoryService.GetCategoryByID(1)
		assert.NoError(t, err)
		assert.Equal(t, getCategory, result)
		mockRepo.AssertExpectations(t)

	})
}
