/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wiphish",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var iface, esid, bsid, network, evilginx, gateway, pool string

var channel uint

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wiphish.yaml)")
	rootCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&iface, "interface", "w", "wlan1", "Wifi interface default wlan1")
	rootCmd.PersistentFlags().StringVarP(&esid, "esid", "e", "", "ESSID name")
	rootCmd.PersistentFlags().StringVarP(&bsid, "bsid", "b", "", "BSSID name")
	rootCmd.PersistentFlags().StringVarP(&network, "network", "n", "10.168.111.0/24", "Network")
	rootCmd.PersistentFlags().StringVarP(&evilginx, "gateway", "g", "10.168.111.2-10.168.111.254", "AP IP address")
	rootCmd.PersistentFlags().StringVarP(&gateway, "pool", "p", "10.168.111.1", "DHCP IP address pool")
	rootCmd.PersistentFlags().StringVarP(&evilginx, "evilginx", "d", "10.168.111.1", "Evilginx IP")
	rootCmd.PersistentFlags().UintVarP(&channel, "channel", "c", 11, "WiFi channel")
}
