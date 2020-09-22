package cloud

import (
	"context"
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/cdn"
	"github.com/qiniu/api.v7/v7/sms/rpc"
	"github.com/qiniu/api.v7/v7/storage"
)

func init() {
	Platform = append(Platform, "qiniu")
}

// Qiniu 七牛云
type Qiniu struct {
	AccessKey     string                 `json:"access_key"`
	SecretKey     string                 `json:"secret_key"`
	Bucket        string                 `json:"bucket"`
	UpToken       string                 `json:"-"`
	Mac           *qbox.Mac              `json:"-"`
	BucketManager *storage.BucketManager `json:"-"`
}

// Init 初始化
func (q *Qiniu) Init() error {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	q.Mac = qbox.NewMac(q.AccessKey, q.SecretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	q.BucketManager = storage.NewBucketManager(q.Mac, &cfg)

	q.UpToken = putPolicy.UploadToken(q.Mac)
	return nil
}

// GetAll 获取全部文件key
func (q *Qiniu) GetAll() (list []string, err error) {
	limit := 1000
	prefix := ""
	delimiter := ""
	//初始列举marker为空
	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := q.BucketManager.ListFiles(q.Bucket, prefix, delimiter, marker, limit)
		if err != nil {
			return nil, err
		}
		//print entries
		for _, entry := range entries {
			list = append(list, entry.Key)
		}
		if hasNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
	return list, nil
}

// DelAll 批量删除
func (q *Qiniu) DelAll(list []string) error {
	deleteOps := make([]string, 0, len(list))
	for _, key := range list {
		deleteOps = append(deleteOps, storage.URIDelete(q.Bucket, key))
	}
	rets, err := q.BucketManager.Batch(deleteOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*rpc.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	}
	return nil
}

// Upload 上传
func (q *Qiniu) Upload(localFile, upKey string) error {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFile(context.Background(), &ret, q.UpToken, upKey, localFile, nil)
	if err != nil {
		return err
	}
	return nil
}

// Prefetch 刷新
func (q *Qiniu) Prefetch() error {
	cdnManager := cdn.NewCdnManager(q.Mac)
	//刷新链接，单次请求链接不可以超过10个，如果超过，请分批发送请求
	urlsToRefresh := []string{
		"http://h5.ygxsj.com/",
	}
	_, err := cdnManager.RefreshDirs(urlsToRefresh)
	return err
}

// Setting 设置
func (q *Qiniu) Setting() config.ProjectConfig {
	var projectName, path, accessKey, secretKey, bucket string
START:
	fmt.Print("Project Name：")
	fmt.Scanln(&projectName)
	err := config.IsExitProjectName(projectName)
	if err != nil {
		fmt.Println("已存在，请重新输入")
		goto START
	}
	fmt.Print("Project Path：")
	fmt.Scanln(&path)
	fmt.Print("Qiniu AccessKey：")
	fmt.Scanln(&accessKey)
	fmt.Print("Qiniu SecretKey：")
	fmt.Scanln(&secretKey)
	fmt.Print("Qiniu Bucket：")
	fmt.Scanln(&bucket)

	project := config.ProjectConfig{
		ProjectName: projectName,
		Platform:    "qiniu",
		Path:        path,
		Args: map[string]string{
			"accessKey": accessKey,
			"secretKey": secretKey,
			"bucket":    bucket,
		},
	}
	return project
}
