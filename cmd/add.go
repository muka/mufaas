// Copyright © 2017 Luca Capra
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

package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"context"
	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/util"
	"github.com/spf13/cobra"
)

var addRequest api.AddRequest

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an image to be used as a function",
	Long: `Create a new image with the provided parameters`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Provide a path to the source code")
			os.Exit(1)
		}

		dir := args[0]

		err := util.CreateTar(dir)
		if err != nil {
			fmt.Printf("Error creating archive: %s", err.Error())
			os.Exit(1)
		}

		src, err := ioutil.ReadFile(dir+".tar")
		if err != nil {
			fmt.Printf("Error reading archive: %s", err.Error())
			os.Exit(1)
		}
		addRequest.Source = src

		url := cmd.Flag("url").Value.String()
		c, conn, err := api.NewClient(url)
		if err != nil {
			fmt.Printf("Failed to connect to %s: %s\n", url, err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		ctx := context.Background()
		res, err := c.Add(ctx, &addRequest)
		if err != nil {
			fmt.Printf("Request failed: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Added image %s (id=%s)", res.Info.Name, res.Info.ID[7:17])

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addRequest.Info = &api.FunctionInfo{}
	addCmd.Flags().StringVarP(&addRequest.Info.Name, "name", "n", "", "The image name to add")
	addCmd.Flags().StringVarP(&addRequest.Info.Type, "type", "t", "", "The language of the source code")
	// addCmd.Flags().StringVarP(&addRequest.Info.Cmd, "command", "c", "", "The command to run from the image (will override the default one)")
}
