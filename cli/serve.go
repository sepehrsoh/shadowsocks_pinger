package cli

import (
	"github.com/joho/godotenv"
	"github.com/sepehrsoh/shadowsocks_pinger/shadowsocks"
	"github.com/sepehrsoh/shadowsocks_pinger/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving pinger service.",
	Long:  `Serving pinger service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

var isOutline = true

func serve() {
	var fullUrl string
	logrus.Infoln("starting server")
	url := getUrlFromEnv("URL")
	if isOutline {
		fullUrl = utils.RegenerateOutlineUrl(url)
	} else {
		fullUrl = utils.RegenerateUrl(url)
	}
	logrus.Infoln("url:", fullUrl)
	err := shadowsocks.RunShadowSocksServer(fullUrl)
	if err != nil {
		logrus.Panic(err)
	}
	return
}
func getUrlFromEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Panic("Error loading .env file")
	}

	return os.Getenv(key)
}
