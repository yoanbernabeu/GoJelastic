package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	rootCmd.AddCommand(configure)
	rootCmd.AddCommand(documentation)
	configure.Flags().String("token", "", "Your Jelastic token")
	configure.Flags().String("url", "", "Your Jelastic url")
	configure.MarkFlagRequired("token")
	configure.MarkFlagRequired("url")
	rootCmd.AddGroup(&cobra.Group{ID: "configure", Title: "Configure CLI"})
}

var configure = &cobra.Command{
	Use:     "configure",
	Short:   "Configure your CLI",
	Long:    "Configure your CLI",
	Example: `jelastic configure --token=your_token --url=your_url`,
	GroupID: "configure",
	Run: func(cmd *cobra.Command, args []string) {
		writeConfig(cmd.Flags().Lookup("url").Value.String(), cmd.Flags().Lookup("token").Value.String())
	},
}

var documentation = &cobra.Command{
	Use:     "documentation",
	Short:   "Generate documentation",
	Long:    "Generate documentation",
	Example: `jelastic documentation`,
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, "./docs/GoJelastic")
		if err != nil {
			log.Fatal(err)
		}
	},
}
