package repository

import (
	"context"

	storage "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func IndexItems(bucket *storage.BucketHandle) ([]string, error) {
	query := &storage.Query{
		Prefix: "",
	}
	itr := bucket.Objects(context.Background(), query)

	fileNames := []string{}
	for {
		attrs, err := itr.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return []string{}, err
		}

		fileNames = append(fileNames, attrs.Name)
	}

	return fileNames, nil
}
