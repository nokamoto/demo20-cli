package cmd

import (
	"log"
	"os"

	"github.com/nokamoto/demo20-cli/internal/cmd/config"
	defaultconfig "github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "cloud",
	Short: "A cloud management tool",
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(config.RootCmd)

	rootCmd.SetOut(os.Stdout)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cloud.yaml)")

	rootCmd.PersistentFlags().String("grpc-address", "localhost:9000", "gRPC server address")
	viper.BindPFlag("grpcaddress", rootCmd.PersistentFlags().Lookup("grpc-address"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(defaultconfig.Filename)
		viper.SetConfigType(defaultconfig.Extension)
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&defaultconfig.Default); err != nil {
		log.Fatal(err)
	}
}
