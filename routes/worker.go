package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadBaseDir = "data"

type DirWorker struct{}

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

	// Ensure the upload directory exists
	uploadDir := filepath.Join(uploadBaseDir, filepath.FromSlash(dirname))
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.String(http.StatusInternalServerError, "Could not create upload directory: %v", err)
		return
	}

	// Save the uploaded file
	filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.String(http.StatusInternalServerError, "Could not save file: %v", err)
		return
	}

	ctx.String(http.StatusOK, "File uploaded successfully: %s", file.Filename)
}

func (dw *DirWorker) DownloadFile(ctx *gin.Context) {
	filePath := filepath.Join(uploadBaseDir, filepath.FromSlash(ctx.Param("filepath")))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.String(http.StatusNotFound, "File not found: %s", ctx.Param("filepath"))
		return
	}

	ctx.File(filePath)
}

func (dw *DirWorker) ListFiles(ctx *gin.Context) {
	dirname := ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}
	var files []string
	uploadDir := filepath.Join(uploadBaseDir, filepath.FromSlash(dirname))

	err := filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relativePath, err := filepath.Rel(uploadDir, path)
			if err != nil {
				return err
			}
			files = append(files, relativePath)
		}
		return nil
	})
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not list files: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, files)
}
