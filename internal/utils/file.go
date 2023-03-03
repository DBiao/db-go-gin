package utils

import (
	"os"
	"path/filepath"
)

func CreatePathIfNotExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return nil
		}
	}
	return err
}

// GetFileList 获取所有文件
func GetFileList(dirPath string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if !f.IsDir() {
				fileList = append(fileList, path)
				return nil
			}

			return nil
		})
	return fileList, err
}

// PathExists 路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
