package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

func RegenerateOutlineUrl(url string) string {
	urls := strings.Split(url, "@")
	hashedCode := strings.TrimLeft(urls[0], "ss://")
	decodeStr, err := base64.StdEncoding.DecodeString(hashedCode)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("ss://%v@%v", string(decodeStr), urls[1])
}

func RegenerateUrl(url string) string {
	hashedCode := strings.TrimLeft(url, "ss://")
	logrus.Println(hashedCode)
	decodeStr, _ := base64.StdEncoding.DecodeString(hashedCode)
	return fmt.Sprintf("ss://%v", string(decodeStr))
}
