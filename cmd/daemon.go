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
	"github.com/muka/mufaas/service"
	"github.com/spf13/cobra"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the daemon",
	Long: `Start the service to listen for request`,
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flag("listen")
		url := f.Value.String()
		fmt.Printf("Daemon listening to %s\n", url)
		err := service.Start(url)
		if err != nil {
			fmt.Printf("Daemon failed to start: %s", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.Flags().String("listen", ":5000", "URL to listen for")
}
