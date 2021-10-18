// source: https://tutorialedge.net/golang/parsing-json-with-golang/
package helpers

import (
	"chandylamport/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const path = "/home/freedy/Documents/master/RedesYDistribuidos/Practica1/chandylamport/"

func readFile(fileName string) []byte {
	data, err := ioutil.ReadFile(path + fileName)
	if err != nil {
		panic(err)
	}

	return data
}

func ReadNetConfig(fileName string) []models.ProcessInfo {
	data := readFile(fileName)
	fmt.Println("Reading file " + fileName)

	var myJson []models.ProcessInfo
	err := json.Unmarshal(data, &myJson)
	if err != nil {
		panic(err)
	}

	return myJson
}
