package cmd

import (
	"fmt"
	"os"

	"inventory-system/handler"
	"inventory-system/utils"

	"github.com/olekukonko/tablewriter"
)

// function to display inventory items used for more than 100 days
func ItemsMoreThan100Days(handler handler.HandlerInventory) {
	inventories, err := handler.ItemsMoreThan100Days()
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
}

// function to display total investment value after depreciation
func TotalInvesmentValue(handler handler.HandlerInventory) {
	totalInvestment, err := handler.TotalInvesmentValue()
	if err != nil {
		fmt.Println("Error fetching inventories:", err)
		return
	}

	fmt.Println("\nTOTAL INVESTMENT VALUE AFTER 20% DEPRECIATION PER YEAR")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(
		"Total Investment Value (Rp)",
	)

	table.Append([]string{
		fmt.Sprintf("%.0f", totalInvestment[0].TotalInvestmentValue),
	})

	table.Render()
}

// function to display investment and depreciation value by inventory ID
func InvesmentAndDepreciationValueByID(handler handler.HandlerInventory) {
	var id int
	fmt.Print("Enter Inventory ID: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	data, err := handler.InvesmentAndDepreciationValueByID(id)
	if err != nil {
		fmt.Println("Error fetching inventory:", err)
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
}

// function to find inventory by name
func FindInventoryByName(handler handler.HandlerInventory) {
	name := utils.ReadLine("Enter Item Name: ")
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
}
