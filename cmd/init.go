package cmd

import (
	"fmt"
	"github.com/Wulfheart/req/requester"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var folder string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new req collection",
	Run: func(cmd *cobra.Command, args []string) {
		folderPath := filepath.Join(folder, requester.ConfigDirectoryName)
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		collectionPath := filepath.Join(folderPath, requester.CollectionVariablesFileName)
		_, err = os.Create(collectionPath)
		if err != nil {
			panic(err)
		}
		envPath := filepath.Join(folderPath, requester.EnvironmentVariablesFileName)
		_, err = os.Create(envPath)
		if err != nil {
			panic(err)
		}

		fmt.Println("Collection created at", folder)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringVarP(&folder, "folder", "f", "req", "Folder under which the collection is placed")
}
