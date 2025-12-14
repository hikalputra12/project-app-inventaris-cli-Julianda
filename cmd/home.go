package cmd

import (
	"fmt"
	"inventory-system/handler"
	"os"
)

// HomePage displays the main menu and handles user input.
func HomePage(handler handler.AllHandlers) {
	for {
		var choise int
		fmt.Println("=== Inventory System ===")
		fmt.Println("1. Go to Category Menu")
		fmt.Println("2. Inventory Items")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choise)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}
		ClearScreen()
		switch choise {
		case 1:
			for {
				fmt.Println("=== Category Items ===")
				fmt.Println("\n1. View Categories")
				fmt.Println("2. Add Category")
				fmt.Println("3. View Category by ID")
				fmt.Println("4. Update Category by ID")
				fmt.Println("5. Delete Category")
				fmt.Println("6. Go to main menu")
				fmt.Println("7. Exit")

				var choise int
				fmt.Print("enter your input: ")
				_, err := fmt.Scanln(&choise)
				if err != nil {
					fmt.Println("Invalid input. Please enter a number.")
					return
				}
				ClearScreen()
				switch choise {
				case 1:
					GetAllCategory(handler.Category)
				case 2:
					CreateCategory(handler.Category)

				case 3:
					GetCategoryByID(handler.Category)

				case 4:
					UpdateCategory(handler.Category)
				case 5:
					DeleteCategory(handler.Category)
				case 6:
					HomePage(handler)
				case 7:
					fmt.Println("Exiting...")
					os.Exit(0)
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 2:
			fmt.Println("=== Inventory Items ===")
			for {
				fmt.Println("\n1. Get Items More Than 100 Days")
				fmt.Println("2. Get Total Investment Value")
				fmt.Println("3. Get Investment And Depreciation By ID")
				fmt.Println("4. Find Inventory By Name")
				fmt.Println("5. Go to main menu")
				fmt.Println("6. Exit")

				var choise int
				fmt.Print("enter your input: ")
				_, err := fmt.Scanln(&choise)
				if err != nil {
					fmt.Println("Invalid input. Please enter a number.")
					return
				}
				ClearScreen()
				switch choise {
				case 1:
					ItemsMoreThan100Days(handler.Inventory)
				case 2:
					TotalInvesmentValue(handler.Inventory)

				case 3:
					InvesmentAndDepreciationValueByID(handler.Inventory)

				case 4:
					FindInventoryByName(handler.Inventory)

				case 5:
					HomePage(handler)
				case 6:
					fmt.Println("Exiting...")
					os.Exit(0)
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")

		}
	}

}
