package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// the write method receives a string slice and encodes in in CSV format.
func main() {
	const million = 1000000
	type Country struct {
		Code, Name string
		Population int
	}
	records := []Country{
		{Code: "IT", Name: "Italy", Population: 60 * million},
		{Code: "ES", Name: "Spain", Population: 46 * million},
		{Code: "JP", Name: "Japan", Population: 126 * million},
		{Code: "US", Name: "United States of America", Population: 327 * million},
	}
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	for _, r := range records {
		if err := w.Write([]string{r.Code, r.Name, strconv.Itoa(r.Population)}); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
}

// writeAll requires us to convert the records into a slice of string slices before calling it
/*func main() {
	const million = 1000000
	type Country struct {
		Code, Name string
		Population int
	}
	records := []Country{
		{Code: "IT", Name: "Italy", Population: 60 * million},
		{Code: "ES", Name: "Spain", Population: 46 * million},
		{Code: "JP", Name: "Japan", Population: 126 * million},
		{Code: "US", Name: "United States of America", Population: 327 * million},
	}
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	var ss = make([][]string, 0, len(records))
	for _, r := range records {
		ss = append(ss, []string{r.Code, r.Name, strconv.Itoa(r.Population)})
	}
	if err := w.WriteAll(ss); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}*/
