package main

import (
	"github.com/cnbattle/upcloud/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey = config.GetDefaultEnv("accessKey", "aAfbj7ZVCWpsuVEnR-0nhLHFX28Fi7b_0-h9L4Cm")
	secretKey = config.GetDefaultEnv("secretKey", "-nWGLHbT9dpb_d6VC4hlrxc0Jf7wEfsmttDAghBS")
	bucket    = config.GetDefaultEnv("bucket", "agent-h5")
)

func init() {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac = qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager = storage.NewBucketManager(mac, &cfg)

	upToken = putPolicy.UploadToken(mac)
}
