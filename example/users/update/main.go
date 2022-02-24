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

	// Create an user passing the user data following the CreateUser struct
	user, err := client.Users().Create(context.Background(), sdmscim.CreateUserBody{
		UserName:   "user@email.com",
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})

	if err != nil {
		log.Fatal("Error creating a user: ", err)
	}

	fmt.Print("\nUser:\n\n")
	if user != nil {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}

	fmt.Println("Updating user id:", user.ID, "...")

	// Update an user passing the user ID and the new user active state
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	user, err = client.Users().Update(ctx, user.ID, false)
	if err != nil {
		log.Fatal("Error updating the user: ", err)
	}
	fmt.Print("\nUser:\n\n")
	if user != nil {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}
}
