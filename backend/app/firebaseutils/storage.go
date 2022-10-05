package firebaseutils

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewApp(storageBucket, storageCredentialPath string) (*firebase.App, error) {
	config := &firebase.Config{
		StorageBucket: storageBucket,
	}
	opt := option.WithCredentialsFile(storageCredentialPath)
	app, err := firebase.NewApp(context.Background(), config, opt)

	return app, err
}
