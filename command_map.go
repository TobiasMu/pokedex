package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	//"encoding/json"
)

func mapCommand() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	fmt.Println("Mapping...")
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	var data map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to decode %v", err)
	}
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data["results"]))

	return nil
}

/*
curl --get https://pokeapi.co/api/v2/location-area/


*/
