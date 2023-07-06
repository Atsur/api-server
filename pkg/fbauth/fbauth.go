package fbauth

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	secretsmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/atsur/api-server/pkg/storage"
)

var (
	firebaseConfigFile = os.Getenv("FIREBASE_CONFIG_FILE")
)

func InitAuth() (*auth.Client, error) {

	file, err := storage.ReadFile(firebaseConfigFile);
	if err != nil {
		return nil, errors.Wrap(err, "error reading firebase config file (init auth)")
	}

	ctx := context.Background()
	creds, err := google.CredentialsFromJSON(ctx, file, secretsmanager.DefaultAuthScopes()...)
	if err != nil {
		// TODO: handle error.
		return nil, errors.Wrap(err, "error getting cloud credentials from json (init auth))")
	}
	// client, err := secretsmanager.NewClient(ctx, option.WithCredentials(creds))
	// if err != nil {
	// 	// TODO: handle error.
	// 	return nil, errors.Wrap(err, "error initializing firebase auth (create firebase app)")
	// }
	// _ = client // Use the client.



	opt := option.WithCredentials(creds)

	// opt := option.WithCredentialsFile(firebaseConfigFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing firebase auth (create firebase app)")
	}

	client, errAuth := app.Auth(ctx)
	if errAuth != nil {
		return nil, errors.Wrap(errAuth, "error initializing firebase auth (creating client)")
	}

	return client, nil
}
