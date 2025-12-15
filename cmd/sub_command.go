package cmd

import "inventory-system/handler"

// daftarkan subcommand ke root command
func Init(allHandler handler.AllHandlers) {
	rootCmd.AddCommand(GetAllCategoryCmd(allHandler.Category), CreateCategory(allHandler.Category), UpdateCategoryCmd(allHandler.Category),
		GetCategoryByIDCmd(allHandler.Category), DeleteCategoryCmd(allHandler.Category), ItemsMoreThan100DaysCmd(allHandler.Inventory),
		InvesmentAndDepreciationValueByIDCmd(allHandler.Inventory), FindInventoryByNameCmd(allHandler.Inventory))

}
