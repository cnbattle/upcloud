package cloud

import (
	"context"
	"github.com/cnbattle/upcloud/config"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func init() {
	//Platform = append(Platform, "tencent")
}

type Tencent struct {
	SecretID  string      `json:"secret_id"`
	SecretKey string      `json:"secret_key"`
	ApiURL    string      `json:"api_url"`
	client    *cos.Client `json:"-"`
}

func (t *Tencent) Init() error {
	u, err := url.Parse(t.ApiURL)
	if err != nil {
		return err
	}
	b := &cos.BaseURL{BucketURL: u}
	t.client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretID,
			SecretKey: t.SecretKey,
		},
	})
	return nil
}

func (t *Tencent) Setting() config.ProjectConfig {
	panic("implement me")
}

func (t *Tencent) GetAll() (list []string, err error) {

	opt := &cos.BucketGetOptions{
		Prefix: "",
		//MaxKeys: 3,
	}
	v, _, err := t.client.Bucket.Get(context.Background(), opt)
	if err != nil {
		return
	}
	for _, c := range v.Contents {
		list = append(list, c.Key)
	}
	return
}

func (t *Tencent) DelAll(list []string) error {

	for _, item := range list {
		_, err := t.client.Object.Delete(context.Background(), item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tencent) Upload(localFile, upKey string) error {
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//name := "test/objectPut.go"
	// 1.通过字符串上传对象
	//f := strings.NewReader("test")
	//
	//_, err := t.client.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	return err
	//}
	// 2.通过本地文件上传对象
	_, err := t.client.Object.PutFromFile(context.Background(), upKey, localFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tencent) Prefetch() error {
	panic("implement me")
}
