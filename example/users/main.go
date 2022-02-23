package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sdmscim/sdmscim"
	"time"
)

func main() {
	// Get the Admin Token from SDM
	token := os.Getenv("SDM_ADMIN_TOKEN")
	if token == "" {
		log.Fatal("You must define SDM_ADMIN_TOKEN env variable.")
	}
	// Initialize the SDM SCIM Client passing the admin token
	client := sdmscim.NewClient(token, nil)

	// Initialize a context (you can use one with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get the users iterator struct to paginate through the data
	userIterator := client.Users.List(ctx, nil)
	fmt.Print("\nUsers:\n\n")
	for userIterator.Next() {
		user := userIterator.Value()
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}
	if userIterator.Err() != "" {
		log.Fatal("Iterator error:", userIterator.Err())
	}
}
