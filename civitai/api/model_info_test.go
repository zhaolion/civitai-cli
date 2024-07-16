package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initTestClient() *CivitaiClient {
	token := os.Getenv("CIVITAI_API_KEY")
	client := NewClient(token)
	client.DevMode()
	return client
}

func TestCivitaiClient_ModelInfoByID(t *testing.T) {
	client := initTestClient()

	model, err := client.ModelInfoByID("352581")
	assert.NoError(t, err)
	assert.NotNil(t, model)

	fmt.Println(model.JSON())

	for _, ver := range model.ModelVersions {
		fmt.Println(ver.ID)
		fmt.Println("---")
	}
}

func TestCivitaiClient_ModelVersion(t *testing.T) {
	client := initTestClient()

	model, err := client.ModelVersionByID("647401")
	assert.NoError(t, err)
	assert.NotNil(t, model)

	fmt.Println(model.JSON())
}
