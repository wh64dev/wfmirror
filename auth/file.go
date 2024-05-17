package auth

import (
	"os"

	"github.com/devproje/plog/log"
	"github.com/wh64dev/wfcloud/util"
)

type CryptFile struct {
	Dirs []struct {
		Directory string `json:"dir"`
	} `json:"dirs"`
}

const (
	dirFilename = "wfconf/crypt_dir.json"
)

func queryDir() *CryptFile {
	res, err := util.ParseJSON[CryptFile](dirFilename)
	if err != nil {
		return nil
	}

	return res
}

func CheckAuthGlobal() bool {
	return false
}

func CheckAuth(dir string) bool {
	res := queryDir()
	for _, d := range res.Dirs {
		if d.Directory == dir {
			return true
		}
	}

	return false
}

func InitFile() {
	env := os.Getenv("GLOBAL_LOCK")
	switch env {
	case "1":
	case "true":
		return
	case "0":
	case "false":
		break
	default:
		log.Fatalln("`GLOBAL_LOCK` variable must be `0` or `1` and `true` or `false`")
	}

	if _, err := os.Stat(dirFilename); err != nil {
		os.WriteFile(dirFilename, []byte("{\"dirs\":[]}"), 0655)
	}
}
