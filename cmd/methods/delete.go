/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package methods

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [url]",
	Short: "Make a DELETE request",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		reponse := deleteReq(url)

		fmt.Printf("status: %v", reponse)
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&url, "url", "u", "", "Data to be deleted")

	if err := deleteCmd.MarkFlagRequired("url"); err != nil{
		log.Fatal(err)
	}
}

func deleteReq(url string) string{
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil{
		log.Fatal(err)
	}

	client := &http.Client{}
	response, err:= client.Do(request)
	if err != nil{
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err !=nil{
		log.Fatal(err)
	}
	return string(body)
}
