package starter

import (
	"github.com/pangum/pangu"
)

// Starter 入口封装，由于语言特性，最好是直接继承的方式，通过显示使用主入口来执行初始化流程
type Starter struct {
	// TODO 从第三方要接入的主入口继承
}

func newAgent(config *pangu.Config) (starter *Starter, err error) {
	_config := new(panguConfig)
	if err = config.Load(_config); nil != err {
		return
	}

	return
}
