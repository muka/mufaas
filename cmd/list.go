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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available functions",
	Long: `List functions deployed and ready to be called`,
	Run: func(cmd *cobra.Command, args []string) {

		url := cmd.Flag("url").Value.String()

		c, conn, err := api.NewClient(url)
		if err != nil {
			fmt.Printf("Failed to connect to %s: %s", url, err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		var f []string

		ctx := context.Background()
		res, err := c.List(ctx, &api.ListRequest{Filter: f})
		if err != nil {
			fmt.Printf("Request failed: %s", err.Error())
			os.Exit(1)
		}

		if len(res.Functions) == 0 {
			fmt.Println("No function available, run `mufaas add -h` to learn how to add one")
			os.Exit(0)
		}

		fmt.Println("ID\t\tName")
		fmt.Println("----------------------------------------------")
		for _, f := range res.Functions {
			id := f.ID[strings.Index(f.ID, ":")+1:][:10]
			name := f.Name
			name = name[:strings.Index(name, ":")]
			fmt.Printf("%s\t%s\n", id, name)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
