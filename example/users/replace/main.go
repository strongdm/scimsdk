package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sdmscim/sdmscim"
)

func main() {
	// Get the Admin Token from SDM
	token := os.Getenv("SDM_SCIM_TOKEN")
	if token == "" {
		log.Fatal("You must define SDM_SCIM_TOKEN env variable.")
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

	fmt.Println("Replacing user id:", user.ID, "...")

	// Create an user passing the user data following the CreateUser struct
	user, err = client.Users().Replace(context.Background(), user.ID, sdmscim.ReplaceUserBody{
		UserName:   "user+01@email.com",
		GivenName:  "test replaced",
		FamilyName: "name replaced",
		Active:     true,
	})
	if err != nil {
		log.Fatal("Error replacing the user: ", err)
	}
	fmt.Print("\nReplaced User:\n\n")
	if user != nil {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}
}
