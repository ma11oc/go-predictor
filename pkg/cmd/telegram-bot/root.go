package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ProgramName is used in help messages and jaeger service name by default
const ProgramName = "predictor-tbot"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   ProgramName,
	Short: "A brief description of your application",
	Long:  "",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("db: %v\n", viper.Get("database"))
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-predictor.yaml)")

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
		viper.SetConfigName(".go-predictor")
	}

	// viper.SetEnvPrefix("pbot")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault("database.auto_migrate", true)
	viper.SetDefault("database.connection_string", fmt.Sprintf("sqlite://%s.db", ProgramName))

	viper.SetDefault("tracer.service_name", ProgramName)
	viper.SetDefault("tracer.endpoint", "http://localhost:14268/api/traces")

	viper.SetDefault("predictor.endpoint", "localhost:50051")

	viper.SetDefault("auth.token", "<AUTH_TOKEN_PLACEHOLDER>")

	// fmt.Println("root init done")
}
