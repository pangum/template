package starter

import (
	"github.com/pangum/pangu"
)

// Starter 方便调用引用时
type Starter = TODO

func newAgent(config *pangu.Config) (starter *Starter, err error) {
	wrap := new(wrapper)
	if err = config.Load(wrap); nil != err {
		return
	}

	return
}
