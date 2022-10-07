package repository

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
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
		t.Errorf("[Failed]:%v", err)
	}

	fileNames, err := IndexItems(firebaseStorageBucket)

	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	fmt.Printf("result of index items:%v\n", fileNames)
}

func TestGetItem(t *testing.T) {
	storageBucketName := os.Getenv("STORAGE_BUCKET")
	firebaseCredentialPath := os.Getenv("FIREBASE_CREDENTIAL")

	firebaseStorageBucket, err := CreateStorageBucketClient(storageBucketName, firebaseCredentialPath)

	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	item, err := GetItem(firebaseStorageBucket, "test/test.txt")
	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(item)

	fmt.Printf("text in test/test.txt:%s\n", buf.String())
}

func TestCreateItem(t *testing.T) {
	storageBucketName := os.Getenv("STORAGE_BUCKET")
	firebaseCredentialPath := os.Getenv("FIREBASE_CREDENTIAL")

	firebaseStorageBucket, err := CreateStorageBucketClient(storageBucketName, firebaseCredentialPath)

	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	itemPath := "test/test_create.txt"
	contentType := "text/plain"
	item := strings.NewReader("created\n")

	err = CreateItem(firebaseStorageBucket, itemPath, contentType, item)
	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	fmt.Println("test/test_create.txt sucessfully created")

	itemReader, err := GetItem(firebaseStorageBucket, itemPath)

	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(itemReader)

	fmt.Printf("text in test/test_create.txt:%s\n", buf.String())
}

func TestDeleteItem(t *testing.T) {
	storageBucketName := os.Getenv("STORAGE_BUCKET")
	firebaseCredentialPath := os.Getenv("FIREBASE_CREDENTIAL")

	firebaseStorageBucket, err := CreateStorageBucketClient(storageBucketName, firebaseCredentialPath)

	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	itemPath := "test/test_create.txt"

	err = DeleteItem(firebaseStorageBucket, itemPath)
	if err != nil {
		t.Errorf("[Failed]:%v", err)
	}

	fmt.Println("test/test_create.txt sucessfully deleted")
}
