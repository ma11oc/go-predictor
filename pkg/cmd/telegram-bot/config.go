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
	ConnectionString string `yaml:"conn_str"    validate:"nonzero,nonnil,regexp=^(sqlite|postgres)://.*$"`
	AutoMigrate      bool   `yaml:"automigrate" validate:"regexp=^(true|false)$"`
}

// Logger specifies logging configuration
type Logger struct {
}

// PredictorServer specifies predictor server configuration
type PredictorServer struct {
	Endpoint string `yaml:"endpoint" validate:"nonzero,nonnil"`
}

// Tracer specifies configuration block for jaeger tracer
type Tracer struct {
	ServiceName string `yaml:"service_name" validate:"nonzero"`
	Endpoint    string `yaml:"endpoint"     validate:"nonzero"`
}

// Config describes the application configuration
type Config struct {
	Auth            *Auth            `yaml:"auth"      validate:"nonzero,nonnil"`
	Database        *Database        `yaml:"database"  validate:"nonzero,nonnil"`
	Logger          *Logger          `yaml:"logger"    validate:"nonzero,nonnil"`
	PredictorServer *PredictorServer `yaml:"predictor" validate:"nonzero,nonnil"`
	Tracer          *Tracer          `yaml:"tracer"    validate:"nonzero,nonnil"`
}

// Validate validates configuration
func (cfg *Config) Validate() error {
	return validator.Validate(cfg)
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var configGenCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates config file based on default values",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := yaml.Marshal(viper.AllSettings())
		fmt.Fprintln(os.Stdout, string(s))
	},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates config file specified",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := yaml.Marshal(viper.AllSettings()) // FIXME: handle error
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cfg := &Config{}

		if err := yaml.Unmarshal(b, cfg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := validator.Validate(cfg); err != nil {
			fmt.Fprintln(os.Stderr, "Configuration errors found:")
			fmt.Fprintln(os.Stderr, err)
		}
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
