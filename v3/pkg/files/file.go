package files

import (
	"os"
	"strings"

	"github.com/pyihe/go-pkg/files"
)

// WritToFile 将data写入到指定目录的指定文件里
func WritToFile(path, fileName string, data []byte) (err error) {
	//判断是否存在目标目录，如果不存在则创建
	if !strings.HasSuffix(path, "/") {
		if path == "" {
			path = "./"
		} else {
			path += "/"
		}
	}
	if err = files.MakeNewPath(path); err != nil {
		return err
	}

	f, err := os.Create(path + fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return
}
