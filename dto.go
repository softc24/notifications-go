package notifications

// NotificationResponse represents the response from sending a notification.
type NotificationResponse struct {
	// UUID is the ID
	UUID string `json:"uuid"`
	// ID is the external ID
	ID *string `json:"id,omitempty"`
	// Recipients are the recipients
	Recipients []Recipient `json:"recipients"`
	// Success is the count of successful deliveries
	Success int `json:"success"`
	// Error is the count of errors
	Error int `json:"error"`
}

// ProvidersResponse is a slice of ProviderInfo.
type ProvidersResponse []ProviderInfo

// ChannelsResponse is a slice of ChannelInfo.
type ChannelsResponse []ChannelInfo
