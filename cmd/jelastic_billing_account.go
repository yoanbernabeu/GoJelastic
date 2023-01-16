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

	rootCmd.AddCommand(getExtendedAccountBillingHistoryByPeriodCmd)
	getExtendedAccountBillingHistoryByPeriodCmd.Flags().String("appid", "", "An appid is required")
	getExtendedAccountBillingHistoryByPeriodCmd.MarkFlagRequired("appid")
	getExtendedAccountBillingHistoryByPeriodCmd.Flags().String("startTime", "", "A startTime is required")
	getExtendedAccountBillingHistoryByPeriodCmd.MarkFlagRequired("startTime")
	getExtendedAccountBillingHistoryByPeriodCmd.Flags().String("endTime", "", "A endTime is required")
	getExtendedAccountBillingHistoryByPeriodCmd.MarkFlagRequired("endTime")
}

var getAccountCmd = &cobra.Command{
	Use:     "getAccount",
	Short:   "Gets account by session",
	Long:    "Gets account by session",
	GroupID: "Billing/Account",
	Example: "GoJelastic billing/account getAccount --token=token --url=url --appid=appid",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")

		finalUrl := url + "/billing/account/rest/getaccount" + "?session=" + token + "&envName=" + appid

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}

var getExtendedAccountBillingHistoryByPeriodCmd = &cobra.Command{
	Use:     "getExtendedAccountBillingHistoryByPeriod",
	Short:   "Gets extended account billing history by period",
	Long:    "Gets extended account billing history by period",
	GroupID: "Billing/Account",
	Example: "GoJelastic billing/account getExtendedAccountBillingHistoryByPeriod --token=token --url=url --appid=appid --startTime=startTime --endTime=endTime",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")
		appid, _ := cmd.Flags().GetString("appid")
		starttime, _ := cmd.Flags().GetString("startTime")
		endtime, _ := cmd.Flags().GetString("endTime")

		finalUrl := url + "/billing/account/rest/getextendedaccountbillinghistorybyperiod" + "?session=" + token + "&appid=" + appid + "&startTime=" + starttime + "&endtime=" + endtime

		response := makeRequest(finalUrl, "GET", "")
		fmt.Println(response)
	},
}
