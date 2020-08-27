package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNum int      `json:"entry_number"`
	Species  PokeSpec `json:"pokemon_species"`
}

type PokeSpec struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	respon, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		panic(err)
	}

	dat, err := ioutil.ReadAll(respon.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObj Response
	json.Unmarshal(dat, &responseObj)

	for i := 0; i < len(responseObj.Pokemon); i++ {
		fmt.Println(responseObj.Pokemon[i].Species.Name)
		fmt.Println(responseObj.Pokemon[i].EntryNum)
	}
}
