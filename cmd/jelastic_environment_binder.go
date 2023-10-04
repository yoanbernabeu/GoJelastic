package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(swapExtIpsCmd)
	swapExtIpsCmd.Flags().String("envName", "", "An envName is required")
	swapExtIpsCmd.MarkFlagRequired("envName")
	swapExtIpsCmd.Flags().String("sourceNodeId", "", "An sourceNodeId is required")
	swapExtIpsCmd.MarkFlagRequired("sourceNodeId")
	swapExtIpsCmd.Flags().String("targetNodeId", "", "An targetNodeId is required")
	swapExtIpsCmd.MarkFlagRequired("targetNodeId")

	rootCmd.AddGroup(&cobra.Group{ID: "Environment/Binder", Title: "Environment/Binder"})
}

var swapExtIpsCmd = &cobra.Command{
	Use:     "swapExtIps",
	Short:   "Swaps external IP addresses between source and target nodes",
	Long:    "Swaps external IP addresses between source and target nodes (just moves IP if the target node does not have one).",
	GroupID: "Environment/Binder",
	Example: `GoJelastic swapExtIps --token=token --url=url --envName=envName --sourceNodeId=sourceNodeId --targetNodeId=targetNodeId`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		envName, _ := cmd.Flags().GetString("envName")
		sourceNodeId, _ := cmd.Flags().GetString("sourceNodeId")
		targetNodeId, _ := cmd.Flags().GetString("targetNodeId")

		finalUrl := url + "/environment/binder/rest/swapextips" + "?session=" + token + "&envName=" + envName + "&sourceNodeId=" + sourceNodeId + "&targetNodeId=" + targetNodeId

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
