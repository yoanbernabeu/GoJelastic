package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getEnvsCmd)

	rootCmd.AddCommand(getEnvCmd)
	getEnvCmd.Flags().String("appid", "", "An appid is required")
	getEnvCmd.MarkFlagRequired("appid")

	rootCmd.AddCommand(startEnvCmd)
	startEnvCmd.Flags().String("appid", "", "An appid is required")
	startEnvCmd.MarkFlagRequired("appid")

	rootCmd.AddCommand(stopEnvCmd)
	stopEnvCmd.Flags().String("appid", "", "An appid is required")
	stopEnvCmd.MarkFlagRequired("appid")

	rootCmd.AddCommand(redeployContainerByIdCmd)
	redeployContainerByIdCmd.Flags().String("nodeid", "", "An nodeid is required")
	redeployContainerByIdCmd.Flags().String("tag", "", "An tag is required")
	redeployContainerByIdCmd.Flags().String("appid", "", "An appid is required")
	redeployContainerByIdCmd.MarkFlagRequired("nodeid")
	redeployContainerByIdCmd.MarkFlagRequired("tag")
	redeployContainerByIdCmd.MarkFlagRequired("appid")
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

var redeployContainerByIdCmd = &cobra.Command{
	Use:   "redeployContainerById",
	Short: "Redeploy a container by id",
	Long:  "Redeploy a container by id and specify Tag",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		nodeid, _ := cmd.Flags().GetString("nodeid")
		tag, _ := cmd.Flags().GetString("tag")
		appId, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/control/rest/redeploycontainerbyid?nodeId=" + nodeid + "&tag=" + tag + "&session=" + token + "&envName=" + appId

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
