// Package notifications provides a client for interacting with the notifications API.
// It allows sending notifications through various channels and providers.
package notifications

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/capcom6/go-restkit"
)

// Client represents a notifications API client.
type Client struct {
	*restkit.Client

	headers http.Header
}

// New creates a new notifications client with the provided configuration.
func New(config Config) (*Client, error) {
	rest, err := restkit.NewClient(restkit.Config{
		Client:  config.Client,
		BaseURL: config.BaseURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &Client{
		Client: rest,

		headers: http.Header{
			"Authorization": []string{
				"Basic " + base64.StdEncoding.EncodeToString([]byte(config.Username+":"+config.Password)),
			},
			"User-Agent": []string{"notifications-go/dev"},
			"Accept":     []string{"application/json"},
		},
	}, nil
}

// GetProviders returns a list of available providers.
func (c *Client) GetProviders(ctx context.Context) (ProvidersResponse, error) {
	var response ProvidersResponse
	err := c.Client.Do(ctx, http.MethodGet, "v1/provider", c.headers, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get providers: %w", err)
	}
	return response, nil
}

// GetChannels returns a list of available channels.
func (c *Client) GetChannels(ctx context.Context) (ChannelsResponse, error) {
	var response ChannelsResponse
	err := c.Client.Do(ctx, http.MethodGet, "v1/channel", c.headers, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get channels: %w", err)
	}
	return response, nil
}

// SendNotification sends a notification through the specified provider.
func (c *Client) SendNotification(
	ctx context.Context,
	provider string,
	request NotificationRequest,
) (*NotificationResponse, error) {
	var response NotificationResponse
	err := c.Client.Do(
		ctx,
		http.MethodPost,
		"v1/notification/"+url.PathEscape(provider),
		c.headers,
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %w", err)
	}
	return &response, nil
}
