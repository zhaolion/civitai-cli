package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

// TODO: 提供选择函数来决定下载哪一个文件
func (c *CivitaiClient) ModelDownloadByID(modelID, modelVerID, targetDir string) error {
	m, err := c.ModelInfoByID(modelID)
	if err != nil {
		return err
	}

	ver := m.MatchVerByID(modelVerID)
	if ver == nil {
		return fmt.Errorf("model version not found - %s", modelVerID)
	}
	if len(ver.Files) == 0 {
		return nil
	}

	file := ver.Files[0]

	return c.downloadWithProgressBar(file.SizeKB, file.DownloadURL, targetDir)
}

func (c *CivitaiClient) downloadWithProgressBar(fileSizeKb float64, fileURL, targetFile string) error {
	totalSize := int(fileSizeKb) * 1024

	bar := progressbar.NewOptions(totalSize,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //you should install "github.com/k0kubun/go-ansi"
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	var (
		errResult ErrorResponse
	)
	resp, err := c.Client.R().
		SetOutputFile(targetFile).
		SetDownloadCallbackWithInterval(func(info req.DownloadInfo) {
			_ = bar.Set(int(info.DownloadedSize))
		}, time.Microsecond).
		Get(fileURL)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return NewNotFoundError("file", fileURL)
		}

		return err
	}

	for k, v := range resp.Header {
		fmt.Println("header:", k, v)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		errResult.StatusCode = resp.StatusCode
		return &errResult
	}

	return nil
}
