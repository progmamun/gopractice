package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	var data = []struct {
		Name  string
		Score int
	}{
		{"Michelangelo", 30},
		{"Donatello", 50},
		{"Leonardo", 80},
		{"Raffaello", 100},
	}
	tpl, err := template.New("question-answer").Parse(`{{range .}}
{{.Name}} scored {{.Score}}. He did {{if lt .Score 50}}bad{{else if lt .Score 75}}okay{{else if lt .Score 90}}good{{else}}great{{end}}
	{{end}}`)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	if err = tpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln("Error:", err)
	}
}
