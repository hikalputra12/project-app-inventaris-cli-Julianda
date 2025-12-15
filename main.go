package main

import (
	"context"
	"inventory-system/cmd"
	"inventory-system/database"
	"inventory-system/handler"
	"inventory-system/repository"
	"inventory-system/service"
	"log"
)

// init function
func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	//for category
	categoryrepository := repository.NewRepositoryCategory(db)
	categoryservice := service.NewServiceCategory(&categoryrepository)
	categoryHandler := handler.NewHandlerCategory(&categoryservice)
	//for inventory
	inventoryRepository := repository.NewRepositoryInventory(db)
	inventoryService := service.NewServiceInventory(&inventoryRepository)
	inventoryHandler := handler.NewHandlerInventory(&inventoryService)
	allHandlers := handler.AllHandlers{ //ambil nanti di struct handler
		Category:  categoryHandler,
		Inventory: inventoryHandler,
	}

	// 5. Panggil fungsi pusat dengan container handler
	cmd.Init(allHandlers)
	cmd.Execute()
}
