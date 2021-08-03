package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ma11oc/go-predictor/pkg/logger"
	"github.com/ma11oc/go-predictor/pkg/protocol/grpc"
	"github.com/ma11oc/go-predictor/pkg/protocol/rest"
	v1 "github.com/ma11oc/go-predictor/pkg/service/v1"
)

// ProgramName is used in help messages and jaeger service name by default
const ProgramName = "predictor-srv"

var cfgFile string
var cfg Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   ProgramName,
	Short: "A brief description of your application",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var runtimeConfig *Config
		var err error

		if runtimeConfig, err = NewConfigFromMap(viper.AllSettings()); err != nil {
			return err
		}

		if err = runtimeConfig.Validate(); err != nil {
			return err
		}

		return RunServer(runtimeConfig)
	},
	Version: "0.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$HOME/."+ProgramName+".yaml", "config file")

	rootCmd.PersistentFlags().StringVar(&cfg.GRPCPort, "grpc-port", "50051", "grpc port to listen on")
	viper.BindPFlag("grpc_port", rootCmd.PersistentFlags().Lookup("grpc-port"))

	rootCmd.PersistentFlags().StringVar(&cfg.HTTPPort, "http-port", "8080", "http port to listen on")
	viper.BindPFlag("http_port", rootCmd.PersistentFlags().Lookup("http-port"))

	rootCmd.PersistentFlags().StringVar(&cfg.Locale, "locale", "", "path to locale")
	viper.BindPFlag("locale", rootCmd.PersistentFlags().Lookup("locale"))

	rootCmd.PersistentFlags().IntVar(&cfg.LogLevel, "log-level", 0, "log level (debug=-1, fatal=5)")
	viper.BindPFlag("log_level", rootCmd.PersistentFlags().Lookup("log-level"))

	// fmt.Println("cobra.init done")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-predictor" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("." + ProgramName)
	}

	// viper.SetEnvPrefix("pbot")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault("grpc_port", 50051)
	viper.SetDefault("http_port", 8080)
	viper.SetDefault("log_level", 8080)

	// fmt.Println("root init done")
}

// RunServer runs gRPC server and HTTP gateway
func RunServer(cfg *Config) error {
	ctx := context.Background()

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	v1API := v1.NewPredictorServiceServer(cfg.Locale)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
