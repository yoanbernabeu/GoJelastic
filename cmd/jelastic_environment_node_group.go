package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getNodeGroupsCmd)

	getNodeGroupsCmd.Flags().String("nodeid", "", "An nodeid is required")
	getNodeGroupsCmd.Flags().String("appid", "", "An appid is required")
	getNodeGroupsCmd.MarkFlagRequired("nodeid")
	getNodeGroupsCmd.MarkFlagRequired("appid")

	rootCmd.AddGroup(&cobra.Group{ID: "Environment/NodeGroup", Title: "Environment/NodeGroup"})

}

var getNodeGroupsCmd = &cobra.Command{
	Use:     "getNodeGroups",
	Short:   "Gets node group's data.",
	Long:    "Gets node group's data.",
	GroupID: "Environment/NodeGroup",
	Example: `GoJelastic getNodeGroups --token=token --url=url --appid=appid --nodeid=nodeid`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")
		nodeid, _ := cmd.Flags().GetString("nodeid")

		finalUrl := url + "/environment/nodegroup/rest/get" + "?session=" + token + "&envName=" + appid + "&nodeGroup=" + nodeid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
