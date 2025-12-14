package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"inventory-system/handler"
	"inventory-system/model"
	"inventory-system/utils"

	"github.com/olekukonko/tablewriter"
)

// function to display all categories
func GetAllCategory(handler handler.HandlerCategory) {
	categories, err := handler.GetAllCategory()
	if err != nil {
		fmt.Println("Error fetching categories:", err)
		return
	}

	fmt.Println("\nLIST OF CATEGORIES")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Category Name", "Description")

	for _, category := range categories {
		table.Append([]string{
			fmt.Sprintf("%d", category.ID),
			category.Name,
			category.Description,
		})
	}

	table.Render()
}

// function to create a new category
func CreateCategory(handler handler.HandlerCategory) {
	fmt.Println("\nCREATE NEW CATEGORY")

	nameInput := utils.ReadLine("Enter Category Name: ")
	descInput := utils.ReadLine("Enter Category Description: ")

	newCategory := &model.Category{
		Name:        nameInput,
		Description: descInput,
	}

	err := handler.CreateCategory(newCategory)
	if err != nil {
		fmt.Println("Error creating category:", err)
		return
	}

	fmt.Println("\nCATEGORY CREATED SUCCESSFULLY")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Category Name", "Description")

	table.Append([]string{
		fmt.Sprintf("%d", newCategory.ID),
		newCategory.Name,
		newCategory.Description,
	})

	table.Render()
}

// function to get category by ID
func GetCategoryByID(handler handler.HandlerCategory) {
	var id int
	fmt.Print("Enter Category ID: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	category, err := handler.GetCategoryByID(id)
	if err != nil {
		fmt.Println("Error fetching category:", err)
		return
	}

	fmt.Println("\nCATEGORY DETAILS")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Category Name", "Description")

	table.Append([]string{
		fmt.Sprintf("%d", id),
		category.Name,
		category.Description,
	})

	table.Render()
}

// function to update category by ID
func UpdateCategory(handler handler.HandlerCategory) {
	var id int
	fmt.Print("Enter Category ID to update: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	oldCategory, err := handler.GetCategoryByID(id)
	if err != nil {
		fmt.Println("Error fetching category:", err)
		return
	}

	nameInput := utils.ReadLine("Enter new Category Name (leave empty to keep old): ")
	descInput := utils.ReadLine("Enter new Category Description (leave empty to keep old): ")

	if nameInput == "" {
		nameInput = oldCategory.Name
	}
	if descInput == "" {
		descInput = oldCategory.Description
	}

	updatedCategory := &model.Category{
		Name:        nameInput,
		Description: descInput,
	}

	_, err = handler.UpdateCategory(id, updatedCategory)
	if err != nil {
		fmt.Println("Error updating category:", err)
		return
	}

	fmt.Println("\nCATEGORY UPDATED SUCCESSFULLY")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Category Name", "Description")

	table.Append([]string{
		fmt.Sprintf("%d", id),
		updatedCategory.Name,
		updatedCategory.Description,
	})

	table.Render()
}

// function to delete category by ID
func DeleteCategory(handler handler.HandlerCategory) {
	var id int
	fmt.Print("Enter Category ID to delete: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = handler.DeleteCategory(ctx, id)
	if err != nil {
		fmt.Println("Error deleting category:", err)
		return
	}

	fmt.Println("\nCATEGORY DELETED SUCCESSFULLY")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("Deleted Category ID")

	table.Append([]string{
		fmt.Sprintf("%d", id),
	})

	table.Render()
}
