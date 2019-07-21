package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int `json:"sAge"`
}

func main() {
	var firstPerson person

	bs := []byte(`{"Name":"SomeName", "sAge":20}`)
	json.Unmarshal(bs, &firstPerson)

	fmt.Println(firstPerson.Name)
	fmt.Println(firstPerson.Age)
	fmt.Printf("%T\n", firstPerson)
}
