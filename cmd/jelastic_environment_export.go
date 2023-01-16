package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getListCmd)
	getListCmd.Flags().String("appid", "", "An appid is required")
	getListCmd.MarkFlagRequired("appid")

	rootCmd.AddGroup(&cobra.Group{ID: "Environment/Export", Title: "Environment/Export"})

}

var getListCmd = &cobra.Command{
	Use:     "getList",
	Short:   "Get list of an environment",
	Long:    "Get list of an environment",
	GroupID: "Environment/Export",
	Example: `GoJelastic getList --token=token --url=url --appid=appid`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/export/rest/getlist" + "?session=" + token + "&envName=" + appid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
