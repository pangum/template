package core

import (
	"github.com/pangum/pangu"
	"github.com/pangum/todo/internal/plugin"
)

func NewWrapper(config *pangu.Config) (wrapper *plugin.Wrapper, err error) {
	wrapper = new(plugin.Wrapper)
	err = config.Build().Get(wrapper)

	return
}
