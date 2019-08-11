package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/enderian/directrd/agent"
	"github.com/enderian/directrd/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app := cli.NewApp()

	app.Name = "directrd"
	app.Usage = "Manage your laboratory with style."
	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "Start directrd as a server",
			Action: server.Setup,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Value: "config.yml",
					Usage: "set the configuration file",
				},
			},
		},
		{
			Name:  "agent",
			Usage: "Start directrd as an agent",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "Starts the service",
					Action: agent.Start,
					Flags: []cli.Flag{
						cli.BoolTFlag{
							Name:  "deamon",
							Usage: "Run the agent as a deamon.",
						},
						cli.StringFlag{
							Name:  "config",
							Value: "config.yml",
							Usage: "set the configuration file",
						},
					},
				},
				{
					Name:   "stop",
					Usage:  "Stop the service",
					Action: agent.Stop,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
