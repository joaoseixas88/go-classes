package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func main() {

	person := Person{
		Name:    "Joao",
		Surname: "Seixas",
		Age:     35,
	}
	fmt.Println(person)

	p, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	os.WriteFile("json.json", p, os.ModePerm)

	file, err := os.ReadFile("json.json")
	if err != nil {
		fmt.Println("Deu ruim", err)
	}
	var person_1 Person
	unMarshalErr := json.Unmarshal(file, &person_1)
	if unMarshalErr != nil {
		fmt.Println("Deu ruim", unMarshalErr)
	}
	fmt.Println("Retorno:", person_1)

}
