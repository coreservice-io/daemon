package main

import (
	"os"

	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/cmd"
)

func main() {
	basic.InitLogger()

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Panicln(errRun)
	}
}
