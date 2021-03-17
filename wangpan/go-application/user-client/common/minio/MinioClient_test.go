package minio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	client *Client
)

func init() {
	client = new(Client)
}

func TestClient_NewMinIoClient(t *testing.T) {
	client := new(Client)
	err := client.NewMinIoClient()
	assert.Nil(t, err)
}

func TestClient_UploadFile(t *testing.T) {
	client.NewMinIoClient()
	file, _ := os.Open("./user.jpg")
	defer file.Close()
	stat, _ := file.Stat()
	err := client.UploadFile("show", "UploadFile.jpg", "jpg", file, stat.Size())
	assert.Nil(t, err)
}

func TestClient_GetUploadOptions(t *testing.T) {
	client := new(Client)
	assert.Equal(t, client.GetUploadOptions("png"), "image/png")
	assert.Equal(t, client.GetUploadOptions("mp4"), "application/octet-stream")
}
