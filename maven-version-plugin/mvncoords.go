package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Coordinates struct {
	Group    string `xml:"groupId"`
	Artifact string `xml:"artifactId"`
	Version  string `xml:"version"`
}

func main() {

	fileName := flag.String("f", "pom.xml", "filename to read")
	showGroup := flag.Bool("group", false, "output group coordinate")
	showApp := flag.Bool("app", false, "output application coordinate")
	showVersion := flag.Bool("version", true, "output version coordinate")
	flag.Parse()

	coords := readCoords(*fileName)

	var result strings.Builder
	if *showGroup {
		result.WriteString(coords.Group)
	}
	result.WriteString(":")
	if *showApp {
		result.WriteString(coords.Artifact)
	}
	result.WriteString(":")
	if *showVersion {
		result.WriteString(coords.Version)
	}

	fmt.Println(result.String())
}

func readCoords(fileName string) Coordinates {

	xmlFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	var coords Coordinates

	xml.Unmarshal(byteValue, &coords)

	return coords
}
