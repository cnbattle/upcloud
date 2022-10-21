// Package cloud 支持的云平台
package cloud

// CommInterface 基础接口
type CommInterface interface {
	Init() error
	GetAll() (list []string, err error)
	DelAll(list []string) error
	Upload(localFile, upKey string) error
	Prefetch([]string) error
}
