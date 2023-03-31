package starter

import (
	"github.com/pangum/pangu"
)

// Starter 方便调用引用时
type Starter = TODO

func newAgent(config *pangu.Config) (starter *Starter, err error) {
	_config := new(panguConfig)
	if err = config.Load(_config); nil != err {
		return
	}

	return
}
