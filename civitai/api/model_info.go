package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	URLModelInfoByID        = "/v1/models/%s"
	URLModelVersionInfoByID = "/v1/model-versions/%s"
	URLModelVersionByID     = "/download/models/%s"
)

/****************************************
*  Model Common Structs
*****************************************/

type ModelStats struct {
	DownloadCount   int `json:"downloadCount,omitempty"` // version 公共字段 - begin
	RatingCount     int `json:"ratingCount,omitempty"`
	Rating          int `json:"rating,omitempty"`
	ThumbsUpCount   int `json:"thumbsUpCount,omitempty"`
	ThumbsDownCount int `json:"thumbsDownCount,omitempty"` // version 公共字段 - end

	FavoriteCount     int `json:"favoriteCount,omitempty"`
	CommentCount      int `json:"commentCount,omitempty"`
	TippedAmountCount int `json:"tippedAmountCount,omitempty"`
}

type ModelMetadata struct {
	Format string `json:"format,omitempty"`
	Size   string `json:"size,omitempty"`
	Fp     string `json:"fp,omitempty"`
}

type ModelHashes struct {
	AutoV1 string `json:"AutoV1,omitempty"`
	AutoV2 string `json:"AutoV2,omitempty"`
	Sha256 string `json:"SHA256,omitempty"`
	Crc32  string `json:"CRC32,omitempty"`
	Blake3 string `json:"BLAKE3,omitempty"`
}

/****************************************
*  ModelInfoByID
*****************************************/

// ModelInfoByID fetch model info by id.
// GET /api/v1/models/:modelId
// API Doc: https://developer.civitai.com/docs/api/public-rest#get-apiv1modelsmodelid
func (c *CivitaiClient) ModelInfoByID(modelID string) (*ModelInfo, error) {
	targetAPI := fmt.Sprintf(URLModelInfoByID, modelID)

	var (
		successResult ModelInfo
		errResult     ErrorResponse
	)
	resp, err := c.Client.R().SetErrorResult(&errResult).SetSuccessResult(&successResult).Get(targetAPI)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, NewNotFoundError("model", modelID)
		}

		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		errResult.StatusCode = resp.StatusCode
		return nil, &errResult
	}

	return &successResult, nil
}

type ModelInfo struct {
	ID                    int                 `json:"id,omitempty"`
	Name                  string              `json:"name,omitempty"`
	Description           string              `json:"description,omitempty"`
	AllowNoCredit         bool                `json:"allowNoCredit,omitempty"`
	AllowCommercialUse    []string            `json:"allowCommercialUse,omitempty"`
	AllowDerivatives      bool                `json:"allowDerivatives,omitempty"`
	AllowDifferentLicense bool                `json:"allowDifferentLicense,omitempty"`
	Type                  string              `json:"type,omitempty"`
	Minor                 bool                `json:"minor,omitempty"`
	Poi                   bool                `json:"poi,omitempty"`
	Nsfw                  bool                `json:"nsfw,omitempty"`
	NsfwLevel             int                 `json:"nsfwLevel,omitempty"`
	Cosmetic              interface{}         `json:"cosmetic,omitempty"`
	Stats                 ModelStats          `json:"stats,omitempty"`
	Creator               ModelCreator        `json:"creator,omitempty"`
	Tags                  []string            `json:"tags,omitempty"`
	ModelVersions         []ModelInfoVersions `json:"modelVersions,omitempty"`
}

func (m *ModelInfo) JSON() string {
	bs, _ := json.MarshalIndent(m, "", "\t")
	return string(bs)
}

type ModelCreator struct {
	Username string `json:"username,omitempty"`
	Image    string `json:"image,omitempty"`
}

type ModelInfoFiles struct {
	ID                int           `json:"id,omitempty"`
	SizeKB            float64       `json:"sizeKB,omitempty"`
	Name              string        `json:"name,omitempty"`
	Type              string        `json:"type,omitempty"`
	PickleScanResult  string        `json:"pickleScanResult,omitempty"`
	PickleScanMessage string        `json:"pickleScanMessage,omitempty"`
	VirusScanResult   string        `json:"virusScanResult,omitempty"`
	VirusScanMessage  interface{}   `json:"virusScanMessage,omitempty"`
	ScannedAt         time.Time     `json:"scannedAt,omitempty"`
	Metadata          ModelMetadata `json:"metadata,omitempty"`
	Hashes            ModelHashes   `json:"hashes,omitempty"`
	DownloadURL       string        `json:"downloadUrl,omitempty"`
	Primary           bool          `json:"primary,omitempty"`
}

type ModelInfoImages struct {
	URL       string `json:"url,omitempty"`
	NsfwLevel int    `json:"nsfwLevel,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	Hash      string `json:"hash,omitempty"`
	Type      string `json:"type,omitempty"`
}

type ModelInfoVersions struct {
	ID            int               `json:"id,omitempty"`
	Index         int               `json:"index,omitempty"`
	Name          string            `json:"name,omitempty"`
	BaseModel     string            `json:"baseModel,omitempty"`
	BaseModelType string            `json:"baseModelType,omitempty"`
	CreatedAt     time.Time         `json:"createdAt,omitempty"`
	PublishedAt   time.Time         `json:"publishedAt,omitempty"`
	Status        string            `json:"status,omitempty"`
	Availability  string            `json:"availability,omitempty"`
	NsfwLevel     int               `json:"nsfwLevel,omitempty"`
	Description   string            `json:"description,omitempty"`
	TrainedWords  []string          `json:"trainedWords,omitempty"`
	Stats         ModelStats        `json:"stats,omitempty"`
	Files         []ModelInfoFiles  `json:"files,omitempty"`
	Images        []ModelInfoImages `json:"images,omitempty"`
	DownloadURL   string            `json:"downloadUrl,omitempty"`
}

/****************************************
* ModelVersion
*****************************************/

// ModelVersionByID fetch model version files by ver-id.
// GET /api/v1/models/:modelId
// API Doc: https://developer.civitai.com/docs/api/public-rest#get-apiv1modelsmodelid
func (c *CivitaiClient) ModelVersionByID(verID string) (*ModelVersion, error) {
	targetAPI := fmt.Sprintf(URLModelVersionInfoByID, verID)

	var (
		successResult ModelVersion
		errResult     ErrorResponse
	)
	resp, err := c.Client.R().SetErrorResult(&errResult).SetSuccessResult(&successResult).Get(targetAPI)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, NewNotFoundError("model version", verID)
		}

		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		errResult.StatusCode = resp.StatusCode
		return nil, &errResult
	}

	return &successResult, nil
}

type ModelVersion struct {
	ID                   int         `json:"id,omitempty"`
	ModelID              int         `json:"modelId,omitempty"`
	Name                 string      `json:"name,omitempty"`
	CreatedAt            time.Time   `json:"createdAt,omitempty"`
	UpdatedAt            time.Time   `json:"updatedAt,omitempty"`
	TrainedWords         []string    `json:"trainedWords,omitempty"`
	BaseModel            string      `json:"baseModel,omitempty"`
	EarlyAccessTimeFrame int         `json:"earlyAccessTimeFrame,omitempty"`
	Description          interface{} `json:"description,omitempty"`
	// Stats only contains:
	//  DownloadCount int
	//	RatingCount   int
	//	Rating        int
	Stats       ModelStats           `json:"stats,omitempty"`
	Model       ModelVersionModel    `json:"model,omitempty"`
	Files       []ModelVersionFiles  `json:"files,omitempty"`
	Images      []ModelVersionImages `json:"images,omitempty"`
	DownloadURL string               `json:"downloadUrl,omitempty"`
}

func (m *ModelVersion) JSON() string {
	bs, _ := json.MarshalIndent(m, "", "\t")
	return string(bs)
}

type ModelVersionStats struct {
	DownloadCount int `json:"downloadCount,omitempty"`
	RatingCount   int `json:"ratingCount,omitempty"`
	Rating        int `json:"rating,omitempty"`
}
type ModelVersionModel struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Nsfw bool   `json:"nsfw,omitempty"`
	Poi  bool   `json:"poi,omitempty"`
}
type ModelVersionFiles struct {
	ID                int           `json:"id,omitempty"`
	Name              string        `json:"name,omitempty"`
	SizeKB            float64       `json:"sizeKB,omitempty"`
	Type              string        `json:"type,omitempty"`
	Metadata          ModelMetadata `json:"metadata,omitempty"`
	PickleScanResult  string        `json:"pickleScanResult,omitempty"`
	PickleScanMessage string        `json:"pickleScanMessage,omitempty"`
	VirusScanResult   string        `json:"virusScanResult,omitempty"`
	ScannedAt         time.Time     `json:"scannedAt,omitempty"`
	Hashes            ModelHashes   `json:"hashes,omitempty"`
	Primary           bool          `json:"primary,omitempty"`
	DownloadURL       string        `json:"downloadUrl,omitempty"`
}
type ModelVersionImages struct {
	URL    string      `json:"url,omitempty"`
	Nsfw   bool        `json:"nsfw,omitempty"`
	Width  int         `json:"width,omitempty"`
	Height int         `json:"height,omitempty"`
	Hash   string      `json:"hash,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}
