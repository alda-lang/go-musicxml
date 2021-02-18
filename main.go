package main

import (
	"alda.io/go-musicxml/model"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	xmlFile, err := os.Open("examples/helloworld.musicxml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var score model.ScorePartwise

	// This unmarshalling doesn't work. I think the model is incorrect? 
	xml.Unmarshal(byteValue, score)


	fmt.Println("First measure number: " + score.Part[0].Measure[0].MeasureAttributes.NumberAttr)
}
