package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"

	"github.com/jonas-p/go-shp"
)

func UnzipFile(zipFile string) {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		f, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = io.Copy(f, rc)
		if err != nil {
			log.Fatal(err)
		}

	}
}

// Pass .shp file name
func readShapefile(file string) {
	shape, err := shp.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	//fields
	fields := shape.Fields()

	//loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()

		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		//print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}

		fmt.Println()
	}
}
