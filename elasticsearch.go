package main

import (
	elastigo "github.com/mattbaird/elastigo/lib"
	"log"
)

type ElasticSettings struct {
	Host string
	Port int
}

func Connection(settings ElasticSettings) *elastigo.Conn {
	c := elastigo.NewConn()
	c.Domain = settings.Host
	c.Port = string(settings.Port)
	return c
}

func DeleteIndex(index string, settings ElasticSettings) {
	c := Connection(settings)
	c.DeleteIndex(index)
}

func load(json string, settings ElasticSettings) {
	//log.Print(json)
	DeleteIndex("addrfeat", settings)
	c := elastigo.NewConn()
	c.Domain = settings.Host

	_, err := c.Index("census", "addrfeat", "", nil, json)
	if err != nil {
		log.Panic(err)
	}
}
