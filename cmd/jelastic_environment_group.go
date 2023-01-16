package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getGroupsCmd)
	getGroupsCmd.Flags().String("appid", "", "An appid is required")

	rootCmd.AddGroup(&cobra.Group{ID: "Environment/Group", Title: "Environment/Group"})

}

var getGroupsCmd = &cobra.Command{
	Use:     "getGroups",
	Short:   "Get group list",
	Long:    "Get group list",
	GroupID: "Environment/Group",
	Example: `GoJelastic getGroups --token=token --url=url --appid=appid`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/group/rest/getgroups" + "?session=" + token + "&envName=" + appid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
