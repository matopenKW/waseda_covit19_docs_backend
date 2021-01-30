package repository

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func OpenAuth() (*auth.Client, error) {
	ctx := context.Background()

	key := os.Getenv("FIREBASE_ADMIN_KEY")
	log.Println(key)
	// opt := option.WithCredentialsJSON([]byte(key))
	opt := option.WithCredentialsFile("firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Firestore Auth Succeeded!")
	return client, nil
}

func OpenFirestore() (*firestore.Client, error) {
	key := os.Getenv("FIREBASE_ADMIN_KEY")
	log.Println(key)
	//	opt := option.WithCredentialsJSON([]byte(key))
	opt := option.WithCredentialsFile("firebase-adminsdk.json")

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetUserRecord(client *auth.Client, uid string) (*auth.UserRecord, error) {
	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return user, err
}