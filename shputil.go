package main

import (
	"log"
	"strings"

	"github.com/jonas-p/go-shp"
	"github.com/kpawlik/geojson"
)

func ReadShapefile(filename string) {
	getLines(filename)

}

func Features(filename string) (features []*geojson.Feature) {
	file, err := shp.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fields := file.Fields()
	props := make(map[string]interface{})

	for file.Next() {
		n, shape := file.Shape()
		line := shape.(*shp.PolyLine)
		geometry := getLineString(line)
		for k, f := range fields {
			name := strings.Trim(f.String(), "\u0000")
			value := file.ReadAttribute(n, k)
			props[name] = value
		}
		feature := geojson.NewFeature(geometry, props, nil)
		features = append(features, feature)
	}
	return features
}

func getLines(filename string) {
	file, err := shp.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fields := file.Fields()

	props := make(map[string]interface{})

	for file.Next() {
		n, shape := file.Shape()
		line := shape.(*shp.PolyLine)
		geometry := getLineString(line)
		for k, f := range fields {
			name := strings.Trim(f.String(), "\u0000")
			value := file.ReadAttribute(n, k)
			props[name] = value
		}
		feature := geojson.NewFeature(geometry, props, nil)
		json, err := geojson.Marshal(feature)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(json)

	}
}

func getLineString(line *shp.PolyLine) *geojson.LineString {
	points := line.Points
	coordinates := geojson.Coordinates{}
	for _, point := range points {
		c := geojson.Coordinate{geojson.Coord(point.X), geojson.Coord(point.Y)}
		coordinates = append(coordinates, c)
	}
	linestring := geojson.NewLineString(coordinates)
	return linestring
}
