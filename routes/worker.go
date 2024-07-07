package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/util"
)

type DirWorker struct{}

type FileType string

const (
	FILE FileType = "file"
	DIR  FileType = "dir"
)

type FileData struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Size     string `json:"size"`
	Type     string `json:"type"`
	Modified string `json:"modified"`
}

func worker(base, dir string) ([]*FileData, error) {
	entries, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}

	var files []*FileData
	var directory []*FileData
	for _, entry := range entries {
		format := "01-02-2006 03:04"
		finfo, _ := entry.Info()
		ftype := FILE
		url := fmt.Sprintf("%s%s", dir, entry.Name())
		if dir != "" {
			if dir[:len(dir)-1] != "/" {
				url = fmt.Sprintf("%s/%s", dir, entry.Name())
			}
		}

		if entry.IsDir() {
			ftype = DIR

			directory = append(directory, &FileData{
				URL:      url,
				Name:     entry.Name(),
				Size:     util.FSize(float64(finfo.Size())),
				Type:     string(ftype),
				Modified: finfo.ModTime().Format(format),
			})

			continue
		}

		files = append(files, &FileData{
			URL:      url,
			Name:     entry.Name(),
			Size:     util.FSize(float64(finfo.Size())),
			Type:     string(ftype),
			Modified: finfo.ModTime().Format(format),
		})
	}

	directory = append(directory, files...)
	return directory, nil
}

func (dw *DirWorker) List(ctx *gin.Context) {
	var start = time.Now()
	var dirname = ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}

	cnf := config.Get()
	baseDir := filepath.Join(cnf.Global.DataDir, dirname)
	file, err := os.Stat(baseDir)
	if err != nil {
		ctx.Status(404)
		return
	}

	if !file.IsDir() {
		ctx.File(baseDir)
		return
	}

	directory, err := worker(baseDir, dirname)
	if err != nil {
		ctx.Status(404)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ok":           1,
		"dir":          dirname,
		"data":         directory,
		"respond_time": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	})
}
