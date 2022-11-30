package factory

import "fmt"

// Notification init interface (can also be an abs class)
type Notification interface {
	publishNotification() string
}

const (
	SMSType = iota
	EmailType
)

// SmsNotification factory object 1
type SmsNotification struct {
	text string
}

func (s *SmsNotification) publishNotification() string {
	return fmt.Sprintf("SMS notification: %s", s.text)
}

// EmailNotification factory object 2
type EmailNotification struct {
	text string
}

func (e *EmailNotification) publishNotification() string {
	return fmt.Sprintf("Email notification: %s", e.text)
}

// NotificationFactory factory
func NotificationFactory(ty int, txt string) Notification {
	switch ty {
	case SMSType:
		return &SmsNotification{text: txt}
	case EmailType:
		return &EmailNotification{text: txt}
	default:
		panic("this notification does not exist")
	}
}

// ExampleService example implementation
func ExampleService() {
	n := NotificationFactory(EmailType, "email notification")
	fmt.Println(n.publishNotification())
}
