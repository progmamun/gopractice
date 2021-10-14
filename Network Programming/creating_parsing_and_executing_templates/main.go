package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	var data = struct {
		Question string
		Answer   int
	}{
		Question: "Answer to the Ultimate Question of Life, " +
			"the Universe, and Everything",
		Answer: 42,
	}
	tpl, err := template.New("question-answer").Parse(`
		<p>Question: {{.Question}}</p>
		<p>Answer: {{.Answer}}</p>
	`)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	if err = tpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln("Error:", err)
	}
}
