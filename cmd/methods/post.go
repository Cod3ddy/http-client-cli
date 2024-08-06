/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package methods

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Cod3ddy/http-client-cli/cmd/common"
	"github.com/spf13/cobra"
)

var (
	url string
	data string
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Make a post request",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string){
		response := postReq(url, data); 
		fmt.Println(response)
	},
}

func init() {
	postCmd.Flags().StringVarP(&data, "data", "d", "", "Data to send in the request body")

	if err := postCmd.MarkFlagRequired("data"); err != nil{
		fmt.Println(err)
		return
	}


	postCmd.Flags().StringVarP(&url, "url", "u","", "the url to send request to")

	if err := postCmd.MarkFlagRequired("url"); err != nil{
		fmt.Println(err)
		return
	}

}

func postReq(urlReg, dataInc string) string {
	inData := dataInc
	inUrl := urlReg


	dataIn, err := common.CheckIfValid(inData)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	dataReader := strings.NewReader(dataIn)
	response, err := http.Post(inUrl, "application/json", dataReader)

	if err != nil{
		fmt.Println(err)
		return "error"
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
	}

	return string(body)
}
