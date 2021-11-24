/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/Wulfheart/req/requester"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Args:  cobra.ExactValidArgs(1),
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		path, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}

		configPath := filepath.Join(path, requester.ConfigDirectoryName)
		file, _ := os.Stat(configPath)

		for file == nil {
			path = filepath.Dir(path)
			if filepath.Dir(configPath) == path {
				panic("Folder not found")
			}
			configPath = filepath.Join(path, requester.ConfigDirectoryName)
			file, _ = os.Stat(configPath)

		}
		var session requester.SessionVariables
		requests := requester.BuildRequestsFromFile(args[0], requester.Config{
			ConfigPath:  configPath,
			HttpVersion: "HTTP/1.1",
		}, &session)

		for _, request := range requests {
			err = request.Prepare()
			if err != nil {
				panic(err)
			}
			err = request.DoRequest()
			if err != nil {
				panic(err)
			}
			fmt.Println(request.Response.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
