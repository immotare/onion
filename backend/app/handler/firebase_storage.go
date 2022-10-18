package handler

import (
	"encoding/json"
	"log"
	"net/http"

	cstorage "cloud.google.com/go/storage"
	fbstorage "firebase.google.com/go/storage"
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

	log.Printf("uid:%s", r.Form.Get("uid"))
	result := map[string]string{
		"itemNames": r.Form.Get("uid"),
	}

	res, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (firebaseAppHandler *FirebaseBucketHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
}

func (firebaseAppHandler *FirebaseBucketHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
}

func (firebaseAppHandler *FirebaseBucketHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
}
