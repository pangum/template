package core

import (
	"github.com/pangum/task/internal/config"
	"github.com/pangum/task/internal/plugin"
)

func NewConfig(wrapper *plugin.Wrapper) *config.Task {
	return wrapper.Task
}
