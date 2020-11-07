# upcloud

[![Build Status](https://cloud.drone.io/api/badges/cnbattle/upcloud/status.svg)](https://cloud.drone.io/cnbattle/upcloud)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)


上传前端项目到CDN平台的工具

## USE
1. 下载对于平台可执行文件,放置到系统PATH目录下, 或
```shell script
go get -u github.com/cnbattle/upcloud
```
2. 根据使用的对象存储平台创建配置 .env 文件, 具体参考下面示例

## TODO
- [x] Qiniu Cloud
- [x] Tencent Cloud
- [x] Aliyun
- [ ] 华为云
- [ ] 百度云
- [ ] 滴滴云
- [ ] UCloud
- [ ] 青云 QingCloud
- [ ] 京东智联云
- [ ] AWS
- [ ] Google Cloud

## 各平台配置示例

### 腾讯云 COS
```.env
UP_CLOUD_PLATFORM=tencent
UP_CLOUD_PATH=dist/

UP_CLOUD_SECRET_ID=your id 
UP_CLOUD_SECRET_KEY=your key
UP_CLOUD_VISIT_NODE=your visis node
UP_CLOUD_PREFETCH_URLS=your prefetch urls (多个用,分割)
```

### 七牛云
```.env
UP_CLOUD_PLATFORM=qiniu
UP_CLOUD_PATH=dist/

UP_CLOUD_ACCESS_KEY=your access key
UP_CLOUD_SECRET_KEY=your secret key
UP_CLOUD_BUCKET=your bucket
UP_CLOUD_PREFETCH_URLS=your prefetch urls (多个用,分割)
```

### 阿里云
```.env
UP_CLOUD_PLATFORM=aliyun
UP_CLOUD_PATH=dist/

UP_CLOUD_ENDPOINT=your endpoint
UP_CLOUD_ACCESS_KEY_ID=your access key id
UP_CLOUD_ACCESS_KEY_SECRET=your access key secret
UP_CLOUD_BUCKET=your bucket
UP_CLOUD_PREFETCH_URLS=your prefetch urls (多个用,分割)
```