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
			AccessKey: config.GetEnvForPanic("UP_CLOUD_ACCESS_KEY"),
			SecretKey: config.GetEnvForPanic("UP_CLOUD_SECRET_KEY"),
			Bucket:    config.GetEnvForPanic("UP_CLOUD_BUCKET"),
		}
		return &platformStruct, nil
	case "tencent":
		platformStruct := cloud.Tencent{
			SecretID:  config.GetEnvForPanic("UP_CLOUD_SECRET_ID"),
			SecretKey: config.GetEnvForPanic("UP_CLOUD_SECRET_KEY"),
			VisitNode: config.GetEnvForPanic("UP_CLOUD_VISIT_NODE"),
		}
		return &platformStruct, nil
	case "aliyun":
		platformStruct := cloud.Aliyun{
			Endpoint:        config.GetEnvForPanic("UP_CLOUD_ENDPOINT"),
			AccessKeyID:     config.GetEnvForPanic("UP_CLOUD_ACCESS_KEY_ID"),
			AccessKeySecret: config.GetEnvForPanic("UP_CLOUD_ACCESS_KEY_SECRET"),
			Bucket:          config.GetEnvForPanic("UP_CLOUD_BUCKET"),
		}
		return &platformStruct, nil
	default:
		return nil, errors.New("config platform is error:" + platform)
	}
}
