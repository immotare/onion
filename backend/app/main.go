package main

import (
	"log"
	"net/http"
	"os"

	"github.com/immotare/onion/firebaseutils"
	"github.com/immotare/onion/handler"
	"golang.org/x/net/context"
)

func OkResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("response sucess"))
	log.Println("response returned")
}

func main() {
	log.Println("server started")
	storageBucketName := os.Getenv("STORAGE_BUCKET")
	firebaseCredentialPath := os.Getenv("FIREBASE_CREDENTIAL")
	firebaseApp, err := firebaseutils.NewApp(storageBucketName, firebaseCredentialPath)
	if err != nil {
		log.Println("failed to initialize firebase app:", err)
	}

	firebaseStorageClient, err := firebaseApp.Storage(context.Background())
	if err != nil {
		log.Println("failed to create firebase client:", err)
	}

	firebaseStorageBucket, err := firebaseStorageClient.DefaultBucket()
	if err != nil {
		log.Println("failed to create firebase bucket:", err)
	}

	firebaseBucketHandler := handler.NewFirebaseBucketHandler(firebaseStorageClient, firebaseStorageBucket)

	http.HandleFunc("/index", firebaseBucketHandler.IndexItems)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("failed starting server")
	}
}
