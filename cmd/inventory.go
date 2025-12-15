package cmd

import (
	"fmt"
	"os"

	"inventory-system/handler"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// // function to display inventory items used for more than 100 days
func ItemsMoreThan100DaysCmd(handler handler.HandlerInventory) *cobra.Command {
	return &cobra.Command{
		Use:   "items-morethan-100days",
		Short: " to display inventory items used for mode than 100 days and must be replaced",
		Run: func(cmd *cobra.Command, args []string) {
			inventories, err := handler.ItemsMoreThan100Days()
			ClearScreen()
			if err != nil {
				fmt.Println("Error fetching inventories:", err)
				return
			}

			fmt.Println("\nINVENTORY ITEMS USED FOR MORE THAN 100 DAYS (NEED REPLACEMENT)")

			table := tablewriter.NewWriter(os.Stdout)
			table.Header(
				"No",
				"Item Name",
				"Price (Rp)",
				"Total Usage Days",
			)

			for i, inventory := range inventories {
				table.Append([]string{
					fmt.Sprintf("%d", i+1),
					inventory.Name,
					fmt.Sprintf("%.0f", inventory.Price),
					fmt.Sprintf("%d", inventory.TotalUsageDays),
				})
			}

			table.Render()

		},
	}
}

// function to display investment and depreciation value by inventory ID
func InvesmentAndDepreciationValueByIDCmd(handler handler.HandlerInventory) *cobra.Command {
	var id int
	InvesmentAndDepreciationValueByID := &cobra.Command{
		Use:   "invesment-depreciation",
		Short: "to display investment and depreciation value by inventory ID",
		Run: func(cmd *cobra.Command, args []string) {
			ClearScreen()
			data, err := handler.InvesmentAndDepreciationValueByID(id)
			if err != nil {
				fmt.Println("Error fetching inventory:", err)
				return
			}
			if data == nil {
				fmt.Printf("Inventory with id %d is not found", id)
				return
			}

			fmt.Println("\nINVESTMENT & DEPRECIATION DETAILS")

			table := tablewriter.NewWriter(os.Stdout)
			table.Header(
				"Item Name",
				"Investment Value (Rp)",
				"Depreciation Value (Rp)",
			)

			table.Append([]string{
				data.Name,
				fmt.Sprintf("%.0f", data.InvestmentValue),
				fmt.Sprintf("%.0f", data.Depreciation),
			})

			table.Render()
		},
	}
	InvesmentAndDepreciationValueByID.Flags().IntVarP(&id, "id", "i", 0, "inventory id")
	InvesmentAndDepreciationValueByID.MarkFlagRequired("id")

	return InvesmentAndDepreciationValueByID
}

// function to find inventory by name
func FindInventoryByNameCmd(handler handler.HandlerInventory) *cobra.Command {
	var name string
	FindInventoryByName := &cobra.Command{
		Use:   "find-inventory-byname",
		Short: "to find inventory by name and display the detail about the inventory",
		Run: func(cmd *cobra.Command, args []string) {
			ClearScreen()
			if name == "" {
				fmt.Println("Item name cannot be empty")
				return
			}

			data, err := handler.FindInventoryByName(name)
			if err != nil {
				fmt.Println("Error fetching inventory:", err)
				return
			}
			if data == nil {
				fmt.Println("Inventory not found")
				return
			}

			fmt.Println("\nINVENTORY DETAILS")

			table := tablewriter.NewWriter(os.Stdout)
			table.Header(
				"Item Name",
				"Price (Rp)",
				"Purchase Date",
			)

			table.Append([]string{
				data.Name,
				fmt.Sprintf("%.0f", data.Price),
				data.PurchaseDate.Format("02/01/2006"),
			})

			table.Render()
		},
	}
	//add flags
	FindInventoryByName.Flags().StringVarP(&name, "name", "n", "", "inventory name")
	FindInventoryByName.MarkFlagRequired("name")

	return FindInventoryByName

}
