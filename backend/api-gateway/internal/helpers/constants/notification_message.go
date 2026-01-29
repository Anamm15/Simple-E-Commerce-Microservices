package constants

const (
	// Notification Success Messages
	MsgNotificationSent   = "Notification sent successfully"
	MsgNotificationQueued = "Notification has been queued for delivery"

	// Notification Error Messages
	MsgNotificationFailed  = "Failed to send notification"
	MsgInvalidRecipient    = "Recipient contact information is invalid"
	MsgTemplateNotFound    = "Notification template not found"
	MsgProviderUnreachable = "Third-party notification provider is down"
)
