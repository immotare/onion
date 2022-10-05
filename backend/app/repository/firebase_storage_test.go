package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/immotare/onion/firebaseutils"
)

func CreateStorageBucketClient(storageBucketName string, firebaseCredentialPath string) (*storage.BucketHandle, error) {
	fmt.Println("BacketName:", storageBucketName)
	firebaseApp, err := firebaseutils.NewApp(storageBucketName, firebaseCredentialPath)
	if err != nil {
		return nil, err
	}
	fmt.Println("firebase app client successfully created")

	firebaseStorageClient, err := firebaseApp.Storage(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Println("firebase storage client successfully created")

	firebaseStorageBucket, err := firebaseStorageClient.DefaultBucket()
	if err != nil {
		return nil, err
	}
	fmt.Println("firebase bucket client successfully created")

	return firebaseStorageBucket, nil
}

func TestIndexItems(t *testing.T) {
	storageBucketName := os.Getenv("STORAGE_BUCKET")
	firebaseCredentialPath := os.Getenv("FIREBASE_CREDENTIAL")

	firebaseStorageBucket, err := CreateStorageBucketClient(storageBucketName, firebaseCredentialPath)

	if err != nil {
		t.Errorf("Test Failed:%v", err)
	}

	fileNames, err := IndexItems(firebaseStorageBucket)

	if err != nil {
		t.Errorf("Test Failed:%v", err)
	}

	fmt.Printf("result of index items:%v", fileNames)
}
