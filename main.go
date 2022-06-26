package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "service-util",
		Usage: "run app as system service",
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "install app as system service",
			},
			{
				Name:  "remove",
				Usage: "remove app from system service",
			},
			{
				Name:  "start",
				Usage: "start app",
			},
			{
				Name:  "stop",
				Usage: "stop app",
			},
			{
				Name:  "restart",
				Usage: "restart app",
			},
			{
				Name:  "status",
				Usage: "show app status",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("removed task template: ", cCtx.Args().First())
					log.Println(cCtx.Args())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func RunServiceCmd(cCtx *cli.Context) error {
	fmt.Printf("Hello %q", cCtx.Args().Get(0))
	log.Println(cCtx.Args())
	return nil
}
