package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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
		cli.IntFlag{
			Name:  "port, p",
			Value: 9200,
			Usage: "Elasticsearch HTTP port",
		},
	}

	app.Action = func(c *cli.Context) {
		start := time.Now()
		state := c.String("state")
		host := c.String("host")
		port := c.Int("port")
		settings := ElasticSettings{Host: host, Port: port}
		if state != "" {
			zipFiles := DownloadAddrFeat(state)
			log.Print("Loading state: " + state + " into " + host + ":" + strconv.Itoa(port))
			for _, zipFile := range zipFiles {
				UnzipFile(zipFile)
				str := strings.Split(zipFile, ".")
				shp := str[0] + ".shp"
				features := Features(shp)
				for _, f := range features {
					load(ToGeoJson(f), settings)
				}
			}
		} else {
			log.Print("Please provide a state to process")
		}
		end := time.Now()
		execTime := end.Sub(start)
		log.Print(execTime)

	}

	app.Run(os.Args)

}
