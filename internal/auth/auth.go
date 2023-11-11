package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey gets the API key from the request headers
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no Authorization header provided")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("missing first part of authorization header")
	}

	return vals[1], nil
}
