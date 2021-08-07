package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

// Auth specifies data required to connect to telegram api
type Auth struct {
	Token string `yaml:"token" validate:"nonzero,nonnil"`
}

// Database specifies a piece of configuration required to connect to a database
type Database struct {
	ConnectionString string `yaml:"connection_string" validate:"nonzero,nonnil"`
	AutoMigrate      bool   `yaml:"auto_migrate"      validate:""`
}

// Logger specifies logging configuration
type Logger struct {
	Level      int    `yaml:"level"       validate:"min=-1,max=5"`
	TimeFormat string `yaml:"time_format" validate:""`
}

// PredictorServer specifies predictor server configuration
type PredictorServer struct {
	Endpoint string `yaml:"endpoint" validate:"nonzero"`
}

// Tracer specifies configuration block for jaeger tracer
type Tracer struct {
	ServiceName string `yaml:"service_name" validate:"nonzero"`
	Endpoint    string `yaml:"endpoint"     validate:"nonzero"`
}

// Config describes the application configuration
type Config struct {
	Auth            *Auth            `yaml:"auth"      validate:"nonzero,nonnil"`
	Database        *Database        `yaml:"database"  validate:"nonzero"`
	Logger          *Logger          `yaml:"logger"    validate:"nonzero"`
	PredictorServer *PredictorServer `yaml:"predictor" validate:"nonzero"`
	Tracer          *Tracer          `yaml:"tracer"    validate:"nonzero"`
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

// NewConfig returns a config with default values set
func NewConfig() *Config {
	return &Config{
		Auth: &Auth{
			Token: "",
		},
		Database: &Database{
			ConnectionString: fmt.Sprintf("sqlite://%s.db", ProgramName),
			AutoMigrate:      true,
		},
		Tracer: &Tracer{
			ServiceName: ProgramName,
			Endpoint:    "http://localhost:14268/api/traces",
		},
		PredictorServer: &PredictorServer{
			Endpoint: "localhost:50051",
		},
		Logger: &Logger{
			Level:      0,
			TimeFormat: "",
		},
	}
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf"},
	Short:   "A brief description of your command",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var configGenCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generates config file based on default values",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := yaml.Marshal(viper.AllSettings())
		fmt.Fprintln(os.Stdout, string(s))
	},
}

var configValidateCmd = &cobra.Command{
	Use:     "validate",
	Aliases: []string{"check"},
	Short:   "Validates config file specified",
	Long:    "",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := NewConfigFromMap(viper.AllSettings())
		if err != nil {
			return err
		}

		if errs := validator.Validate(cfg); errs != nil {
			return errs
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(configGenCmd)
	configCmd.AddCommand(configValidateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
