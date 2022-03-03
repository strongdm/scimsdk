package service

import "strings"

func extractAuthorizationToken(authHeaderValue string) string {
	token := strings.Split(authHeaderValue, "Bearer")[1]
	return strings.TrimSpace(token)
}
