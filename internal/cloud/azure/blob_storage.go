package azure

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// get azure blob container
func getContainer() (*storage.Container, error) {
	storageAccountName := os.Getenv("ACCOUNT_NAME")
	accessKey := os.Getenv("ACCESS_KEY")
	containerName := os.Getenv("AZURE_STORAGE_CONTAINER_NAME")

	// Create a storage client
	client, err := storage.NewBasicClient(storageAccountName, accessKey)
	if err != nil {
		return nil, err
	}

	// Create client for Azure Blob Storage
	blobService := client.GetBlobService()
	container := blobService.GetContainerReference(containerName)
	return container, nil
}

// upload file to azure blob storage
func Upload(f_name string, content io.Reader) error {
	container, err := getContainer()
	if err != nil {
		return err
	}

	blob := container.GetBlobReference(f_name)
	err = blob.CreateBlockBlobFromReader(content, nil)
	if err != nil {
		return err
	}
	return nil
}

// download file from azure blob storage
func Download(f_name string) ([]byte, error) {
	container, err := getContainer()
	if err != nil {
		return nil, err
	}

	blob := container.GetBlobReference(f_name)
	opt := &storage.GetBlobOptions{}
	reader, err := blob.Get(opt)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return b, nil
}
