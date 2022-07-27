// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.

// Package daemon linux version
package daemon_util

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var openwrtNameArr = []string{
	"wrt",
}

var bobcatNameArr = []string{
	"bobcat",
}

// Get the daemon properly
func newDaemon(name, description string, kind Kind, dependencies []string) (Daemon, error) {
	// newer subsystem must be checked first
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return &systemDRecord{name, description, kind, dependencies}, nil
	}
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return &upstartRecord{name, description, kind, dependencies}, nil
	}

	if isOpenWrt() {
		log.Println("[info] openwrt detected")
		return &openWrtRecord{name, description, kind, dependencies}, nil
	}

	if isBobCat() {
		log.Println("[info] bobcat detected")
		return &bobCatRecord{name, description, kind, dependencies}, nil
	}

	log.Println("[warning] using default systemV type")
	return &systemVRecord{name, description, kind, dependencies}, nil
}

// Get executable path
func execPath() (string, error) {
	return os.Readlink("/proc/self/exe")
}

func isBobCat() bool {
	osInfo, err := uname()
	if err == nil {
		for _, v := range bobcatNameArr {
			if strings.Index(osInfo, v) != -1 { //exist
				return true
			}
		}
	}
	return false
}

func isOpenWrt() bool {
	osInfo, err := uname()
	if err == nil {
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

	if err := cmd.Start(); err != nil { // 运行命令
		return "", err
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		return "", err
	} else {
		return strings.ToLower(string(opBytes)), nil
	}
}
