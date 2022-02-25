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

	// Create a group passing the group data following the CreateGroupBody struct
	group, err := client.Groups().Create(context.Background(), sdmscim.CreateGroupBody{
		DisplayName: "xxx",
		Members:     []*sdmscim.GroupMember{},
	})
	if err != nil {
		log.Fatal("Error creating a group: ", err.Error())
	}

	fmt.Println("Finding group id:", group.ID)

	// Initialize a context (you can use one with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	// Update the group adding members with the specified group id
	ok, err := client.Groups().UpdateAddMembers(ctx, group.ID, []sdmscim.UpdateGroupMemberBody{
		{
			Value:   "a-0001",
			Display: "myUser@example.test",
		},
	})
	if err != nil {
		log.Fatal("Error updating the user: ", err)
	}
	if ok {
		fmt.Println("User updated successfullyy")
	}
}
