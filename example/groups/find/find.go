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

	// Create a group passing the group data following the CreateGroupBody struct
	group, err := client.Groups().Create(context.Background(), models.CreateGroupBody{
		DisplayName: "xxx",
		Members:     []models.GroupMember{},
	})
	if err != nil {
		log.Fatal("Error creating a group: ", err.Error())
	}
	fmt.Print("\nCreated Group:\n\n")
	if group != nil {
		fmt.Println("ID:", group.ID)
		fmt.Println("Display Name:", group.DisplayName)
		if group.Members != nil && len(group.Members) > 0 {
			fmt.Println("Members:")
			for _, member := range group.Members {
				fmt.Println("\t- Display:", member.Email)
				fmt.Println("\t- Value:", member.ID)
			}
		} else {
			fmt.Println("Members: no members found")
		}
		fmt.Println("Meta:")
		fmt.Println("\t- Resource Type:", group.Meta.ResourceType)
		fmt.Println("\t- Location:", group.Meta.Location)
		fmt.Printf("\n----------------\n\n")
	}

	fmt.Println("Finding group id:", group.ID)

	// Initialize a context (you can use one with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get the user with the specified user id
	group, err = client.Groups().Find(ctx, group.ID)
	if err != nil {
		log.Fatal("Error finding group: ", err.Error())
	}
	fmt.Print("\nGroup Found:\n\n")
	if group != nil {
		fmt.Println("ID:", group.ID)
		fmt.Println("Display Name:", group.DisplayName)
		if group.Members != nil && len(group.Members) > 0 {
			fmt.Println("Members:")
			for _, member := range group.Members {
				fmt.Println("\t- Display:", member.Email)
				fmt.Println("\t- Value:", member.ID)
			}
		} else {
			fmt.Println("Members: no members found")
		}
		fmt.Println("Meta:")
		fmt.Println("\t- Resource Type:", group.Meta.ResourceType)
		fmt.Println("\t- Location:", group.Meta.Location)
		fmt.Printf("\n----------------\n\n")
	}
}
