package api

import (
	"os"
	"testing"
)

func TestSetAPIToken(t *testing.T) {
	token := os.Getenv("CIVITAI_API_KEY")
	SetAPIToken(token)
}
