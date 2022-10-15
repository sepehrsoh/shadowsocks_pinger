package ping

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
	"os/exec"
)

var ErrPacketLost = errors.New("packet lost")

func PingUrl(url string) {
	cmd := exec.Command("curl", "-v", "--proxy", "socks5h://127.0.0.1:1080", url)
	logrus.Println(cmd.String())
	go func() {
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			logrus.Errorln(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
		logrus.Println("done ping")
	}()
}

func PingIp(ip string) error {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	if err = pinger.Run(); err != nil {
		return err
	}
	status := pinger.Statistics()
	if status.PacketLoss > 0 {
		return ErrPacketLost
	}
	return nil
}
