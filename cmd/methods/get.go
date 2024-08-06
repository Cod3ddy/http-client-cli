/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package methods

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Make a Get request",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		response  := getReq(url)
		fmt.Println(response)
	},
}

func init() {
	getCmd.Flags().StringVarP(&url, "url", "u", "", "Data to be fetched")

	if err := getCmd.MarkFlagRequired("url"); err != nil{
		fmt.Println(err)
		return
	}
}

func getReq (url string ) string {
	response, err := http.Get(url)

	if err != nil{
		return err.Error()
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	
	if err != nil{
		return err.Error()
	}
	return string (body)
}