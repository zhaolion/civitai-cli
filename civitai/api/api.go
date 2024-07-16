package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/imroc/req/v3"
)

const (
	CivitaiAPIBaseV1 = "https://civitai.com/api/"
	UserAgent        = "civitai-go-client (https://github.com/zhaolion/civitai-cli)"
)

var (
	abortRetryStatusCodes = []int{http.StatusNotFound, http.StatusUnauthorized}
)

type CivitaiClient struct {
	*req.Client
	// To make authorized requests as a user you must use an API Key.
	// You can generate an API Key from your User Account Settings.
	// Your Account Settings URL - https://civitai.com/user/account
	//
	// Once you have an API Key you can authenticate with either an Authorization Header or Query String.
	// Creators can require that people be logged in to download their resources.
	// That is an option we provide but not something we require – it's entirely up to the resource owner.
	//
	// API Doc: https://developer.civitai.com/docs/api/public-rest#authorization
	token string
}

// NewClient with an API Key you can authenticate
func NewClient(token string, opts ...CivitaiClientOption) *CivitaiClient {
	cli := req.C().
		SetBaseURL(CivitaiAPIBaseV1).
		SetUserAgent(UserAgent).
		SetCommonHeader("Content-Type", "application/json").
		SetCommonBearerAuthToken(token).
		SetCommonRetryCount(3).
		SetCommonRetryBackoffInterval(time.Second, time.Second*30).
		SetCommonRetryInterval(func(resp *req.Response, attempt int) time.Duration {
			// Sleep seconds from "Retry-After" response header if it is present and correct.
			// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
			if resp.Response != nil {
				if ra := resp.Header.Get("Retry-After"); ra != "" {
					after, err := strconv.Atoi(ra)
					if err == nil {
						return time.Duration(after) * time.Second
					}
				}
			}
			return 2 * time.Second // Otherwise, sleep 2 seconds
		}).
		SetCommonRetryCondition(func(resp *req.Response, err error) bool {
			if resp.StatusCode >= http.StatusInternalServerError {
				return true
			}
			if resp.StatusCode <= 299 || containsInt(abortRetryStatusCodes, resp.StatusCode) {
				return false
			}
			return err != nil
		}).
		SetRedirectPolicy(
			// Only allow up to 5 redirects
			req.MaxRedirectPolicy(3),
		)
	client := &CivitaiClient{
		Client: cli,
		token:  token,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

/****************************************
*  CivitaiClientOption
*****************************************/

type CivitaiClientOption func(*CivitaiClient)
