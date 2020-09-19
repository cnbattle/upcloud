package cloud

import "errors"

// Platform 平台列表
var Platform []string

// SelectPlatform 选择平台
func SelectPlatform(platform string) (CommInterface, error) {
	switch platform {
	case "qiniu":
		var qiniu Qiniu
		return &qiniu, nil
	default:
		return nil, errors.New("platform is error:" + platform)
	}
}
