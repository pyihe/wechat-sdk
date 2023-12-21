package files

import (
	"os"
	"path"
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
	if err = MakeNewPath(filePath); err != nil {
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

// MakeNewPath 判断目录是否存在，如果不存在，则新建一个目录
func MakeNewPath(targetPath string) error {
	if _, err := os.Stat(targetPath); err != nil {
		if !os.IsExist(err) {
			//创建目录
			if mErr := os.MkdirAll(targetPath, os.ModePerm); mErr != nil {
				return mErr
			}
		}
	}
	return nil
}
