package api

const (
	URLModelInfoByID        = "/v1/models/%s"
	URLModelVersionInfoByID = "/v1/model-versions/%s"
)

const (
	CivitaiAPIBaseV1 = "https://civitai.com/api/"
	UserAgent        = "civitai-go-client (https://github.com/zhaolion/civitai-cli)"
)

const (
	ModelTypeCheckpoint        = "Checkpoint"
	ModelTypeTextualInversion  = "TextualInversion"
	ModelTypeHypernetwork      = "Hypernetwork"
	ModelTypeAestheticGradient = "AestheticGradient"
	ModelTypeLORA              = "LORA"
	ModelTypeControlnet        = "Controlnet"
	ModelTypePoses             = "Poses"
)
