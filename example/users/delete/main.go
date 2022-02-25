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

	fmt.Print("\nUser:\n\n")
	if user != nil {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}

	if err != nil {
		log.Fatal("Error creating an user: ", err.Error())
	}

	// Initialize a context (you can use one with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Deleting user id:", user.ID)

	// Delete the user with the specified id
	ok, err := client.Users().Delete(ctx, user.ID)
	if err != nil {
		log.Fatal("Error deleting the user: ", err.Error())
	}
	if ok {
		fmt.Println("User", user.ID, "deleted successfully")
	}
}
