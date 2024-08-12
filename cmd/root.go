package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yahon",
	Short: "convertion tool",
	Long:  `convertion tool`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.yahon.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbosity")
	if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		log.Printf("Failed to bind flag: %v\n", err)
	}
	viper.SetDefault("verbose", false)

	rootCmd.PersistentFlags().StringP("source", "s", "", "Source file")
	viper.BindPFlag("source", rootCmd.PersistentFlags().Lookup("source"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in working directory, and then home directory with name ".cobra-app" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".yahon")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found! Counting on flags!")
	}

	verbose = viper.GetBool("verbose")
	if verbose {
		// If a config file is found, read it in.
		fmt.Fprintln(os.Stderr, "Using config:", viper.ConfigFileUsed())

		log.Println("--- Configuration ---")
		for s, i := range viper.AllSettings() {
			log.Printf("\t%s = %s\n", s, i)
		}
		log.Println("---")
	}
}
