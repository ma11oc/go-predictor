package cmd

import (
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string `yaml:"grpc_port" validate:"nonnil,nonzero,min=1,max=65535"`

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string `yaml:"http_port" validate:"nonnil,nonzero,min=1,max=65535"`

	// Path to locale
	Locale string `yaml:"locale" validate:"nonnil,nonzero"`

	// Log parameters section

	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int `yaml:"log_level" validate:"nonnil,min=-1,max=5"`

	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string `yaml:"log_time_format"`
}

// Validate validates configuration
func (cfg *Config) Validate() error {
	return validator.Validate(cfg)
}

// NewConfigFromMap return a new Config struct
func NewConfigFromMap(m map[string]interface{}) (*Config, error) {
	b, err := yaml.Marshal(m)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
