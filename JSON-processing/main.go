package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Province struct {
	ProvinceId uint   `json:"provincia_id"`
	Name       string `json:"provincia"`
}

type City struct {
	ProvinceId uint   `json:"provincia_id"`
	CityId     uint   `json:"municipio_id"`
	Name       string `json:"municipio"`
}

type Sector struct {
	CityId     uint   `json:"municipio_id"`
	SectorId uint `json:"sector_id"`
	Name string `json:"sector"`
}

func main() {

	provsFile, ProvsErr := os.Open("./data/provincias.json")

	if ProvsErr != nil {
		fmt.Printf("An error ocurred while loading the provs file: %s", ProvsErr.Error())
		return
	}

	defer (func() {
		fmt.Println("Clossing the provs file")
		provsFile.Close()
	})()

	byteValue, ReadErr := io.ReadAll(provsFile)

	if ReadErr != nil {
		fmt.Printf("An error ocurred while reading the provs bytes: %s", ReadErr.Error())
		return
	}

	//Como el JSON es un array de provincias se puede hacer de esta forma, y no agregar un valor extra para la raiz del JSON.
	var provs []Province

	provsParseErr := json.Unmarshal(byteValue, &provs)

	if provsParseErr != nil {
		fmt.Printf("An error ocurred while trying to  parse the provs: %s", provsParseErr.Error())
		return
	}

	fmt.Println("Lista de provincias:----------------------------------------------------------------------")
	for i := range provs {
		fmt.Println(provs[i].Name)
	}

	//TODO: abstraer toda esta logica repetida usando un funcion que retorne un generic type.
	cityFile, cityErr := os.Open("./data/municipios.json")

	if cityErr != nil {
		fmt.Printf("An error ocurred while loading the provs file: %s", cityErr.Error())
		return
	}

	defer (func() {
		fmt.Println("Clossing the cities file")
		cityFile.Close()
	})()

	citiesBytes, citiesReadError := io.ReadAll(cityFile)

	if citiesReadError != nil {
		fmt.Printf("An error ocurred while reading the cities bytes: %s", citiesReadError.Error())
		return
	}

	var cities []City

	citiesParseErro := json.Unmarshal(citiesBytes, &cities)

	if citiesParseErro != nil {
		fmt.Printf("An error ocurred while trying to parse the cities: %s", citiesParseErro.Error())
		return
	}

	fmt.Println("Lista de cuidades: ----------------------------------------------------------------------")
	for i := range cities {
		fmt.Println(cities[i].Name)
	}

	





}
