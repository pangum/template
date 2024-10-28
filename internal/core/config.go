package core

import (
	"github.com/pangum/todo/internal/config"
	"github.com/pangum/todo/internal/plugin"
)

func NewConfig(wrapper *plugin.Wrapper) *config.Todo {
	return wrapper.Todo
}
