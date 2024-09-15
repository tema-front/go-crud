package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example
// Auth: ApiKey {insert apikey here}
func GetApiKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")

	if value == "" {
		return "", errors.New("no auth info found")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		return "", errors.New("malformed auth header")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return values[1], nil
}