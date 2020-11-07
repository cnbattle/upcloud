package cloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// Tencent 腾讯云 cos
type Tencent struct {
	SecretID  string
	SecretKey string
	VisitNode string
	Client    *cos.Client
}

func (t *Tencent) Init() error {
	u, err := url.Parse(t.VisitNode)
	if err != nil {
		return err
	}
	b := &cos.BaseURL{BucketURL: u}
	t.Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretID,
			SecretKey: t.SecretKey,
		},
	})
	return nil
}

func (t *Tencent) GetAll() (list []string, err error) {
	opt := &cos.BucketGetOptions{
		Prefix: "",
		//MaxKeys: 3,
	}
	v, _, err := t.Client.Bucket.Get(context.Background(), opt)
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
		_, err := t.Client.Object.Delete(context.Background(), item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tencent) Upload(localFile, upKey string) error {
	_, _, err := t.Client.Object.Upload(context.Background(), upKey, localFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tencent) Prefetch(urls []string) error {
	credential := common.NewCredential(t.SecretID, t.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)
	request := cdn.NewPurgeUrlsCacheRequest()
	request.Urls = common.StringPtrs(urls)
	_, err := client.PurgeUrlsCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return fmt.Errorf("An API error has returned: %s", err)
	}
	if err != nil {
		return err
	}
	request2 := cdn.NewPushUrlsCacheRequest()
	request2.Urls = common.StringPtrs(urls)
	_, err = client.PushUrlsCache(request2)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return fmt.Errorf("An API error has returned: %s", err)
	}
	if err != nil {
		return err
	}
	return nil
}
