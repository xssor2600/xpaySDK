package utils

import (
	"io"
	"os"
	"path/filepath"
)

func ReadPemFile(filePath string) ([]byte, error) {
	// 确保文件路径以.pem结尾
	if filepath.Ext(filePath) != ".pem" {
		return nil, os.ErrInvalid
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // 确保在函数返回时关闭文件

	// 读取文件的全部内容
	pemBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return pemBytes, nil
}
