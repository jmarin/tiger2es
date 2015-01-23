package main

import "log"

type ElasticSettings struct {
	Host string
	Port int
}

func load(json string, settings ElasticSettings) {
	log.Print(json)
}
