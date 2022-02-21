package sdmscim

import (
	"strings"
)

type Client struct {
	Users      *UserModule
	Groups     *GroupModule
	adminToken string
}

func NewClient(adminToken string) *Client {
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{adminToken: trimmedToken}
	client.Users = &UserModule{client: client}
	client.Groups = &GroupModule{client: client}
	return client
}
