package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	ginEngine := gin.Default()
	firebaseCredentials := os.Getenv("FIREBASE_CREDENTIALS")
	opt := option.WithCredentialsJSON([]byte(firebaseCredentials)) // Replace with the path to your service account key JSON file
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	var usersRecords []*auth.ExportedUserRecord = make([]*auth.ExportedUserRecord, 0)

	// Get an Auth client from the Firebase Admin SDK
	authClient, err := app.Auth(context.Background())
	if err != nil {
		err.Error()
		log.Fatalf("Failed to create Firebase Auth client: %v", err)
	}

	// List all users
	users := authClient.Users(context.Background(), "")
	if err != nil {
		log.Fatalf("Failed to retrieve users: %v", err)
	}

	sum := 0
	for sum < users.PageInfo().Remaining() {
		userRecord, err := users.Next()
		if err != nil {
			log.Fatalf("Failed to get the user: %v", err)
		}
		usersRecords = append(usersRecords, userRecord)
	}
	ginEngine.POST("/token", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, usersRecords)
	})
	ginEngine.Run()
}
