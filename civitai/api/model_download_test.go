package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCivitaiClient_downloadWithProgressBar(t *testing.T) {
	client := initTestClient()

	err := client.ModelVerDownloadByID(context.Background(), "352581", "647401", "./aaa.ckpt")
	assert.NoError(t, err)
}
