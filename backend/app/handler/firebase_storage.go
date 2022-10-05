package handler

import (
	"net/http"

	storage "cloud.google.com/go/storage"
)

type FirebaseBucketHandler struct {
	bucket *storage.BucketHandle
}

func NewFirebaseBucketHandler(bucket *storage.BucketHandle) *FirebaseBucketHandler {
	return &FirebaseBucketHandler{
		bucket: bucket,
	}
}

func (firebaseAppHandler *FirebaseBucketHandler) IndexItems(w http.ResponseWriter, r *http.Request) {

}
