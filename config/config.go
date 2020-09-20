package config

import (
	"errors"
	"strings"
)

func IsExitProjectName(name string) error {
	for _, i2 := range Conf {
		if strings.EqualFold(i2.ProjectName, name) {
			return errors.New("is exit")
		}
	}
	return nil
}

func GetProjectConfig(name string) (ProjectConfig, error) {
	for _, i2 := range Conf {
		if strings.EqualFold(i2.ProjectName, name) {
			return i2, nil
		}
	}
	return ProjectConfig{}, errors.New("不存在该项目")
}
