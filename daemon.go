package main

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/coreservice-io/daemon/daemon_util"
	"github.com/urfave/cli/v2"
)

func NewService(serviceName string) (daemon_util.Daemon, error) {
	kind := daemon_util.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon_util.UserAgent
	}
	return daemon_util.New(serviceName, serviceName, kind)
}

func install(cCtx *cli.Context) error {
	serviceName := cCtx.Args().First()
	if serviceName == "" {
		return errors.New("service name error")
	}

	service, err := NewService(serviceName)
	if err != nil {
		return err
	}

	args := cCtx.Args().Slice()
	result, err := service.Install(args[1:]...)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func remove(cCtx *cli.Context) error {
	serviceName := cCtx.Args().First()
	if serviceName == "" {
		return errors.New("service name error")
	}

	service, err := NewService(serviceName)
	if err != nil {
		return err
	}

	service.Stop()
	result, err := service.Remove()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func start(cCtx *cli.Context) error {
	serviceName := cCtx.Args().First()
	if serviceName == "" {
		return errors.New("service name error")
	}

	service, err := NewService(serviceName)
	if err != nil {
		return err
	}
	result, err := service.Start()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func stop(cCtx *cli.Context) error {
	serviceName := cCtx.Args().First()
	if serviceName == "" {
		return errors.New("service name error")
	}

	service, err := NewService(serviceName)
	if err != nil {
		return err
	}
	result, err := service.Stop()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func restart(cCtx *cli.Context) error {
	serviceName := cCtx.Args().First()
	if serviceName == "" {
		return errors.New("service name error")
	}

	service, err := NewService(serviceName)
	if err != nil {
		return err
	}
	service.Stop()
	result, err := service.Start()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func status(cCtx *cli.Context) error {
	serviceName := cCtx.Args().First()
	if serviceName == "" {
		return errors.New("service name error")
	}

	service, err := NewService(serviceName)
	if err != nil {
		return err
	}
	result, err := service.Status()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}
