package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().String("jps", "", "A JPS manifest is required")
	installCmd.Flags().String("envName", "", "An environment name is required")
	installCmd.Flags().String("envGroups", "", "An environment groups is optional")
	installCmd.MarkFlagRequired("jps")
	installCmd.MarkFlagRequired("envName")

	rootCmd.AddGroup(&cobra.Group{ID: "Marketplace/JPS", Title: "Marketplace/JPS"})
}

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Installs a custom JPS manifest.",
	Long:    "Installs a custom JPS manifest.",
	GroupID: "Marketplace/JPS",
	Example: `GoJelastic install --token=token --url=url --jps=jps --envName=envName --envGroups=envGroups`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		jps, _ := cmd.Flags().GetString("jps")
		envName, _ := cmd.Flags().GetString("envName")
		envGroups, _ := cmd.Flags().GetString("envGroups")

		finalUrl := url + "/marketplace/jps/rest/install" + "?session=" + token + "&envName=" + envName + "&jps=" + jps + "&envGroups=" + envGroups

		response := makeRequest(finalUrl, "POST", "")
		fmt.Println(response)
	},
}
