package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type FilesStr struct {
	Local string
	UpKey string
}

// Local 获取指定目录下所有文件
func Local(pwd string) (files []FilesStr) {
	//获取当前目录下的所有文件或目录信息
	err := filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, handlePathToStruct(pwd, path))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

// handlePathToStruct handlePathToStruct
func handlePathToStruct(pwd, path string) FilesStr {
	path = strings.Replace(path, "\\", "/", -1)
	pwdLen := len(pwd)
	if !strings.EqualFold(pwd[len(pwd)-1:], "/") {
		pwdLen++
	}
	return FilesStr{
		Local: path,
		UpKey: path[pwdLen:],
	}
}
