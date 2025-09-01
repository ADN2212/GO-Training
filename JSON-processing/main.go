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
	CityId   uint   `json:"municipio_id"`
	SectorId uint   `json:"sector_id"`
	Name     string `json:"sector"`
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

	sectorFile, sectorErr := os.Open("./data/sectores.json")

	if sectorErr != nil {
		fmt.Printf("An error ocurred while loading the provs file: %s", sectorErr.Error())
		return
	}

	defer (func() {
		fmt.Println("Clossing the sectors file")
		sectorFile.Close()
	})()

	sectorBytes, sectorsReadError := io.ReadAll(sectorFile)

	if sectorsReadError != nil {
		fmt.Printf("An error ocurred while reading the sectors bytes: %s", sectorsReadError.Error())
		return
	}

	var sectors []Sector

	sectorsParseError := json.Unmarshal(sectorBytes, &sectors)

	if sectorsParseError != nil {
		fmt.Printf("An error ocurred while trying to parse the sectors: %s", sectorsParseError.Error())
		return
	}

	provsMap := make(map[uint]Province) //De esta forma se declara e inizializa el map al mismo tiempo.
	var currentProv Province

	for i := range provs {
		currentProv = provs[i]
		provsMap[currentProv.ProvinceId] = currentProv
	}

	//En este mapa la key es el id de la cuidad y el value es la cuidad, esto es para poder buscar cuidades en O(1) time
	citiesMap := make(map[uint]City)
	var currentCity City

	for i := range cities {
		currentCity = cities[i]
		citiesMap[currentCity.CityId] = currentCity
	}

	reducedProvs := make(map[string]map[string][]string)

	var currentSector Sector

	for i := range sectors {

		currentSector = sectors[i]
		currentCity = citiesMap[currentSector.CityId]
		currentProv = provsMap[currentCity.ProvinceId]

		if len(reducedProvs[currentProv.Name]) == 0 {
			reducedProvs[currentProv.Name] = map[string][]string{
				fmt.Sprint(currentCity.Name): {currentSector.Name},
			}
		} else {
			if len(reducedProvs[currentProv.Name][currentCity.Name]) == 0 {
				reducedProvs[currentProv.Name][currentCity.Name] = []string{currentSector.Name}
			} else {
				reducedProvs[currentProv.Name][currentCity.Name] = append(reducedProvs[currentProv.Name][currentCity.Name], currentSector.Name)
			}
		}
	}

	for provName, cityMap := range reducedProvs {
		fmt.Println("+", provName, "-----------------------------------------------------")
		for cityname, sectorsArray := range cityMap {
			fmt.Println("	-", cityname)
			for i := range sectorsArray {
				fmt.Println("		->", sectorsArray[i])
			}
		}
	}

}
