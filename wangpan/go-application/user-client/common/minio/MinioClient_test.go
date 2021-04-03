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
	err := client.NewMinIoClient()
	assert.Nil(t, err)
}

func TestClient_UploadFile(t *testing.T) {
	client.NewMinIoClient()
	file, _ := os.Open("./张蕊芝是个猪猪.jpg")
	defer file.Close()
	stat, _ := file.Stat()
	err := client.UploadFile("show1", "张猪猪.jpg", "jpg", file, stat.Size())
	assert.Nil(t, err)
}

func TestClient_GetUploadOptions(t *testing.T) {
	assert.Equal(t, client.GetUploadOptions("png"), "image/png")
	assert.Equal(t, client.GetUploadOptions("mp4"), "application/octet-stream")
}

func TestClient_CreateBucket(t *testing.T) {
	err := client.CreateBucket("show1")
	assert.Nil(t, err)
}
