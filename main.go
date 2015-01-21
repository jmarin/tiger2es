package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tiger2es"
	//app.Usage = "tiger2es <state> <host> <port>"
	app.Author = "Juan Marin Otero"
	app.Email = "juan.marin.otero@gmail.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "state, s",
			Value: "all",
			Usage: "state to process (FIPS code)",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
			Usage: "Elasticsearch host",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "9200",
			Usage: "Elasticsearch HTTP port",
		},
	}
	app.Action = func(c *cli.Context) {
		state := c.String("state")
		host := c.String("host")
		port := c.String("port")
		fmt.Println("Loading state: " + state + " into " + host + ":" + port)
	}

	app.Run(os.Args)

}
