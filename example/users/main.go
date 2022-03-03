package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"scimsdk/scimsdk"
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
	user, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   "user@email.com",
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})
	if err != nil {
		log.Fatal("Error creating a user: ", err)
	}
	fmt.Print("\nCreated User:\n\n")
	if user != nil {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}

	fmt.Println("Finding user id:", user.ID)

	// Get the user with the specified user id
	user, err = client.Users().Find(context.Background(), user.ID)
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

	fmt.Println("Updating user id:", user.ID)

	ok, err := client.Users().Update(context.Background(), user.ID, scimsdk.UpdateUser{
		Active: true,
	})
	if err != nil {
		log.Fatal("Error updating the user: ", err)
	}
	if ok {
		fmt.Println("User updated successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Finding user id:", user.ID)

	// Get the user with the specified user id
	user, err = client.Users().Find(context.Background(), user.ID)
	if err != nil {
		log.Fatal("Error finding user: ", err.Error())
	}
	fmt.Print("\nUser Found after update:\n\n")
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
	user, err = client.Users().Replace(context.Background(), user.ID, scimsdk.ReplaceUser{
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

	fmt.Println("Deleting user id:", user.ID)

	// Delete the user with the specified id
	ok, err = client.Users().Delete(context.Background(), user.ID)
	if err != nil {
		log.Fatal("Error deleting the user: ", err.Error())
	}
	if ok {
		fmt.Println("User", user.ID, "deleted successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Listing users...")

	userIterator := client.Users().List(context.Background(), nil)
	fmt.Print("\nUsers:\n\n")
	for userIterator.Next() {
		user := userIterator.Value()
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n--------\n\n")
	}
	if userIterator.Err() != "" {
		log.Fatal("Iterator error:", userIterator.Err())
	}
}
