package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataType struct {
	ID           int     `json:"id"`
	English      *string `json:"English"`
	GlossEnglish *string `json:"Gloss (english)"`
	GlossDharug  *string `json:"Dharug(Gloss)"`
	Dharug       *string `json:"Dharug"`
	Topic        *string `json:"Topic"`
	ImageName    *string `json:"Image Name (optional)"`
	Recording    *string `json:"recording"`
	Completed    bool    `json:"completed"`
}

var Data []DataType

func readJson() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		panic(fmt.Sprintf("Error opening json file\n%v\n", err))
	}

	err = json.Unmarshal(file, &Data)
	if err != nil {
		panic(fmt.Sprintf("Error decoding json file\n%v\n", err))
	}
}

func insertData() {
}

func main() {
	readJson()
}
