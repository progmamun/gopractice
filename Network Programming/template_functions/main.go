package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func main() {
	var data = struct {
		Name, Surname, Occupation, City string
	}{
		"Bojack", "Horseman", "Actor", "Los Angeles",
	}
	tpl, err := template.New("question-answer").Funcs(template.FuncMap{
		"upper": func(s string) string { return strings.ToUpper(s) },
		"lower": func(s string) string { return strings.ToLower(s) },
	}).Parse(`{{.Name}} {{.Surname}} - {{lower .Occupation}} from {{upper .City}}`)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	if err = tpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln("Error:", err)
	}
}
