package download

import (
	"github.com/zhaolion/civitai-cli/civitai/api"
)

type Client struct {
	api *api.CivitaiClient
}

func NewClient(api *api.CivitaiClient) *Client {
	return &Client{
		api: api,
	}
}

func DefaultClient() *Client {
	return NewClient(api.NewClient(api.GetAPIToken()))
}
