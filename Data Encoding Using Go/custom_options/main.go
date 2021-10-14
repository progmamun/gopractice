package main

import (
	"encoding/csv"
	"log"
	"strings"
)

func main() {
	r := csv.NewReader(strings.NewReader("a b\ne f g\n1"))
	r.Comma = ' '
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range records {
		log.Println(r)
	}
}
