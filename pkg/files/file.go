package files

import (
	"os"
	"path"

	"github.com/pyihe/go-pkg/files"
)

// WritToFile 将data写入到指定目录的指定文件里
func WritToFile(filePath, fileName string, data []byte) (err error) {
	//判断是否存在目标目录，如果不存在则创建
	if filePath == "" {
		filePath = "files"
	}
	if fileName == "" {
		fileName = "newfile"
	}
	if err = files.MakeNewPath(filePath); err != nil {
		return err
	}

	f, err := os.Create(path.Join(filePath, fileName))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return
}
