package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhangmingkai4315/geoip-server/config"
	"github.com/zhangmingkai4315/geoip-server/web"
)

func init() {
	serverCmd.PersistentFlags().String("http.ListenAt", "", "http api server url")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a web server for api query",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetAppConfig()
		httpServerURL := cmd.Flag("http.ListenAt").Value.String()
		if httpServerURL == "" {
			httpServerURL = cfg.GlobalConfig.ListenAt
		}
		log.Printf("Try starting web serve at %s\n", httpServerURL)
		web.Start(httpServerURL)
	},
}
