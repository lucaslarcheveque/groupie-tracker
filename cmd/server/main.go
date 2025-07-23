package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println("Error fetching:", err)
		os.Exit(1)

	}
	defer response.Body.Close()

	reponseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(reponseData))

}
