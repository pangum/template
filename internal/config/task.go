package config

type Task struct {
	Retries uint32 `json:"retries,omitempty" default:"10"`
}
