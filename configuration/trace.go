package configuration

// trace config
type TraceConfig struct {
	Agent   string `yaml:"agent"`
	FromEnv bool   `yaml:"fromEnv"`
}
