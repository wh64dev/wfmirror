package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/service"
	"github.com/wh64dev/wfcloud/service/auth"
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

func checkAuth(ctx *gin.Context) bool {
	_, validation := auth.Validate(ctx, false)
	return validation
}

func (dw *DirWorker) UploadFile(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	dirname := ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{
			"ok":    0,
			"errno": fmt.Sprintf("Bad request: %v", err),
		})
		return
	}

	uploadDir := filepath.Join(uploadBaseDir, filepath.FromSlash(dirname))
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": fmt.Sprintf("Could not create upload directory: %v", err),
		})
		return
	}

	filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.String(http.StatusInternalServerError, "Could not save file: %v", err)
		return
	}

	log.Infof("File uploaded successfully: %s\n", file.Filename)
	ctx.Status(200)
}

func (dw *DirWorker) AddSecret(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	dirname := ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}

	priv := new(service.PrivDir)

	err := priv.Add(dirname)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"errno":  fmt.Sprintf("Cannot add private directory: %v", err),
		})
		return
	}

	ctx.Status(200)
}

func (dw *DirWorker) DropSecret(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	id := ctx.Param("id")
	priv := new(service.PrivDir)

	numberId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"ok":    0,
			"errno": "id must be number",
		})
		return
	}

	err = priv.Drop(int(numberId))
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": err,
		})
		return
	}
}

func (dw *DirWorker) QuerySecret(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	priv := new(service.PrivDir)

	dirs, err := priv.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":   1,
		"dirs": dirs,
	})
}

func (dw *DirWorker) List(ctx *gin.Context) {
	var start = time.Now()
	var dirname = ctx.Param("dirname")
	if dirname == "/" || dirname == "/root" {
		dirname = ""
	}

	baseDir := filepath.Join(uploadBaseDir, dirname)
	file, err := os.Stat(baseDir)
	if err != nil {
		ctx.Status(404)
		return
	}

	if !file.IsDir() {
		ctx.FileAttachment(baseDir, file.Name())
		return
	}

	directory, err := worker(baseDir, dirname)
	if err != nil {
		ctx.Status(404)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"dir":          dirname,
		"data":         directory,
		"respond_time": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	})
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
