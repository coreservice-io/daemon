package service

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/basic/conf"
	"github.com/coreservice-io/service-util/plugin/daemon_plugin"
	"github.com/urfave/cli/v2"
)

func RunServiceCmd(clictx *cli.Context) {

	toml_conf := conf.Get_config().Toml_config

	daemon_name := toml_conf.Daemon_name
	if daemon_name == "" {
		basic.Logger.Errorln("daemon_name in config should not be vacant")
		return
	}

	exe_path, exe_path_err := os.Executable()
	if exe_path_err != nil {
		basic.Logger.Errorln(exe_path_err)
		return

	}

	exeDir := filepath.Dir(exe_path)

	if _, dir_err := os.Stat(path.Join(exeDir, "assets")); dir_err != nil {
		basic.Logger.Errorln("error -> please check:")
		basic.Logger.Errorln("1.dont directly `go run` for service, always `go build` first")
		basic.Logger.Errorln("2.the assets folder exist parellel to the excutable file ")
		return
	}

	basic.Logger.Infoln("exefile:" + exe_path + " to be service target")

	//check command
	subCmds := clictx.Command.Names()
	if len(subCmds) == 0 {
		basic.Logger.Fatalln("no sub command")
	}

	action := subCmds[0]
	err := daemon_plugin.Init(daemon_name)
	if err != nil {
		basic.Logger.Fatalln("init daemon service error:", err)
	}

	var status string
	var e error
	switch action {
	case "install":
		status, e = daemon_plugin.GetInstance(daemon_name).Install()
		basic.Logger.Debugln("cmd install")
	case "remove":
		daemon_plugin.GetInstance(daemon_name).Stop()
		status, e = daemon_plugin.GetInstance(daemon_name).Remove()
		basic.Logger.Debugln("cmd remove")
	case "start":
		status, e = daemon_plugin.GetInstance(daemon_name).Start()
		basic.Logger.Debugln("cmd start")
	case "stop":
		status, e = daemon_plugin.GetInstance(daemon_name).Stop()
		basic.Logger.Debugln("cmd stop")
	case "restart":
		daemon_plugin.GetInstance(daemon_name).Stop()
		status, e = daemon_plugin.GetInstance(daemon_name).Start()
		basic.Logger.Debugln("cmd restart")
	case "status":
		status, e = daemon_plugin.GetInstance(daemon_name).Status()
		basic.Logger.Debugln("cmd status")
	default:
		basic.Logger.Debugln("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}
