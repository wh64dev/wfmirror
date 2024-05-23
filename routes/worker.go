package routes

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/util"
)

type FileData struct {
	URL      string
	Name     string
	Size     string
	Modified string
}

func read(path string) []*FileData {
	dir, _ := os.ReadDir(fmt.Sprintf("data/%s", path))
	var back string
	var files []*FileData

	defDir := ""

	if path != defDir {
		split := strings.Split(path, "/")
		split = split[:len(split)-1]

		back = "/"
		for i, p := range split {
			if i == len(split)-1 {
				back += p
				break
			}

			back += fmt.Sprintf("%s/", p)
		}

		if back == "" {
			back = "../"
		}

		files = append(files, &FileData{
			URL:      back,
			Name:     "../",
			Size:     "-",
			Modified: "-",
		})
	}

	for _, file := range dir {
		directory := fmt.Sprintf("/%s/%s", path, file.Name())
		if path == defDir {
			directory = fmt.Sprint(file.Name())
		}

		format := "01-02-2006 03:04"
		var name, size string
		finfo, _ := file.Info()
		modified := finfo.ModTime().Format(format)
		if file.IsDir() {
			name = file.Name() + "/"
			size = "-"
			modified = "-"
		} else {
			name = file.Name()
			size = util.FSize(float64(finfo.Size()))
		}

		files = append(files, &FileData{
			URL:      directory,
			Name:     name,
			Size:     size,
			Modified: modified,
		})
	}

	return files
}

func DirWorker(ctx *gin.Context, path string) {
	dir := fmt.Sprintf("data/%s", path)
	file, err := os.Stat(dir)
	if err != nil {
		ctx.JSON(404, gin.H{
			"status": 404,
			"error":  err.Error(),
		})
		return
	}

	if !file.IsDir() {
		ctx.FileAttachment(dir, file.Name())
		return
	}

	data := read(path)
	ctx.JSON(200, gin.H{
		"dir":  fmt.Sprintf("/%s", path),
		"data": data,
	})
}
