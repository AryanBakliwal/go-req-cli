/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "GET request to a URL",
	Long:  `Makes a GET request to a provided URL.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverURL := args[0]

		resp, err := http.Get(serverURL)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		defer resp.Body.Close()

		dataBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(dataBytes))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	//getCmd.Flags().StringVarP(&message, "message", "m", "", "Body for GET request")
}
