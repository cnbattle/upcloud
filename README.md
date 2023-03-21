# upcloud

[![Build Status](https://github.com/cnbattle/upcloud/actions/workflows/vet.yml/badge.svg)](https://cloud.drone.io/cnbattle/upcloud)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)


上传前端项目到CDN平台及自动刷新节点缓存的工具

## USE
1. 下载对于平台可执行文件,放置到系统PATH目录下, 或
```shell script
go get -u github.com/cnbattle/upcloud
```
2. 根据使用的对象存储平台创建配置 `.upcloud.env` 文件, 具体参考下面示例

## TODO

### 功能

- [x] 多线程上传

### 多平台
- [x] 腾讯云
- [x] 七牛云
- [x] 阿里云

## 各平台`.upcloud.env`配置示例

- UP_CLOUD_PLATFORM： 平台
- UP_CLOUD_PATH： 需上传静态资源路径
- UP_CLOUD_POOL_SIZE： 上传时的并发数，默认为10
- UP_CLOUD_PREFETCH_URLS： 上传完成，要刷新缓存的链接，VUE React等静态站点一般刷新首页index.html即可

### 腾讯云 COS
```.env
UP_CLOUD_PLATFORM=tencent
UP_CLOUD_PATH=dist/
UP_CLOUD_POOL_SIZE=10
UP_CLOUD_PREFETCH_URLS=your prefetch urls (多个用,分割)

UP_CLOUD_SECRET_ID=your id 
UP_CLOUD_SECRET_KEY=your key
UP_CLOUD_VISIT_NODE=your visis node
```

### 七牛云
```.env
UP_CLOUD_PLATFORM=qiniu
UP_CLOUD_PATH=dist/
UP_CLOUD_POOL_SIZE=10
UP_CLOUD_PREFETCH_URLS=your prefetch urls (多个用,分割)

UP_CLOUD_ACCESS_KEY=your access key
UP_CLOUD_SECRET_KEY=your secret key
UP_CLOUD_BUCKET=your bucket
```

### 阿里云
```.env
UP_CLOUD_PLATFORM=aliyun
UP_CLOUD_PATH=dist/
UP_CLOUD_POOL_SIZE=10
UP_CLOUD_PREFETCH_URLS=your prefetch urls (多个用,分割)

UP_CLOUD_ENDPOINT=your endpoint
UP_CLOUD_ACCESS_KEY_ID=your access key id
UP_CLOUD_ACCESS_KEY_SECRET=your access key secret
UP_CLOUD_BUCKET=your bucket
```
