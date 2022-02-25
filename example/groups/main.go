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

	// Create an group passing the user data following the CreateGroupBody struct
	group, err := client.Groups().Create(context.Background(), sdmscim.CreateGroupBody{
		DisplayName: "xxx",
		Members:     []sdmscim.GroupMember{},
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

	fmt.Println("Finding group id:", group.ID)

	group, err = client.Groups().Find(context.Background(), group.ID)
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

	fmt.Println("Replacing group...")

	group, err = client.Groups().Replace(context.Background(), group.ID, sdmscim.ReplaceGroupBody{
		DisplayName: "Replaced Display Name",
		Members:     []sdmscim.GroupMember{},
	})

	if err != nil {
		log.Fatal("Error replacing group: ", err.Error())
	}

	fmt.Print("\nReplaced Group:\n\n")
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

	fmt.Println("Deleting group...")

	ok, err := client.Groups().Delete(context.Background(), group.ID)

	if err != nil {
		log.Fatal("Error deleting group: ", err.Error())
	}
	if ok {
		fmt.Println("Group deleted successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Listing groups...")

	iterator := client.Groups().List(context.Background(), nil)
	if iterator.Err() != "" {
		log.Fatal("Error finding group: ", iterator.Err())
	}
	fmt.Print("\nGroup List:\n\n")
	for iterator.Next() {
		group := iterator.Value()

		fmt.Println("ID:", group.ID)
		fmt.Println("Display Name:", group.DisplayName)
		if group.Members != nil && len(group.Members) > 0 {
			fmt.Println("Members:")
			for _, member := range group.Members {
				fmt.Println("\t- Display:", member.Display)
				fmt.Println("\t- Value:", member.Value)
				fmt.Println()
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
