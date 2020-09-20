package utils

import (
	"os"
)

// CreateDir 调用os.MkdirAll递归创建文件夹
func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// IsExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsExistProjectConfig 判断项目名称对应的文件是否存在(返回true是存在)
func IsExistProjectConfig(projectName string) bool {
	path := GetConfigDir()
	if IsExist(path + projectName + ".json") {
		return true
	}
	return false
}

func GetConfigDir() string {
	home, err := Home()
	if err != nil {
		panic(err)
	}
	path := home + "/.config/upcloud/"
	err = CreateDir(path)
	if err != nil {
		panic(err)
	}
	return path
}

// GetExistProjectConfig 获取项目名称对应的文件的路径名称
func GetExistProjectConfig(projectName string) string {
	home, err := Home()
	if err != nil {
		panic(err)
	}
	path := home + "/.config/upcloud/"
	err = CreateDir(path)
	if err != nil {
		panic(err)
	}
	return path + projectName + ".json"
}
