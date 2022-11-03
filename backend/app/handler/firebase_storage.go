package handler

import (
	"encoding/json"
	"log"
	"net/http"

	cstorage "cloud.google.com/go/storage"
	fbstorage "firebase.google.com/go/storage"
	"github.com/immotare/onion/repository"
)

type FirebaseBucketHandler struct {
	client *fbstorage.Client
	bucket *cstorage.BucketHandle
}

func NewFirebaseBucketHandler(client *fbstorage.Client, bucket *cstorage.BucketHandle) *FirebaseBucketHandler {
	return &FirebaseBucketHandler{
		client: client,
		bucket: bucket,
	}
}

func (firebaseAppHandler *FirebaseBucketHandler) IndexItems(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k, v := range r.Form {
		log.Printf("[] key: %s, value:%s", k, v)
	}

	uid := r.Form.Get("uid")
	log.Printf("uid:%s", uid)

	itemNames, err := repository.IndexItemsByUserId(firebaseAppHandler.bucket, uid)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := map[string][]string{
		"itemNames": itemNames,
	}

	log.Printf("index result:%v", itemNames)

	res, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (firebaseAppHandler *FirebaseBucketHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
}

func (firebaseAppHandler *FirebaseBucketHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
}

func (firebaseAppHandler *FirebaseBucketHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
}
