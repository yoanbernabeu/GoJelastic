package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getUserInfoCmd)
	rootCmd.AddCommand(recoverPasswordCmd)
	recoverPasswordCmd.Flags().String("email", "", "An email is required")
	recoverPasswordCmd.MarkFlagRequired("email")

	rootCmd.AddGroup(&cobra.Group{ID: "Users/Account", Title: "Users/Account"})
}

var getUserInfoCmd = &cobra.Command{
	Use:     "getUserInfo",
	Short:   "Gets information about the user.",
	Long:    "Gets information about the user.",
	GroupID: "Users/Account",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")

		finalUrl := url + "/users/account/rest/getuserinfo" + "?session=" + token

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var recoverPasswordCmd = &cobra.Command{
	Use:     "recoverPassword",
	Short:   "Sends an email with the link to reset the account password",
	Long:    "Sends an email with the link to reset the account password",
	GroupID: "Users/Account",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		email, _ := cmd.Flags().GetString("email")

		finalUrl := url + "/users/account/rest/recoverpassword" + "?email=" + email

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
