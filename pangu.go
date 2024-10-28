package starter

import (
	"github.com/pangum/pangu"
	"github.com/pangum/todo/internal/core"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		core.NewWrapper,
		core.NewConfig,
		core.NewTodo,
	).Build().Apply()
}
