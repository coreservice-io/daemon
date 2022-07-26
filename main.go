package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "daemon-util",
		Usage: "run app as system service",
		Commands: []*cli.Command{
			{
				Name:   "install",
				Usage:  "install app as system service",
				Action: install,
			},
			{
				Name:   "remove",
				Usage:  "remove app from system service",
				Action: remove,
			},
			{
				Name:   "start",
				Usage:  "start app",
				Action: start,
			},
			{
				Name:   "stop",
				Usage:  "stop app",
				Action: stop,
			},
			{
				Name:   "restart",
				Usage:  "restart app",
				Action: restart,
			},
			{
				Name:   "status",
				Usage:  "show app status",
				Action: status,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
