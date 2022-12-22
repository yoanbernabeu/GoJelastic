package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getAccountCmd)
	getAccountCmd.Flags().String("appid", "", "An appid is required")
	getAccountCmd.MarkFlagRequired("appid")

	rootCmd.AddGroup(&cobra.Group{ID: "Billing/Account", Title: "Billing/Account"})

}

var getAccountCmd = &cobra.Command{
	Use:     "getAccount",
	Short:   "Gets account by session.",
	Long:    "Gets account by session.",
	GroupID: "Billing/Account",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/billing/account/rest/getaccount" + "?session=" + token + "&envName=" + appid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
