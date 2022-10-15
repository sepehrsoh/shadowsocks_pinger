package cli

import (
	"github.com/sepehrsoh/shadowsocks_pinger/ping"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client pinger service.",
	Long:  `client pinger service.`,
	Run: func(cmd *cobra.Command, args []string) {
		client()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}

func client() {
	logrus.Infoln("starting curl")
	time.Sleep(time.Second * 5)
	for true {
		ping.PingUrl("youtube.com")
		time.Sleep(time.Second * 5)
		logrus.Infoln("try to connect")
	}
	return
}
