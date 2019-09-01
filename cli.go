package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	FileName string
	Port     string
	EndPoint string
)

func main() {

	mainApp := cli.NewApp()

	mainApp.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "serve a json file restfully",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "endpoint , e",
					Value:       "/mockrest",
					Usage:       "The api endpoint to be hit default /mockrest ie http://localhost:3000/mockrest",
					Destination: &EndPoint,
				},
				cli.StringFlag{
					Name:        "port , p",
					Value:       "3000",
					Usage:       "port to be served ",
					Destination: &Port,
				},
			},
			Action: func(c *cli.Context) error {
				StartServer(c.Args().First())
				return nil
			},
		},
	}
	mainApp.Action = func(c *cli.Context) {
		if _, err := os.Stat("mockrest.json"); err == nil && c.Args().First() == "" {
			StartServerWithConfig("mockrest.json")
		} else {
			if _, err := os.Stat(c.Args().First()); err == nil {
				StartServerWithConfig(c.Args().First())
			} else {
				fmt.Printf("Cannot find file %s \n", c.Args().First())
			}
		}
	}
	mainApp.Run(os.Args)
}
