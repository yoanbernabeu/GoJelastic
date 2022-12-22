/*
Copyright © 2022 Yoan BERNABEU <yoan.bernabeu@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GoJelastic",
	Short: "A Alternative CLI for Jelastic",
	Long:  `A Alternative and not official CLI for Jelastic`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config := initConfig()

	if config["token"] != nil {
		rootCmd.PersistentFlags().String("token", config["token"].(string), "A token is required")
	} else {
		rootCmd.PersistentFlags().String("token", "", "A token is required")
		rootCmd.MarkPersistentFlagRequired("token")
	}

	if config["url"] != nil {
		rootCmd.PersistentFlags().String("url", config["url"].(string), "A url is required")
	} else {
		rootCmd.PersistentFlags().String("url", "", "A url is required")
		rootCmd.MarkPersistentFlagRequired("url")
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() map[string]interface{} {
	// Don't forget to read config either from cfgFile or from home directory!
	viper.SetConfigType("env")
	viper.SetConfigName("gojelastic.env") // name of config file (without extension)
	viper.AddConfigPath("$HOME")          // adding home directory as first search path
	viper.AutomaticEnv()                  // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config file found")
	}

	return viper.AllSettings()
}

//writeConfig writes config file and ENV variables if set.
func writeConfig(url string, token string) {
	viper.SetConfigType("env")
	viper.SetConfigName("gojelastic") // name of config file (without extension)
	viper.AddConfigPath("$HOME")      // adding home directory as first search path
	viper.AutomaticEnv()              // read in environment variables that match

	viper.Set("url", url)
	viper.Set("token", token)

	viper.WriteConfig()

	fmt.Println("Config file written or updated")
}
