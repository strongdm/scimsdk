package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/strongdm/scimsdk"
	"github.com/strongdm/scimsdk/models"
)

func main() {
	// Get the Admin Token from SDM
	token := os.Getenv("SDM_SCIM_TOKEN")
	if token == "" {
		log.Fatal("You must define SDM_SCIM_TOKEN env variable.")
	}
	// Initialize the SDM SCIM Client passing the admin token
	client := scimsdk.NewClient(token, nil)

	// Create an user passing the user data following the CreateUser struct
	user, err := client.Users().Create(context.Background(), models.CreateUser{
		UserName:   "user@email.com",
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})
	if err != nil {
		log.Fatal("Error creating an user: ", err.Error())
	}

	fmt.Println("Finding user id:", user.ID)

	// Initialize a context (you can use one with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get the user with the specified user id
	user, err = client.Users().Find(ctx, user.ID)
	if err != nil {
		log.Fatal("Error finding user: ", err.Error())
	}
	fmt.Print("\nUser Found:\n\n")
	if user != nil {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}
}
