package api

import "time"

const (
	APIModelInfoByID = "/api/v1/models/%d"
)

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
	Stats                 ModelInfoStats      `json:"stats,omitempty"`
	Creator               ModelInfoCreator    `json:"creator,omitempty"`
	Tags                  []string            `json:"tags,omitempty"`
	ModelVersions         []ModelInfoVersions `json:"modelVersions,omitempty"`
}

type ModelInfoStats struct {
	DownloadCount   int `json:"downloadCount,omitempty"` // version 公共字段 - begin
	RatingCount     int `json:"ratingCount,omitempty"`
	Rating          int `json:"rating,omitempty"`
	ThumbsUpCount   int `json:"thumbsUpCount,omitempty"`
	ThumbsDownCount int `json:"thumbsDownCount,omitempty"` // version 公共字段 - end

	FavoriteCount     int `json:"favoriteCount,omitempty"`
	CommentCount      int `json:"commentCount,omitempty"`
	TippedAmountCount int `json:"tippedAmountCount,omitempty"`
}

type ModelInfoCreator struct {
	Username string `json:"username,omitempty"`
	Image    string `json:"image,omitempty"`
}

type ModelInfoMetadata struct {
	Format string `json:"format,omitempty"`
	Size   string `json:"size,omitempty"`
	Fp     string `json:"fp,omitempty"`
}

type ModelInfoHashes struct {
	AutoV1 string `json:"AutoV1,omitempty"`
	AutoV2 string `json:"AutoV2,omitempty"`
	Sha256 string `json:"SHA256,omitempty"`
	Crc32  string `json:"CRC32,omitempty"`
	Blake3 string `json:"BLAKE3,omitempty"`
}

type ModelInfoFiles struct {
	ID                int               `json:"id,omitempty"`
	SizeKB            float64           `json:"sizeKB,omitempty"`
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	PickleScanResult  string            `json:"pickleScanResult,omitempty"`
	PickleScanMessage string            `json:"pickleScanMessage,omitempty"`
	VirusScanResult   string            `json:"virusScanResult,omitempty"`
	VirusScanMessage  interface{}       `json:"virusScanMessage,omitempty"`
	ScannedAt         time.Time         `json:"scannedAt,omitempty"`
	Metadata          ModelInfoMetadata `json:"metadata,omitempty"`
	Hashes            ModelInfoHashes   `json:"hashes,omitempty"`
	DownloadURL       string            `json:"downloadUrl,omitempty"`
	Primary           bool              `json:"primary,omitempty"`
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
	Stats         ModelInfoStats    `json:"stats,omitempty"`
	Files         []ModelInfoFiles  `json:"files,omitempty"`
	Images        []ModelInfoImages `json:"images,omitempty"`
	DownloadURL   string            `json:"downloadUrl,omitempty"`
}
