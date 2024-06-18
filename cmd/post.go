/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var mflag string

type Msg struct {
	MsgData string `json:"message"`
}

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "POST request to a URL",
	Long: `Makes a POST request to a provided URL.
Use the --message or -m flag to specify the message.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverURL := args[0]
		//marshal
		m := Msg{
			MsgData: mflag,
		}
		jsondata, _ := json.Marshal(m)

		//request server
		client := &http.Client{}
		req, _ := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsondata))
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		//read response and unmarshal
		body, _ := io.ReadAll(resp.Body)
		var mresp Msg
		err2 := json.Unmarshal([]byte(body), &mresp)
		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println("Response from server: ", mresp.MsgData)
		//fmt.Println("post called")
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	postCmd.Flags().StringVarP(&mflag, "message", "m", "", "Body for POST request")
}
