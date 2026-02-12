package notifications

// RecipientState represents the state of a notification recipient.
type RecipientState string

// ChannelID represents the ID of a notification channel.
type ChannelID string

const (
	// ParamUserID is the parameter name for user ID.
	ParamUserID = "userId"
	// ParamUserSecret is the parameter name for user secret.
	ParamUserSecret = "userSecret"
	// ParamSender is the parameter name for sender.
	ParamSender = "sender"

	// RecipientStatePending represents a pending recipient state.
	RecipientStatePending RecipientState = "pending"
	// RecipientStateSent represents a sent recipient state.
	RecipientStateSent RecipientState = "sent"
	// RecipientStateDelivered represents a delivered recipient state.
	RecipientStateDelivered RecipientState = "delivered"
	// RecipientStateError represents an error recipient state.
	RecipientStateError RecipientState = "error"

	// ChannelPhone represents the phone channel.
	ChannelPhone ChannelID = "phone"
	// ChannelEmail represents the email channel.
	ChannelEmail ChannelID = "email"
	// ChannelTelegram represents the telegram channel.
	ChannelTelegram ChannelID = "telegram"

	// FieldSubject is the field name for subject.
	FieldSubject = "subject"
	// FieldContentType is the field name for content type.
	FieldContentType = "contentType"
	// FieldInlineAttachments is the field name for inline attachments.
	FieldInlineAttachments = "inlineAttachments"

	// ContentTypeText represents text content type.
	ContentTypeText = "text"
	// ContentTypeHTML represents HTML content type.
	ContentTypeHTML = "html"
	// ContentTypeMarkdown represents markdown content type.
	ContentTypeMarkdown = "markdown"
)

// NotificationRequest represents a request to send a notification.
type NotificationRequest struct {
	// ID is the external ID
	ID *string `json:"id,omitempty"`
	// Recipients are the addresses of recipients
	Recipients []string `json:"recipients"`
	// Payload is the content of the notification
	Payload string `json:"payload"`
	// Options are notification parameters, when using multiple providers,
	// the provider name prefix with two underscores is used: `backend__`
	Options map[string]string `json:"options"`
	// Fields are additional fields specific to the sending channel
	Fields map[string]string `json:"fields"`
}

// Recipient represents a notification recipient.
type Recipient struct {
	// Address is the address
	Address string `json:"address"`
	// State is the state
	State RecipientState `json:"state" enums:"pending,sent,delivered,error"`
	// Error is the error text
	Error *string `json:"error,omitempty"`
}

// ProviderInfo represents information about a notification provider.
type ProviderInfo struct {
	// ID is the ID
	ID string `json:"id"`
	// Title is the title
	Title string `json:"title"`
	// Description is the description
	Description string `json:"description,omitempty"`
	// Help is the help information
	Help string `json:"help,omitempty"`
	// Channels are the channels
	Channels []ChannelID `json:"channels,omitempty"`
	// Params are the available parameters
	Params []ParamsItem `json:"params"`
}

// ParamsItem represents a parameter item.
type ParamsItem struct {
	// ID is the ID
	ID string `json:"id"`
	// Title is the title
	Title string `json:"title"`
	// Description is the description
	Description string `json:"description,omitempty"`
	// Regex is the regular expression for validation
	Regex string `json:"regex,omitempty"`
	// IsRequired indicates if it's required
	IsRequired bool `json:"isRequired"`
}

// ChannelInfo represents information about a notification channel.
type ChannelInfo struct {
	// ID is the ID
	ID ChannelID `json:"id"`
	// Title is the title
	Title string `json:"title"`
	// Fields are the additional fields
	Fields []ChannelField `json:"fields"`
}

// ChannelField represents a field in a notification channel.
type ChannelField struct {
	// ID is the ID
	ID string `json:"id"`
	// Title is the title
	Title string `json:"title"`
	// Description is the description
	Description string `json:"description,omitempty"`
	// Regex is the regular expression for validation
	Regex string `json:"regex,omitempty"`
	// IsRequired indicates if it's required
	IsRequired bool `json:"isRequired"`
}

// Attachment represents an attachment for a notification.
type Attachment struct {
	// ContentType is the content type
	ContentType string `json:"contentType"`
	// ContentBase64 is the content encoded in base64
	ContentBase64 string `json:"content"`
	// Name is the name
	Name string `json:"name"`
}

// Attachments is a slice of attachments.
type Attachments []Attachment
