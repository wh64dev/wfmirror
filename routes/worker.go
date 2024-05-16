package routes

import (
	"fmt"
	"html/template"
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

func arrToStr(arr []FileData) *string {
	var str = ""
	for _, i := range arr {
		str += createElement(i.URL, i.Name, i.Size, i.Modified)
	}

	return &str
}

func createElement(path, name, size, modified string) string {
	return fmt.Sprintf(
		`<a class="file_item animated" href=%s>
			<p class="file_name">%s</p>
			<p>%s</p>
			<p>%s</p>
		</a>`,
		path,
		name,
		size,
		modified,
	)
}

func MirrorWorker(ctx *gin.Context, path string) {
	iPath := fmt.Sprintf("data/%s", path)
	file, err := os.Stat(iPath)
	if err != nil {
		ctx.JSON(404, gin.H{
			"status": 404,
			"error":  err.Error(),
		})
		return
	}

	if !file.IsDir() {
		ctx.FileAttachment(iPath, file.Name())
		return
	}

	dir := read(path)
	ctx.HTML(200, "index.html", gin.H{
		"dir":     path,
		"content": template.HTML(*dir),
	})
}

func read(path string) *string {
	dir, _ := os.ReadDir(fmt.Sprintf("data/%s", path))
	var back string
	var files []FileData

	if path != "/" {
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

		files = append(files, FileData{
			URL:      back,
			Name:     "../",
			Size:     "-",
			Modified: "-",
		})
	}

	for _, file := range dir {
		ph := fmt.Sprintf("/%s/%s", path, file.Name())
		if path == "/" {
			ph = fmt.Sprint(file.Name())
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

		files = append(files, FileData{
			URL:      ph,
			Name:     name,
			Size:     size,
			Modified: modified,
		})
	}

	return arrToStr(files)
}
