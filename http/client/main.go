package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Let's do some requests ...")
	//baseUrl := "https://restcountries.com/v3.1/all"

	//http.Client()
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Dark")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//This handle can be better 
	if resp.StatusCode != 200 {
		fmt.Println("Not a good awswer for the API")
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	responseText := ""
	currLine := "" 
	//Esto funciona como un interator, scanea linea por linea 
	//el bucle se detendra cundo Scan() retorne false
	for scanner.Scan() {
		currLine = scanner.Text()
		fmt.Println(currLine)
		responseText += currLine
	}

	//Luego podriamos parcear esta data a una struct the Go.
	//...

}
