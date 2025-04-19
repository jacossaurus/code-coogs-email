package main

import (
	"context"
	"encoding/base64"

	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	credentials, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Failed to read credentials file: %s", err)
	}

	config, err := google.JWTConfigFromJSON(credentials, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	config.Subject = "main@codecoogs.com"

	service, err := gmail.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx)), option.WithScopes(gmail.MailGoogleComScope))
	if err != nil {
		log.Fatalf("Failed to create Gmail service: %s", err)
	}

	headers, err := os.ReadFile("headers.txt")
	if err != nil {
		log.Fatalf("Failed to get headers file: %s", err)
	}

	emailContent, err := os.ReadFile("email.html")
	if err != nil {
		log.Fatalf("Failed to get email content file: %s", err)
	}

	message := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte(string(headers) + "\r\n" + string(emailContent))),
	}

	_, err = service.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Failed to send message: %s", err)
	}
}
