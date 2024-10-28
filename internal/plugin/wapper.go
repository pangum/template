package plugin

import (
	"github.com/pangum/todo/internal/config"
)

type Wrapper struct {
	Todo *config.Todo `json:"todo,omitempty" validate:"required"`
}
