package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	URL      string
	Name     string
	Size     string
	Type     string
	Modified string
}

func (dw *DirWorker) CreateDir(ctx *gin.Context) {
	dirname := ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}
	dirPath := filepath.Join(uploadBaseDir, filepath.FromSlash(dirname))

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		ctx.String(http.StatusInternalServerError, "Could not create directory: %v", err)
		return
	}

	ctx.String(http.StatusOK, "Directory created successfully: %s", dirname)
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

func (dw *DirWorker) DownloadFile(ctx *gin.Context) {
	path := ctx.Param("filepath")
	filePath := filepath.Join(uploadBaseDir, filepath.FromSlash(path))

	data, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.String(http.StatusNotFound, "File not found: %s", ctx.Param("filepath"))
		return
	}

	if data.IsDir() {
		ctx.Redirect(301, fmt.Sprintf("/f/%s", strings.Replace(path, "data/", "", 1)))
	}

	ctx.FileAttachment(filePath, data.Name())
}

func (dw *DirWorker) ListFiles(ctx *gin.Context) {
	var dirname = ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}

	baseDir := filepath.Join(uploadBaseDir, dirname)

	var files []*FileData
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

	for _, entry := range entries {
		format := "01-02-2006 03:04"
		finfo, _ := entry.Info()
		ftype := FILE
		if entry.IsDir() {
			ftype = DIR
		}

		url := fmt.Sprintf("%s%s", dirname, entry.Name())
		if dirname[:len(dirname)-1] != "/" {
			url = fmt.Sprintf("%s/%s", dirname, entry.Name())
		}
		files = append(files, &FileData{
			URL:      url,
			Name:     entry.Name(),
			Size:     util.FSize(float64(finfo.Size())),
			Type:     string(ftype),
			Modified: finfo.ModTime().Format(format),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"dir":  dirname,
		"data": files,
	})
}
