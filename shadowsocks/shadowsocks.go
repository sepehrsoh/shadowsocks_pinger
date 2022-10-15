package shadowsocks

import (
	"os/exec"
)

func RunShadowSocksServer(url string) error {
	for true {
		err := exec.Command("./run.sh", url).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
