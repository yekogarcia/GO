package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (sms SMSNotification) SendNotification() {
	fmt.Println("Sending SMS notification")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (sms SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Mobile"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Email notification")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	// switch notificationType {
	// case "SMS":
	// 	return SMSNotification{}
	// case "Email":
	// 	return EmailNotification{}
	// default:
	// 	return nil
	// }

	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	} else if notificationType == "Email" {
		return &EmailNotification{}, nil
	} else {
		return nil, fmt.Errorf("invalid notification type")
	}
}

func sendNotification(notificationFactory INotificationFactory) {
	notificationFactory.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	// notificationType := "SMS"
	// notificationFactory := getNotificationFactory(notificationType)
	// if notificationFactory != nil {
	// 	notificationFactory.SendNotification()
	// 	sender := notificationFactory.GetSender()
	// 	fmt.Printf("Sender Method: %s\n", sender.GetSenderMethod())
	// 	fmt.Printf("Sender Channel: %s\n", sender.GetSenderChannel())
	// } else {
	// 	fmt.Println("Invalid notification type")
	// }

	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	sendNotification(smsFactory)
	sendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)
}
