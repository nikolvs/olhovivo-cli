package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/nikolvs/olhovivo-cli/olhovivo/cmd/predict"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "olhovivo",
	Short: "Olho Vivo is a public transport monitoring system for São Paulo",
	Long: `Olho Vivo is a public transport monitoring system for São Paulo.
This command-line interface is a tool to find bus information,
like arrival times, bus stops, lines and more.`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.olhovivo.yaml)")

	// Add subcommands
	rootCmd.AddCommand(predict.Command())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".olhovivo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".olhovivo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// TODO: Display this only in verbose mode
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	token := viper.GetString("token")
	if token == "" {
		fmt.Println("You MUST set the 'token' attribute in the config file.")
		os.Exit(1)
	}
}
