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
	"time"

	// "github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"

	// "github.com/golang/protobuf/ptypes/empty"
	// "google.golang.org/grpc"

	pb "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := time.Parse("2006-01-02", cmd.Flags().Lookup("birthday").Value.String())

		r, err := predictorClient.GetCardByBirthday(ctx, &pb.Date{
			Year:  uint32(d.Year()),
			Month: uint32(d.Month()),
			Day:   uint32(d.Day()),
		})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		fmt.Println(r)
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	cardCmd.Flags().StringP("birthday", "b", "", "Birthday")
}
