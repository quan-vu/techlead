package main

import (
	"fmt"
	"myproject/pkg/api"
	"myproject/pkg/utils"
)

func main() {
	fmt.Println("Starting server on http://127.0.0.1:8080")
	api.HandleRequests()
	// db.Connect()
	utils.Log("Server started successfully!")
}
