package daemon_util

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func Test_uname(t *testing.T) {
	log.Println(uname())
}

var openwrtNameArr = []string{
	"bobcatminer",
	"wrt",
}

func isOpenWrt() bool {
	osInfo, err := uname()
	if err == nil {
		log.Println("[warning] xxxxxxxxx")
		for _, v := range openwrtNameArr {
			if strings.Index(osInfo, v) != -1 { //exist
				return true
			}
		}
	}

	return false
}

func uname() (string, error) {
	cmd := exec.Command("uname", "-a")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil { // run command
		return "", err
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // read the output
		return "", err
	} else {
		return strings.ToLower(string(opBytes)), nil
	}
}
