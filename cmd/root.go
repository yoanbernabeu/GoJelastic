/*
Copyright © 2022 Yoan BERNABEU <yoan.bernabeu@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GoJelastic",
	Short: "An Alternative CLI for Jelastic",
	Long:  `An Alternative and not official CLI for Jelastic`,
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

	viper.ReadInConfig()

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

func makeHTTPRequest(url string, method string, body string) (*http.Response, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func formatResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	var data interface{}
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

//makeRequest makes a request to the Jelastic API
func makeRequest(url string, method string, body string) string {
	resp, err := makeHTTPRequest(url, method, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	formattedResponse, err := formatResponse(resp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return formattedResponse
}
