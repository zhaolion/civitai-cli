package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCivitaiClient_downloadWithProgressBar(t *testing.T) {
	client := initTestClient()

	err := client.ModelDownloadByID("352581", "647401", "./aaa.ckpt")
	assert.NoError(t, err)
}
