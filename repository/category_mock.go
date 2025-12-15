package repository

import (
	"context"
	"inventory-system/model"

	"github.com/stretchr/testify/mock"
)

// RepositoryCategoryMock implementasi dari mock mengambil dari interface di repo asli
type RepositoryCategoryMock struct {
	mock.Mock
}

//	GetAllCategory() ([]*model.Category, error)
//CreateCategory(category *model.Category) error
//GetCategoryByID(id int) (*model.Category, error)
//UpdateCategory(id int, category *model.Category) (*model.Category, error)
//DeleteCategory(ctx context.Context, id int) error

// buat fungsi buatan
func (m *RepositoryCategoryMock) GetAllCategory() ([]*model.Category, error) {
	//buat argument
	args := m.Called()      // panggil fungsi
	if args.Get(0) == nil { //data yang di masukkan tidak ada
		return nil, args.Error(1) //jika tidak ada maka tampilkan error
	}

	//jika ada maka tampilkan datanya
	return args.Get(0).([]*model.Category), args.Error(1)
}

// buat fungsi buatan
func (m *RepositoryCategoryMock) CreateCategory(category *model.Category) error {
	//buat argument
	args := m.Called(category) // panggil fungsi
	return args.Error(0)       //jika tidak ada maka tampilkan error
}

func (m *RepositoryCategoryMock) GetCategoryByID(id int) (*model.Category, error) {
	args := m.Called(id)
	err := args.Error(1)

	// Ambil Objek dengan Pengecekan Defensif (untuk menghindari panic pada Get(0))
	// Jika nilai yang dikembalikan nil, kita kembalikan nil pointer
	if args.Get(0) == nil {
		return nil, err // Kembalikan nil *model.Category dan error yang aman (err)
	}

	//Konversi Objek (Ini aman karena kita sudah cek nil)
	return args.Get(0).(*model.Category), err
}
func (m *RepositoryCategoryMock) UpdateCategory(id int, category *model.Category) (*model.Category, error) {
	//buat argument
	args := m.Called(id, category) // panggil fungsi
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} //jika tidak ada maka tampilkan datanya
	return args.Get(0).(*model.Category), args.Get(1).(error)

}

func (m *RepositoryCategoryMock) DeleteCategory(ctx context.Context, id int) error {
	//buat argument
	args := m.Called(ctx, id) // panggil fungsi
	return args.Error(1)

}
