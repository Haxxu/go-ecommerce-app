package notification

import "go-ecommerce-app/config"

type NotificationClient interface {
	SendSMS(phone string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSMS(phone string, message string) error {
	return nil
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}
