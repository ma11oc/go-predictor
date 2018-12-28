// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	"context"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// "os"
	// "bitbucket.org/shchukin_a/go-predictor/internal/core"
	// "time"

	pb "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
	// "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	cfgFile         string
	predictorClient pb.PredictorClient
	conn            *grpc.ClientConn
	ctx             context.Context
	cancel          context.CancelFunc
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crystal-ball",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun:  predictorConnect,
	PersistentPostRun: predictorDisconnect,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd has been called")
		fmt.Printf("rootCmd args: %v\n", cmd.Flags())
		fmt.Printf("rootCmd args: %v\n", args)

		defer func() {
			fmt.Println("rootCmd.Run finished")
		}()

		// for _, v := range r.M.Values {
		// 	fmt.Printf("%v\n", v.GetNumberValue())
		// }

		// birthday, err := time.Parse("2006-01-02", cfg.birthday)

		// if err != nil {
		// 	log.Fatalf("Unable to parse birthday")
		// }

		// fmt.Println(birthday)

		// newDate := time.Date(2000, birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)

		// pc := core.GetCurrentPeriodicityCircles(newDate)

		// core.PrintAllPeriodicityCicles()

		// fmt.Println()

		// for i, v := range pc {
		// 	fmt.Printf(" %v: %v-%v | %v\n", i, v.Start.Format("02/01"), v.End.Format("02/01"), core.Planets[i])
		// }

		// fmt.Printf("your card is: %v\n", core.GetCardByBirthDate(&birthday))
		// fmt.Printf("originMatrix: %v", core.originMatrix)

	},
}

func predictorConnect(cmd *cobra.Command, args []string) {
	var err error
	// Set up a connection to the server.
	addr := cmd.Flags().Lookup("grpc-addr").Value.String()

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()

	predictorClient = pb.NewPredictorClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
}

func predictorDisconnect(cmd *cobra.Command, args []string) {
	defer cancel()
	defer conn.Close()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.crystal-ball.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "grpc-addr", "", "the address of predictor endpoint in format `host:port`")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.MarkFlagRequired("grpc-addr")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".crystal-ball" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".crystal-ball")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
