package main

import (
	"github.com/codegangsta/cli"
	"github.com/enderian/directrd/agent"
	"github.com/enderian/directrd/server"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "start directrd as a server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Value: "config.yml",
					Usage: "set the configuration file",
				},
			},
			Action: server.Setup,
		},
		{
			Name:  "agent",
			Usage: "start directrd as an agent",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Value: "config.yml",
					Usage: "set the configuration file",
				},
			},
			Action: agent.Setup,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
