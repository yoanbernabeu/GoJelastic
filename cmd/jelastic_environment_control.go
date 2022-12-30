package cmd

import (
	"fmt"

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

	rootCmd.AddCommand(getRegionsCmd)
	getRegionsCmd.Flags().String("appid", "", "An appid is required")
	getRegionsCmd.MarkFlagRequired("appid")

	rootCmd.AddCommand(getContainerEnvVarsCmd)
	getContainerEnvVarsCmd.Flags().String("appid", "", "An appid is required")
	getContainerEnvVarsCmd.MarkFlagRequired("appid")
	getContainerEnvVarsCmd.Flags().String("nodeid", "", "An nodeid is required")
	getContainerEnvVarsCmd.MarkFlagRequired("nodeid")

	rootCmd.AddGroup(&cobra.Group{ID: "Environment/Control", Title: "Environment/Control"})
}

var getEnvsCmd = &cobra.Command{
	Use:     "getEnvs",
	Short:   "Gets the information about all environments of a user",
	Long:    "Gets the information about all environments of a user",
	GroupID: "Environment/Control",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")

		finalUrl := url + "/environment/control/rest/getenvs" + "?session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var getEnvCmd = &cobra.Command{
	Use:     "getEnv",
	Short:   "Gets the full information about environment",
	Long:    "Gets the full information about environment (list of the nodes, settings etc.).",
	GroupID: "Environment/Control",
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
	Use:     "startEnv",
	Short:   "Start one environment",
	Long:    "Start one environment",
	GroupID: "Environment/Control",
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
	Use:     "stopEnv",
	Short:   "Stop one environment",
	Long:    "Stop one environment",
	GroupID: "Environment/Control",
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
	Use:     "redeployContainerById",
	Short:   "Redeploy a container by id",
	Long:    "Redeploy a container by id and specify Tag",
	GroupID: "Environment/Control",
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

var getRegionsCmd = &cobra.Command{
	Use:     "getRegions",
	Short:   "Gets available regions for the user",
	Long:    "Gets available regions for the user",
	GroupID: "Environment/Control",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/control/rest/getregions?envName=" + appid + "&session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var getContainerEnvVarsCmd = &cobra.Command{
	Use:     "getContainerEnvVars",
	Short:   "Gets env vars of container",
	Long:    "Gets env vars of container",
	GroupID: "Environment/Control",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")
		nodeid, _ := cmd.Flags().GetString("nodeid")

		finalUrl := url + "/environment/control/rest/getcontainerenvvars?envName=" + appid + "&session=" + token + "&nodeId=" + nodeid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
