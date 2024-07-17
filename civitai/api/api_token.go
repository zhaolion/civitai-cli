package api

import (
	"os"
	"sync"

	"github.com/mitchellh/go-homedir"
)

const (
	TokenEnvKey   = "CIVITAI_API_KEY"
	ConfigDir     = "~/.civitai"
	TokenFilePath = "~/.civitai/token"
)

var (
	apiToken string

	loadTokenOnce sync.Once
)

func absConfigDir() string {
	dir, _ := homedir.Expand(ConfigDir)
	return dir
}

func absTokenFilePath() string {
	file, _ := homedir.Expand(TokenFilePath)
	return file
}

func SetAPIToken(token string) {
	configDir := absConfigDir()

	// check config dir is already exist.
	_, err := os.ReadDir(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(configDir, 0744); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	// touch and write to token file
	tokenFile := absTokenFilePath()
	if err := os.WriteFile(tokenFile, []byte(token), 0644); err != nil {
		panic(err)
	}
}

func GetAPIToken() string {
	if apiToken != "" {
		return apiToken
	}

	// load from env or file
	loadTokenOnce.Do(func() {
		apiTokenEnv := os.Getenv(TokenEnvKey)

		if apiTokenEnv == "" {
			bs, err := os.ReadFile(absTokenFilePath())
			if err != nil {
				panic(err)
			}

			apiToken = string(bs)
		} else {
			apiToken = apiTokenEnv
		}
	})

	return apiToken
}

func GetAPITokenMask() string {
	token := GetAPIToken()
	rs := []rune(token)
	for i := 0; i < len(rs)-4; i++ {
		rs[i] = '*'
	}
	return string(rs)
}
