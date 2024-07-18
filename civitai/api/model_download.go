package api

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

// ModelDownloadByID 下载指定 modelID 的模型文件 - 所有版本
// TODO: 提供选择函数来决定下载哪些个文件
func (c *CivitaiClient) ModelDownloadByID(ctx context.Context, modelID, targetDir string) error {
	m, err := c.ModelInfoByID(modelID)
	if err != nil {
		return err
	}

	wg, cctx := errgroup.WithContext(ctx)
	wg.SetLimit(3) // TODO: 替换为 config 选项 or 参数
	for _, ver := range m.ModelVersions {
		if len(ver.Files) == 0 {
			continue
		}

		ver := ver
		wg.Go(func() error {
			return c.downloadVersion(cctx, &ver, targetDir)
		})
	}
	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}

// ModelVerDownloadByID 下载指定 modelID 和 modelVerID 的模型文件
// TODO: 提供选择函数来决定下载哪一个文件
func (c *CivitaiClient) ModelVerDownloadByID(ctx context.Context, modelID, modelVerID, targetDir string) error {
	m, err := c.ModelInfoByID(modelID)
	if err != nil {
		return err
	}

	ver := m.MatchVerByID(modelVerID)
	if ver == nil {
		return fmt.Errorf("model version not found - %s", modelVerID)
	}
	return c.downloadVersion(ctx, ver, targetDir)
}

func (c *CivitaiClient) downloadVersion(ctx context.Context, ver *ModelInfoVersions, targetDir string) error {
	if len(ver.Files) == 0 {
		return nil
	}

	wg, cctx := errgroup.WithContext(ctx)
	wg.SetLimit(3) // TODO: 替换为 config 选项 or 参数
	for _, file := range ver.Files {
		file := file
		wg.Go(func() error {
			targetFile := filepath.Join(targetDir, file.Name)
			return c.downloadFileWithProgressBar(cctx, &file, targetFile)
		})
	}
	return wg.Wait()
}

func (c *CivitaiClient) downloadFileWithProgressBar(ctx context.Context, file *ModelInfoFiles, targetFile string) error {
	totalSize := int(file.SizeKB) * 1024
	fileURL := file.DownloadURL
	desc := fmt.Sprintf("[cli][%s] download...", file.Name)

	bar := progressbar.NewOptions(totalSize,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //you should install "github.com/k0kubun/go-ansi"
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(desc),
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

	// TODO: 替换为 config 选项 or 参数
	cctx, cancel := context.WithTimeout(ctx, time.Minute*45)
	defer cancel()
	resp, err := c.Client.Clone().SetTimeout(30*time.Minute).R().
		SetContext(cctx).
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

	if resp.StatusCode >= http.StatusBadRequest {
		errResult.StatusCode = resp.StatusCode
		return &errResult
	}

	return nil
}
