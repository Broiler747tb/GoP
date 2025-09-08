package main

import (
	"GoP/api"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Load environment variables first
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}

	// Initialize API
	api.Api()

	// Define flags
	create := flag.Bool("create", false, "create a json")
	update := flag.Bool("update", false, "update a json")
	delete := flag.Bool("delete", false, "delete a json")
	get := flag.Bool("get", false, "get a json")
	list := flag.Bool("list", false, "get a list")

	file := flag.String("file", "", "file name")
	name := flag.String("name", "", "name")
	id := flag.String("id", "", "id")

	flag.Parse()

	// Check if at least one operation is specified
	operationsCount := 0
	if *create {
		operationsCount++
	}
	if *update {
		operationsCount++
	}
	if *delete {
		operationsCount++
	}
	if *get {
		operationsCount++
	}
	if *list {
		operationsCount++
	}

	if operationsCount == 0 {
		fmt.Println("Error: Please specify an operation (create, update, delete, get, or list)")
		flag.Usage()
		os.Exit(1)
	}

	if operationsCount > 1 {
		fmt.Println("Error: Please specify only one operation at a time")
		os.Exit(1)
	}

	// Execute operations
	if *create {
		api.Create(file, name)
	}
	if *update {
		api.Update(file, id)
	}
	if *delete {
		api.Delete(id)
	}
	if *get {
		api.Get(id)
	}
	if *list {
		api.List()
	}
}
