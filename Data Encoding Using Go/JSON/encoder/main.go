package main

import (
	"encoding/json"
	"log"
	"os"
)

type Character struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Job         string `json:"job,omitempty"`
	YearOfBirth int    `json:"year_of_birth,omitempty"`
}

func main() {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	c := Character{
		Name:        "Charles Dexter",
		Surname:     "Ward",
		YearOfBirth: 1902,
	}
	if err := e.Encode(c); err != nil {
		log.Fatalln(err)
	}
}
