package utils

import (
	"net/url"
	"strings"
)

// IsValidURL checks if a string is a valid HTTP/HTTPS URL
func IsValidURL(str string) bool {
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	// Ensure the scheme is web-based
	scheme := strings.ToLower(u.Scheme)
	if scheme != "http" && scheme != "https" {
		return false
	}

	// Ensure there is actually a host/domain name
	if u.Host == "" {
		return false
	}

	return true
}
