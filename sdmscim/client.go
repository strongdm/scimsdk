package sdmscim

import (
	"strings"
)

type Client struct {
	Users      *UserService
	Groups     *GroupService
	adminToken string
}

func New(adminToken string) *Client {
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{adminToken: trimmedToken}
	client.Users = &UserService{client: client}
	client.Groups = &GroupService{client: client}
	return client
}
