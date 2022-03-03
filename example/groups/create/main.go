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

	// Create an group passing the user data following the CreateGroupBody struct
	group, err := client.Groups().Create(context.Background(), scimsdk.CreateGroupBody{
		DisplayName: "xxx",
		Members:     []scimsdk.GroupMember{},
	})
	if err != nil {
		log.Fatal("Error creating a group: ", err)
	}
	fmt.Print("\nCreated Group:\n\n")
	if group != nil {
		fmt.Println("ID:", group.ID)
		fmt.Println("Display Name:", group.DisplayName)
		if group.Members != nil && len(group.Members) > 0 {
			fmt.Println("Members:")
			for _, member := range group.Members {
				fmt.Println("\t- Display:", member.Display)
				fmt.Println("\t- Value:", member.Value)
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
