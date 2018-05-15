package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zhangmingkai4315/geoip-server/cache"
	"github.com/zhangmingkai4315/geoip-server/config"
)

var cfgFile string

// AppConfigInstance will store all config items in memory
var AppConfigInstance *config.AppConfig

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.toml", "config file for read")
	rootCmd.AddCommand(serverCmd)
}
func initConfig() {
	appConfig, err := config.NewAppConfig(cfgFile)
	if err != nil || appConfig == nil {
		log.Printf("Read config file error:%s\n", err)
		os.Exit(1)
	}

	_, err = cache.InitConnect(appConfig.DatabaseConfig.HostAndPort)
	if err != nil {
		log.Printf("Connect redis server fail : %s\n", err)
	}
	log.Printf("Connect redis server success\n")
}

var rootCmd = &cobra.Command{
	Use:   "geoip",
	Short: "Geoip is a ip to location query service",
	Long: `A Fast ip query service based on free geoip database and using redis cache for store ip data
                Complete documentation is available at https://github.com/zhangmingkai4315/geoip-server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute will start the cobra app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
