package main

import (
	"encoding/csv"
	"log"
	"strings"
)

func main() {
	r := csv.NewReader(strings.NewReader("a,b,c\ne,f,g\n1,2,3"))
	for {
		r, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(r)
	}
}

/*func main() {
r := csv.NewReader(strings.NewReader("a,b,c\ne,f,g\n1,2,3"))
records, err := r.ReadAll()
if err != nil {
log.Fatal(err)
}
for _, r := range records {
log.Println(r)
}
}*/
