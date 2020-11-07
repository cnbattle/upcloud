package main

import (
	"errors"
	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/core/cloud"
)

func selectInterFace(platform string) (cloud.CommInterface, error) {
	switch platform {
	case "qiniu":
		platformStruct := cloud.Qiniu{
			AccessKey: config.GetEnv("UP_CLOUD_ACCESS_KEY"),
			SecretKey: config.GetEnv("UP_CLOUD_SECRET_KEY"),
			Bucket:    config.GetEnv("UP_CLOUD_BUCKET"),
		}
		return &platformStruct, nil
	case "tencent":
		platformStruct := cloud.Tencent{
			SecretID:  config.GetEnv("UP_CLOUD_SECRET_ID"),
			SecretKey: config.GetEnv("UP_CLOUD_SECRET_KEY"),
			VisitNode: config.GetEnv("UP_CLOUD_VISIT_NODE"),
		}
		return &platformStruct, nil
	case "aliyun":
		platformStruct := cloud.Aliyun{
			Endpoint:        config.GetEnv("UP_CLOUD_ENDPOINT"),
			AccessKeyID:     config.GetEnv("UP_CLOUD_ACCESS_KEY_ID"),
			AccessKeySecret: config.GetEnv("UP_CLOUD_ACCESS_KEY_SECRET"),
			Bucket:          config.GetEnv("UP_CLOUD_BUCKET"),
		}
		return &platformStruct, nil
	default:
		return nil, errors.New("config platform is error:" + platform)
	}
}
