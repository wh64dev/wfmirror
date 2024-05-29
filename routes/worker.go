package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/util"
)

const uploadBaseDir = "data"

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

func (dw *DirWorker) UploadFile(ctx *gin.Context) {
	dirname := ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request: %v", err)
		return
	}

	uploadDir := filepath.Join(uploadBaseDir, filepath.FromSlash(dirname))
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.String(http.StatusInternalServerError, "Could not create upload directory: %v", err)
		return
	}

	filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.String(http.StatusInternalServerError, "Could not save file: %v", err)
		return
	}

	ctx.String(http.StatusOK, "File uploaded successfully: %s", file.Filename)
}

func (dw *DirWorker) RawFiles(ctx *gin.Context) {
	path := ctx.Param("filepath")
	filePath := filepath.Join(uploadBaseDir, filepath.FromSlash(path))

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.String(http.StatusNotFound, "File not found: %s", ctx.Param("filepath"))
		return
	}

	_, file := filepath.Split(path)
	ctx.FileAttachment(filePath, file)
}

func (dw *DirWorker) ListFiles(ctx *gin.Context) {
	var start = time.Now()
	var dirname = ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}

	baseDir := filepath.Join(uploadBaseDir, dirname)
	var directory []*FileData
	file, err := os.Stat(baseDir)
	if err != nil {
		ctx.Status(404)
		return
	}

	if !file.IsDir() {
		ctx.Redirect(301, fmt.Sprintf("/d/%s", dirname))
		return
	}

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		ctx.Status(404)
		return
	}

	var files []*FileData
	for _, entry := range entries {
		format := "01-02-2006 03:04"
		finfo, _ := entry.Info()
		ftype := FILE
		url := fmt.Sprintf("%s%s", dirname, entry.Name())
		if dirname != "" {
			if dirname[:len(dirname)-1] != "/" {
				url = fmt.Sprintf("%s/%s", dirname, entry.Name())
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
	ctx.JSON(http.StatusOK, gin.H{
		"dir":          dirname,
		"data":         directory,
		"respond_time": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	})
}
