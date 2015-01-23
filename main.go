package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "tiger2es"
	app.Author = "Juan Marin Otero"
	app.Email = "juan.marin.otero@gmail.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "state, s",
			//Value: "all",
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
		if state != "" {
			log.Print("Loading state: " + state + " into " + host + ":" + port)
			zipFiles := DownloadAddrFeat(state)
			for _, zipFile := range zipFiles {
				UnzipFile(zipFile)
				str := strings.Split(zipFile, ".")
				shp := str[0] + ".shp"
				features := Features(shp)
				for _, f := range features {
					log.Print(ToGeoJson(f))
				}
			}
		} else {
			log.Print("Please provide a state to process")
		}
	}

	app.Run(os.Args)

}
