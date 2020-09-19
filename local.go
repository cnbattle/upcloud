package main

import (
	"os"
	"path/filepath"
)

func local(pwd string) (files []string) {
	//获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path[len(pwd)+1:])
		}
		return nil
	})
	return
}
