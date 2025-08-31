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

	// fmt.Println("Lista de provincias:----------------------------------------------------------------------")
	// for i := range provs {
	// 	fmt.Println(provs[i].Name)
	// }

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

	// fmt.Println("Lista de cuidades: ----------------------------------------------------------------------")
	// for i := range cities {
	// 	fmt.Println(cities[i].Name)
	// }

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

	// fmt.Println("Lista de sectores: ----------------------------------------------------------------------")
	// for i := range sectors {
	// 	fmt.Println(sectors[i])
	// }

	//Este map servira para agrupar los sectores por cuidad
	sectorsMap := make(map[uint][]Sector) //De esta forma se declara e inizializa el map al mismo tiempo.
	var currentSector Sector

	for i := range sectors {
		currentSector = sectors[i]
		if len(sectorsMap[currentSector.CityId]) == 0 {
			sectorsMap[currentSector.CityId] = []Sector{currentSector}
		} else {
			sectorsMap[currentSector.CityId] = append(sectorsMap[currentSector.CityId], currentSector)
		}
	}

	//fmt.Println(sectorsMap)
	// for cityId, sectorsArray := range sectorsMap {
	// 	fmt.Println(cityId)
	// 	for i := range sectorsArray {
	// 		fmt.Println(sectorsArray[i])
	// 	}
	// }

	//El siguiente para aculumar las cuidades por porvincia:
	citiesMap := make(map[uint][]City)
	var currentCity City

	for i := range cities {

		currentCity = cities[i]

		if len(citiesMap[currentCity.ProvinceId]) == 0 {
			citiesMap[currentCity.ProvinceId] = []City{currentCity}
		} else {
			citiesMap[currentCity.ProvinceId] = append(citiesMap[currentCity.ProvinceId], currentCity)
		}
	}

	// for provId, cityArray := range citiesMap {
	// 	fmt.Println(provId)
	// 	for i := range cityArray {
	// 		fmt.Println(cityArray[i])
	// 	}
	// }

	



}
