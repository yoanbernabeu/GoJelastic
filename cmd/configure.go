package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configure)
	configure.Flags().String("token", "", "Your Jelastic token")
	configure.Flags().String("url", "", "Your Jelastic url")
	configure.MarkFlagRequired("token")
	configure.MarkFlagRequired("url")
}

var configure = &cobra.Command{
	Use:   "configure",
	Short: "Configure your Jelastic account",
	Long:  "Configure your Jelastic account",
	Run: func(cmd *cobra.Command, args []string) {
		writeConfig(cmd.Flags().Lookup("url").Value.String(), cmd.Flags().Lookup("token").Value.String())
	},
}
