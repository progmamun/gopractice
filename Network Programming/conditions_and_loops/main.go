package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	var data = []struct {
		Question, Answer string
	}{{
		Question: "Answer to the Ultimate Question of Life, " +
			"the Universe, and Everything",
		Answer: "42",
	}, {
		Question: "Who you gonna call?",
		Answer:   "Ghostbusters",
	}}
	tpl, err := template.New("question-answer").Parse(`{{range .}}
Question: {{.Question}}
Answer: {{.Answer}}
	{{end}}`)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	if err = tpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln("Error:", err)
	}
}
