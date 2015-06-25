package main

import (
	"flag"
	"github.com/pravj/geopattern"
	"log"
)
import "path"
import "os"

func main() {
	wd, _ := os.Getwd()

	phrase := flag.String("phrase", "no phrase", "phrase to generate the pattern from")
	color := flag.String("color", "#0000EE", "base color of the pattern")
	filepath := flag.String("file", path.Join(wd, "result.svg"), "filepath of the result")
	generator := flag.String("generator", "xes", "generator name")
	flag.Parse()

	strings := map[string]string{
		"phrase":    *phrase,
		"generator": *generator,
		"color":     *color,
	}
	result := geopattern.Generate(strings)
	file, _ := os.Create(*filepath)
	_, err := file.Write([]byte(result))
	if err != nil {
		log.Fatal(err)
	}
}
