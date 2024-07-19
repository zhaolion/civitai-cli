package api

import (
	"context"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
	"github.com/zhaolion/civitai-cli/civitai/util"
)

// FileDownload download file from fileURL to targetFile.
func (c *CivitaiClient) FileDownload(ctx context.Context, file *ModelInfoFiles, targetFile string, opts ...*FileDownloadOption) error {
	if err := util.CheckTargetFileNotExists(targetFile); err != nil {
		return err
	}

	var opt *FileDownloadOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	opt = normalizeFileDownloadOption(opt)

	downTimeout := time.Duration(opt.DownloadTimeoutSec) * time.Second
	ctxTimeout := time.Duration(opt.DownloadTimeoutSec+60) * time.Second
	fileURL := file.DownloadURL

	// build request
	cctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	var errResult ErrorResponse
	downloadReq := c.Client.Clone().SetTimeout(downTimeout).R().
		SetContext(cctx).
		SetOutputFile(targetFile)
	if opt.Callback != nil {
		downloadReq = downloadReq.SetDownloadCallbackWithInterval(opt.Callback, time.Microsecond)
	}

	// fetch file
	resp, err := downloadReq.Get(fileURL)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return NewNotFoundError("file", fileURL)
		}

		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		errResult.StatusCode = resp.StatusCode
		return &errResult
	}

	return nil
}

type FileDownloadOption struct {
	DownloadTimeoutSec uint32
	Callback           req.DownloadCallback
}

func normalizeFileDownloadOption(input *FileDownloadOption) *FileDownloadOption {
	if input == nil {
		return &FileDownloadOption{
			DownloadTimeoutSec: 60 * 30,
		}
	}
	if input.DownloadTimeoutSec == 0 {
		input.DownloadTimeoutSec = 60 * 30
	}

	return input
}
