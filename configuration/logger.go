package configuration

type LoggerConfig struct {
	Level  string `yaml:"level"`
	LogDir string `yaml:"logDir"`
}
