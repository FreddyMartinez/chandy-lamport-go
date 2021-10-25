// source: https://tutorialedge.net/golang/parsing-json-with-golang/
package helpers

import (
	"chandylamport/models"
	"encoding/json"
	"io/ioutil"
)

// const path = "/home/a846866/Documents/"
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

	var myJson []models.ProcessInfo
	err := json.Unmarshal(data, &myJson)
	if err != nil {
		panic(err)
	}

	return myJson
}

func ReadTaskList(fileName string) []models.Task {
	data := readFile(fileName)

	var myJson []models.Task
	err := json.Unmarshal(data, &myJson)
	if err != nil {
		panic(err)
	}

	return myJson
}
