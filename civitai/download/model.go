package download

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"github.com/zhaolion/civitai-cli/civitai/api"
	"github.com/zhaolion/civitai-cli/civitai/util"
	"golang.org/x/sync/errgroup"
)

// ModelDownloadByID 下载指定 modelID 的模型文件 - 所有版本
func (c *Client) ModelDownloadByID(ctx context.Context, modelID, targetDir string, opts ...*ModelDownloadOption) error {
	var opt *ModelDownloadOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	opt = normalizeFileDownloadOption(opt)

	// check dir
	if err := util.CheckTargetDirExists(targetDir); err != nil {
		return err
	}

	// fetch model info and check model is valid.
	m, err := c.api.ModelInfoByID(modelID)
	if err != nil {
		return err
	}

	files, dirMapping := chooseModelFileDirMapping(targetDir, m, opt)
	if len(files) == 0 {
		// no file to download
		return nil
	}

	// download single file
	if len(files) == 1 {
		fileDir := dirMapping[files[0].ID]
		targetFile := filepath.Join(fileDir, files[0].Name)
		return c.downloadFileWithProgressBar(ctx, files[0], targetFile, opt)
	}

	// download specific versions
	targetFileMapping := make(map[int]string)
	for _, file := range files {
		// TODO: 支持检查目录文件是否已经存在，可以跳过？or ignore
		targetFile := filepath.Join(dirMapping[file.ID], file.Name)
		targetFileMapping[file.ID] = targetFile
	}
	return c.downloadBatchFileWithProgressBar(ctx, files, targetFileMapping, opt)
}

// downloadFileWithProgressBar download one file with progress bar
func (c *Client) downloadFileWithProgressBar(ctx context.Context, file *api.ModelInfoFiles, targetFile string, opt *ModelDownloadOption) error {
	totalSize := int(file.SizeKB) * 1024
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

	err := c.api.FileDownload(ctx, file, targetFile, &api.FileDownloadOption{
		DownloadTimeoutSec: opt.DownloadTimeoutSec,
		Callback: func(info req.DownloadInfo) {
			_ = bar.Set(int(info.DownloadedSize))
		},
	})

	return err
}

func (c *Client) downloadBatchFileWithProgressBar(ctx context.Context, files []*api.ModelInfoFiles, targetFiles map[int]string, opt *ModelDownloadOption) error {
	// calculate total size
	var totalSize int
	for _, file := range files {
		totalSize += int(file.SizeKB) * 1024
	}
	// download desc
	desc := fmt.Sprintf("[cli][%d files] download...", len(files))
	// download size state, key is file id, value is downloaded size
	sizeState := sync.Map{}

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

	// download specific versions
	wg, cctx := errgroup.WithContext(ctx)
	wg.SetLimit(int(opt.DownloadBatchSize))
	for _, file := range files {
		file := file
		wg.Go(func() error {
			targetFile := targetFiles[file.ID]
			err := c.api.FileDownload(cctx, file, targetFile, &api.FileDownloadOption{
				DownloadTimeoutSec: opt.DownloadTimeoutSec,
				Callback: func(info req.DownloadInfo) {
					sizeState.Store(file.ID, int(info.DownloadedSize))

					// update bar
					var totalDownloadedSize int
					sizeState.Range(func(_, value interface{}) bool {
						size := value.(int)
						totalDownloadedSize += size
						return true
					})

					_ = bar.Set(totalDownloadedSize)
				},
			})
			return err
		})
	}
	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}

func chooseModelFileDirMapping(targetDir string, m *api.ModelInfo, opt *ModelDownloadOption) ([]*api.ModelInfoFiles, map[int]string) {
	files := make([]*api.ModelInfoFiles, 0)
	dirMapping := make(map[int]string)

	for _, ver := range m.ModelVersions {
		if len(opt.VersionIDList) != 0 {
			if !util.Int64sContains(opt.VersionIDList, int64(ver.ID)) {
				continue
			}
		}
		if len(opt.VersionNameList) != 0 {
			if !util.StringsContains(opt.VersionNameList, ver.Name) {
				continue
			}
		}

		for _, file := range ver.Files {
			if len(opt.FileIDList) != 0 {
				if !util.Int64sContains(opt.FileIDList, int64(file.ID)) {
					continue
				}
			}
			if len(opt.FileNameList) != 0 {
				if !util.StringsContains(opt.VersionNameList, file.Name) {
					continue
				}
			}

			// TODO: 支持 comfyUI model dir mapping.
			files = append(files, &file)
			dirMapping[file.ID] = targetDir
		}
	}

	return files, dirMapping
}

type ModelDownloadOption struct {
	DownloadBatchSize uint32
	// DownloadTimeoutSec is the timeout for downloading a file in seconds.
	DownloadTimeoutSec uint32
	// VersionIDList is a list of version id to download. If empty, download all versions.
	VersionIDList []int64
	// VersionNameList is a list of version name to download. If empty, download all versions.
	VersionNameList []string
	// FileIDList is a list of file id to download. If empty, download all files.
	FileIDList []int64
	// FileNameList is a list of file name to download. If empty, download all files.
	FileNameList []string
}

func normalizeFileDownloadOption(input *ModelDownloadOption) *ModelDownloadOption {
	if input == nil {
		return &ModelDownloadOption{
			DownloadBatchSize:  1,
			DownloadTimeoutSec: 60 * 30,
		}
	}
	if input.DownloadBatchSize == 0 {
		input.DownloadBatchSize = 1
	}
	if input.DownloadTimeoutSec == 0 {
		input.DownloadTimeoutSec = 60 * 30
	}

	return input
}
