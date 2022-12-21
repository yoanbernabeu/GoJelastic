package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getEnvsCmd)
	getEnvsCmd.Flags().String("token", "", "A token is required")
	getEnvsCmd.Flags().String("url", "", "A url is required")
	getEnvsCmd.MarkFlagRequired("token")
	getEnvsCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(getEnvCmd)
	getEnvCmd.Flags().String("token", "", "A token is required")
	getEnvCmd.Flags().String("url", "", "A url is required")
	getEnvCmd.Flags().String("appid", "", "An appid is required")
	getEnvCmd.MarkFlagRequired("appid")
	getEnvCmd.MarkFlagRequired("token")
	getEnvCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(startEnvCmd)
	startEnvCmd.Flags().String("token", "", "A token is required")
	startEnvCmd.Flags().String("url", "", "A url is required")
	startEnvCmd.Flags().String("appid", "", "An appid is required")
	startEnvCmd.MarkFlagRequired("appid")
	startEnvCmd.MarkFlagRequired("token")
	startEnvCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(stopEnvCmd)
	stopEnvCmd.Flags().String("token", "", "A token is required")
	stopEnvCmd.Flags().String("url", "", "A url is required")
	stopEnvCmd.Flags().String("appid", "", "An appid is required")
	stopEnvCmd.MarkFlagRequired("appid")
	stopEnvCmd.MarkFlagRequired("token")
	stopEnvCmd.MarkFlagRequired("url")
}

var getEnvsCmd = &cobra.Command{
	Use:   "getEnvs",
	Short: "Get all your environments",
	Long:  "Get all your environments",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")

		finalUrl := url + "/environment/control/rest/getenvs" + "?session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var getEnvCmd = &cobra.Command{
	Use:   "getEnv",
	Short: "Get one environment",
	Long:  "Get one environment",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/control/rest/getenvinfo?envName=" + appid + "&session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var startEnvCmd = &cobra.Command{
	Use:   "startEnv",
	Short: "Start one environment",
	Long:  "Start one environment",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/control/rest/startenv?envName=" + appid + "&session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var stopEnvCmd = &cobra.Command{
	Use:   "stopEnv",
	Short: "Stop one environment",
	Long:  "Stop one environment",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/control/rest/stopenv?envName=" + appid + "&session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var makeRequest = func(url string, method string, body string) string {

	httpClient := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	return bodyString
}
