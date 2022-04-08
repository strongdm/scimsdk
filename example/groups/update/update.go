package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	fmt.Println("Creating an user...")

	// Create an user passing the user data following the CreateUser struct
	user, err := client.Users().Create(context.Background(), models.CreateUser{
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

	fmt.Println("Updating group adding members...")

	// Update the group adding members with the specified group id
	ok, err := client.Groups().UpdateAddMembers(context.Background(), group.ID, []models.GroupMember{
		{
			ID:    user.ID,
			Email: user.UserName,
		},
	})
	if err != nil {
		log.Fatal("Error updating the group: ", err)
	}
	if ok {
		fmt.Println("Group updated successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Finding group id:", group.ID)

	// Get the group with the specified group id
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

	fmt.Println("Updating group removing member...")

	// Update the group adding members with the specified group id
	ok, err = client.Groups().UpdateRemoveMemberByID(context.Background(), group.ID, user.ID)
	if err != nil {
		log.Fatal("Error updating the group: ", err)
	}
	if ok {
		fmt.Println("Group updated successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Finding group id:", group.ID)

	// Get the group with the specified group id
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

	fmt.Println("Updating group replacing name...")

	// Update the group adding members with the specified group id
	ok, err = client.Groups().UpdateReplaceName(context.Background(), group.ID, models.UpdateGroupReplaceName{DisplayName: "New Group Name"})
	if err != nil {
		log.Fatal("Error updating the group: ", err)
	}
	if ok {
		fmt.Println("Group updated successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Finding group id:", group.ID)

	// Get the group with the specified group id
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

	fmt.Println("Creating another user...")

	// Create an user passing the user data following the CreateUser struct
	user, err = client.Users().Create(context.Background(), models.CreateUser{
		UserName:   "user+01@email.com",
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

	fmt.Println("Updating group replacing members...")

	// Update the group adding members with the specified group id
	ok, err = client.Groups().UpdateReplaceMembers(context.Background(), group.ID, []models.GroupMember{
		{
			ID:    user.ID,
			Email: user.UserName,
		},
	})
	if err != nil {
		log.Fatal("Error updating the group: ", err)
	}
	if ok {
		fmt.Println("Group updated successfully")
	}

	fmt.Printf("\n----------------\n\n")

	fmt.Println("Finding group id:", group.ID)

	// Get the group with the specified group id
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
