package plugin

import (
	"github.com/pangum/task/internal/config"
)

type Wrapper struct {
	Task *config.Task `json:"task,omitempty" validate:"required"`
}
