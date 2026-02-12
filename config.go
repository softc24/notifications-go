package notifications

import "net/http"

// Config represents the configuration for the notifications client.
type Config struct {
	// BaseURL is the base URL of the notifications API
	BaseURL string
	// Username is the username for authentication
	Username string
	// Password is the password for authentication
	Password string

	// Client is the HTTP client to use (optional, will use http.DefaultClient if nil)
	Client *http.Client
}
