// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.

package daemon

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/coreservice-io/utils/path_util"
)

// openWrtRecord - standard record (struct) for linux openWrtRecord version of daemon package
type openWrtRecord struct {
	name         string
	description  string
	kind         Kind
	dependencies []string
}

// Standard service path for systemV daemons
func (linux *openWrtRecord) servicePath() string {
	return "/etc/init.d/" + linux.name
}

// Is a service installed
func (linux *openWrtRecord) isInstalled() bool {

	if _, err := os.Stat(linux.servicePath()); err == nil {
		return true
	}

	return false
}

// Check service is running
func (linux *openWrtRecord) checkRunning() (string, bool) {
	output, err := exec.Command("service", linux.name, "status").Output()
	if err == nil {
		if matched, err := regexp.MatchString(linux.name, string(output)); err == nil && matched {
			reg := regexp.MustCompile("pid  ([0-9]+)")
			data := reg.FindStringSubmatch(string(output))
			if len(data) > 1 {
				return "Service (pid  " + data[1] + ") is running...", true
			}
			return "Service is running...", true
		}
	}

	return "Service is stopped", false
}

// Install the service
func (linux *openWrtRecord) Install(args ...string) (string, error) {
	installAction := "Install " + linux.description + ":"

	if ok, err := checkPrivileges(); !ok {
		return installAction + failed, err
	}

	srvPath := linux.servicePath()

	if linux.isInstalled() {
		return installAction + failed, ErrAlreadyInstalled
	}

	execPatch := path_util.ExE_Path(linux.name)
	_, err := os.Stat(execPatch)
	if err != nil {
		return installAction + failed, err
	}

	file, err := os.Create(srvPath)
	if err != nil {
		return installAction + failed, err
	}
	defer file.Close()

	templ, err := template.New("openWrtConfig").Parse(openWrtConfig)
	if err != nil {
		return installAction + failed, err
	}

	if err := templ.Execute(
		file,
		&struct {
			Name, Description, Path, Args string
		}{linux.name, linux.description, execPatch, strings.Join(args, " ")},
	); err != nil {
		return installAction + failed, err
	}

	if err := os.Chmod(srvPath, 0755); err != nil {
		return installAction + failed, err
	}

	return installAction + success, nil
}

// Remove the service
func (linux *openWrtRecord) Remove() (string, error) {
	removeAction := "Removing " + linux.description + ":"

	if ok, err := checkPrivileges(); !ok {
		return removeAction + failed, err
	}

	if !linux.isInstalled() {
		return removeAction + failed, ErrNotInstalled
	}

	if err := os.Remove(linux.servicePath()); err != nil {
		return removeAction + failed, err
	}

	return removeAction + success, nil
}

// Start the service
func (linux *openWrtRecord) Start() (string, error) {
	startAction := "Starting " + linux.description + ":"

	if ok, err := checkPrivileges(); !ok {
		return startAction + failed, err
	}

	if !linux.isInstalled() {
		return startAction + failed, ErrNotInstalled
	}

	if _, ok := linux.checkRunning(); ok {
		return startAction + failed, ErrAlreadyRunning
	}

	if err := exec.Command("service", linux.name, "start").Run(); err != nil {
		return startAction + failed, err
	}

	return startAction + success, nil
}

// Stop the service
func (linux *openWrtRecord) Stop() (string, error) {
	stopAction := "Stopping " + linux.description + ":"

	if ok, err := checkPrivileges(); !ok {
		return stopAction + failed, err
	}

	if !linux.isInstalled() {
		return stopAction + failed, ErrNotInstalled
	}

	if _, ok := linux.checkRunning(); !ok {
		return stopAction + failed, ErrAlreadyStopped
	}

	if err := exec.Command("service", linux.name, "stop").Run(); err != nil {
		return stopAction + failed, err
	}

	return stopAction + success, nil
}

// Status - Get service status
func (linux *openWrtRecord) Status() (string, error) {

	if ok, err := checkPrivileges(); !ok {
		return "", err
	}

	if !linux.isInstalled() {
		return statNotInstalled, ErrNotInstalled
	}

	statusAction, _ := linux.checkRunning()

	return statusAction, nil
}

// Run - Run service
func (linux *openWrtRecord) Run(e Executable) (string, error) {
	runAction := "Running " + linux.description + ":"
	e.Run()
	return runAction + " completed.", nil
}

// GetTemplate - gets service config template
func (linux *openWrtRecord) GetTemplate() string {
	return systemVConfig
}

// SetTemplate - sets service config template
func (linux *openWrtRecord) SetTemplate(tplStr string) error {
	systemVConfig = tplStr
	return nil
}

var openWrtConfig = `#!/bin/sh /etc/rc.common
#
#       /etc/init.d/{{.Name}}
#
#       Starts {{.Name}} as a daemon
#
# Copyright (C) 2008 OpenWrt.org
# description: Starts and stops a single {{.Name}} instance on this system

START=98
STOP=98

USE_PROCD=1

DAEMON={{.Name}}
PROG={{.Path}}

start_service() {
	echo "start user service!"

	# ubus call service list -check instance
	procd_open_instance
	
	#respawn
	# threshold：3600；timeout：5；retry：5
	procd_set_param respawn 0 60 100
	
	# run 
	procd_set_param command $PROG

	# pidfile
	procd_set_param pidfile /var/run/${DAEMON}.pid

	# 完成进程实例的增加
	procd_close_instance
}

stop_service() {
	echo "stop user service!"
	rm -f /var/run/${DAEMON}.pid
	service_stop "$PROG"
	killall $DAEMON
}

reload_service(){
	echo "reload user service!"
	stop
	start
}

# service_started(){

# }

restart() {
 　　stop
 　　start
}

`
