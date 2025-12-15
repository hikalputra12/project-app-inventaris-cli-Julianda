package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"inventory-system/handler"
	"inventory-system/model"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// untuk mendapatkan semua kategori
// ubah jadi contructor cobra
func GetAllCategoryCmd(handler handler.HandlerCategory) *cobra.Command {
	getAllCategory := &cobra.Command{
		Use:   "get-all-category",
		Short: "Get all categories",
		Run: func(cmd *cobra.Command, args []string) {
			ClearScreen()
			categories, err := handler.GetAllCategory()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Name", "Description"})

			for _, c := range categories {
				table.Append([]string{
					fmt.Sprintf("%d", c.ID),
					c.Name,
					c.Description,
				})
			}

			table.Render()
		},
	}
	return getAllCategory
}

// // function to create a new category
func CreateCategory(handler handler.HandlerCategory) *cobra.Command {
	var name string
	var description string
	createCategory := &cobra.Command{
		Use:   "create-category",
		Short: "Create a new category",
		Run: func(cmd *cobra.Command, args []string) {
			ClearScreen()
			fmt.Println("\nCREATE NEW CATEGORY")

			if name == "" || description == "" {
				fmt.Println("Error: --name and --desc are required")
				return
			}

			newCategory := &model.Category{
				Name:        name,
				Description: description,
			}

			err := handler.CreateCategory(newCategory)
			if err != nil {
				fmt.Println("Error creating category:", err)
				return
			}

			fmt.Println("\nCATEGORY CREATED SUCCESSFULLY")

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Category Name", "Description"})

			table.Append([]string{
				fmt.Sprintf("%d", newCategory.ID),
				newCategory.Name,
				newCategory.Description,
			})

			table.Render()
		},
	}
	//add falgs
	createCategory.Flags().StringVarP(&name, "name", "n", "", "category name")
	createCategory.Flags().StringVarP(&description, "desc", "d", "", "category desk")

	//memaksa user untuk wajib isi
	createCategory.MarkFlagRequired("name")
	createCategory.MarkFlagRequired("desc")

	return createCategory

}

// function to get category by ID
func GetCategoryByIDCmd(handler handler.HandlerCategory) *cobra.Command {
	var id int
	GetCategoryByID := &cobra.Command{
		Use:   "get-category-byID",
		Short: "To get category by ID",
		Run: func(cmd *cobra.Command, args []string) {
			categories, err := handler.GetAllCategory()
			if err != nil {
				fmt.Println("Error fetching categories:", err)
				return
			}
			if id < 1 {
				fmt.Println("Category ID not found.")
				return
			}
			found := false

			for _, rows := range categories {
				if id == rows.ID {
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Category ID not found")
				return
			}

			category, err := handler.GetCategoryByID(id)
			if err != nil {
				fmt.Println("Error fetching category:", err)
				return
			}

			fmt.Println("\nCATEGORY DETAILS")

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Category Name", "Description"})

			table.Append([]string{
				fmt.Sprintf("%d", id),
				category.Name,
				category.Description,
			})

			table.Render()

		},
	}
	//add flag
	GetCategoryByID.Flags().IntVarP(&id, "id", "i", 0, "category id")
	GetCategoryByID.MarkFlagRequired("id")

	return GetCategoryByID
}

// unction to update category by ID
func UpdateCategoryCmd(handler handler.HandlerCategory) *cobra.Command {
	var id int
	var name string
	var description string
	UpdateCategory := &cobra.Command{
		Use:   "updated-category-byID",
		Short: "To Updated Category By ID",
		Long:  `to updated a category and return new name and description By ID input  By user`,
		Run: func(cmd *cobra.Command, args []string) {
			ClearScreen()
			categories, _ := handler.GetAllCategory()
			if id < 1 {
				fmt.Println("Category ID not found.")
				return
			}
			found := false

			for _, rows := range categories {
				if id == rows.ID {
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Category ID not found")
				return
			}

			oldCategory, err := handler.GetCategoryByID(id)
			if err != nil {
				fmt.Println("Error fetching category:", err)
				return
			}

			if name == "" {
				name = oldCategory.Name
			}
			if description == "" {
				description = oldCategory.Description
			}

			updatedCategory := &model.Category{
				Name:        name,
				Description: description,
			}

			_, err = handler.UpdateCategory(id, updatedCategory)
			if err != nil {
				fmt.Println("Error updating category:", err)
				return
			}

			fmt.Println("\nCATEGORY UPDATED SUCCESSFULLY")

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Category Name", "Description"})

			table.Append([]string{
				fmt.Sprintf("%d", id),
				updatedCategory.Name,
				updatedCategory.Description,
			})

			table.Render()
		},
	}
	//add flags
	UpdateCategory.Flags().IntVarP(&id, "id", "i", 0, "category id")
	UpdateCategory.Flags().StringVarP(&name, "name", "n", "", "category name")
	UpdateCategory.Flags().StringVarP(&description, "desc", "d", "", "categoryy description")
	UpdateCategory.MarkFlagRequired("id")

	return UpdateCategory

}

// function to delete category by ID
func DeleteCategoryCmd(handler handler.HandlerCategory) *cobra.Command {
	var id int
	DeleteCategory := &cobra.Command{
		Use:   "delete-category",
		Short: "to delete category by ID",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			categories, _ := handler.GetAllCategory()
			if id < 1 {
				fmt.Println("Category ID not found.")
				return
			}
			found := false

			for _, rows := range categories {
				if id == rows.ID {
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Category ID not found")
				return
			}
			err := handler.DeleteCategory(ctx, id)
			if err != nil {
				fmt.Println("Error deleting category:", err)
				return
			}

			fmt.Printf("\nCATEGORY WITH ID %d DELETED SUCCESSFULLY\n", id)

		},
	}
	DeleteCategory.Flags().IntVarP(&id, "id", "i", 0, "category id")
	DeleteCategory.MarkFlagRequired("id")

	return DeleteCategory
}
