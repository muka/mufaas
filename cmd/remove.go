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
	"strings"
	"golang.org/x/net/context"
	"os"
	"github.com/muka/mufaas/api"
	"github.com/spf13/cobra"
)

var forceRemove *bool

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		url := cmd.Flag("url").Value.String()

		c, conn, err := api.NewClient(url)
		if err != nil {
			fmt.Printf("Failed to connect to %s: %s", url, err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		ctx := context.Background()
		res, err := c.Remove(ctx, &api.RemoveRequest{Name: args, Force: *forceRemove})
		if err != nil {
			fmt.Printf("Request failed: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("ID\t\tName\t\tStatus")
		fmt.Println("----------------------------------------------")
		for _, f := range res.Functions {
			id := f.ID[strings.Index(f.ID, ":")+1:][:10]
			name := f.Name
			status := "removed"
			if len(f.Error) != 0 {
				status = f.Error
			}
			fmt.Printf("%s\t%s\t%s\n", id, name, status)
		}

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	forceRemove = removeCmd.Flags().BoolP("force", "f", false, "Force function remove")
}
