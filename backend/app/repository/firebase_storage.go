package repository

import (
	"context"
	"io"

	storage "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func IndexItems(bucket *storage.BucketHandle) ([]string, error) {
	query := &storage.Query{
		Prefix: "",
	}
	itr := bucket.Objects(context.Background(), query)

	itemNames := []string{}
	for {
		attrs, err := itr.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return []string{}, err
		}

		itemNames = append(itemNames, attrs.Name)
	}

	return itemNames, nil
}

func GetItem(bucket *storage.BucketHandle, itemPath string) (io.Reader, error) {
	obj := bucket.Object(itemPath)
	reader, err := obj.NewReader(context.Background())

	if err != nil {
		return nil, err
	}
	return reader, nil
}

func CreateItem(bucket *storage.BucketHandle, itemPath string, contentType string, item io.Reader) error {
	obj := bucket.Object(itemPath)
	writer := obj.NewWriter(context.Background())
	writer.ContentType = contentType

	if _, err := io.Copy(writer, item); err != nil {
		return err
	}

	return writer.Close()
}

func DeleteItem(bucket *storage.BucketHandle, itemPath string) error {
	obj := bucket.Object(itemPath)

	if err := obj.Delete(context.Background()); err != nil {
		return err
	}

	return nil
}
