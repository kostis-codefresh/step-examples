package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Coordinates struct {
	Group    string `xml:"groupId" json:"groupId"`
	Artifact string `xml:"artifactId" json:"artifactId"`
	Version  string `xml:"version" json:"version"`
}

func main() {

	fileName := flag.String("f", "pom.xml", "filename to read")
	useJsonOutput := flag.Bool("json", false, "format output in json")
	flag.Parse()

	coords := readCoords(*fileName)

	if *useJsonOutput {
		bolB, _ := json.Marshal(coords)
		fmt.Println(string(bolB))
	} else {
		fmt.Printf("%s:%s:%s\n", coords.Group, coords.Artifact, coords.Version)
	}

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
