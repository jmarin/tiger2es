package main

import (
	"log"

	"github.com/kpawlik/geojson"
)

func ToGeoJson(f *geojson.Feature) string {
	json, err := geojson.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
