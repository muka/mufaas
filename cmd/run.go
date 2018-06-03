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
	"time"
	"os"
	"context"
	"github.com/muka/mufaas/api"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var runRequest api.RunRequest

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a function",
	Long: `Run a function providing arguments and other context information`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Provide the function name to call")
			os.Exit(1)
		}

		runRequest.Name = args[0]
		runRequest.Args = args[1:]

		envs := os.Environ()
		for _, e := range envs {
			runRequest.Env = append(runRequest.Env, e)
		}

		url := cmd.Flag("url").Value.String()

		c, conn, err := api.NewClient(url)
		if err != nil {
			fmt.Printf("Failed to connect to %s: %s\n", url, err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		ctx := context.Background()
		t1 := time.Now()
		res, err := c.Run(ctx, &runRequest)
		if err != nil {
			fmt.Printf("Request failed: %s\n", err.Error())
			os.Exit(1)
		}

		t2 := time.Since(t1)
		log.Debugf("Run took %dms", t2.Nanoseconds()/1000000)

		if len(res.Err) > 0 {
			fmt.Printf("Error: %s\n", string(res.Err))
			os.Exit(1)
		}

		fmt.Print(string(res.Output))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
