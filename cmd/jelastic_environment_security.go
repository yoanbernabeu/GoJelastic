package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getRulesCmd)
	getRulesCmd.Flags().String("appid", "", "An appid is required")
	getRulesCmd.MarkFlagRequired("appid")

	rootCmd.AddGroup(&cobra.Group{ID: "Environment/Security", Title: "Environment/Security"})
}

var getRulesCmd = &cobra.Command{
	Use:     "getRules",
	Short:   "Provides information about firewall rules for the environment",
	Long:    "Provides information about firewall rules for the environment",
	GroupID: "Environment/Security",
	Example: `GoJelastic getRules --token=token --url=url --appid=appid`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/environment/security/rest/getrules" + "?session=" + token + "&envName=" + appid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
