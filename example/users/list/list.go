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

	// Initialize a context (you can use one with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get the users iterator struct to paginate through the data. You can
	// add pagination options to specify page size, offset and filter
	userIterator := client.Users().List(ctx, &models.PaginationOptions{
		PageSize: 5,
		Offset:   1,
	})
	fmt.Print("\nUser List:\n\n")
	for userIterator.Next() {
		user := userIterator.Value()
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name.Formatted)
		fmt.Println("Display Name:", user.DisplayName)
		fmt.Println("UserName:", user.UserName)
		fmt.Println("Active:", user.Active)
		fmt.Printf("\n----------------\n\n")
	}
	if userIterator.Err() != nil {
		log.Fatal("Iterator error:", userIterator.Err())
	}
}
