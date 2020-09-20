package cloud

import "github.com/cnbattle/upcloud/config"

// CommInterface 基础接口
type CommInterface interface {
	Init() error
	Setting() config.ProjectConfig
	GetAll() (list []string, err error)
	DelAll(list []string) error
	Upload(localFile, upKey string) error
	Prefetch() error
}
