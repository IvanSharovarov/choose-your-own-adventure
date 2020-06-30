package main

import (
	"flag"
	"fmt"
	choose_your_own_adventure "github.com/IvanSharovarov/choose-your-own-adventure"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	fname := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fname)

	f, err := os.Open(*fname)
	if err != nil {
		panic(err)
	}

	story, err := choose_your_own_adventure.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := choose_your_own_adventure.NewHandler(story,
		choose_your_own_adventure.WithTemplate(tpl),
		choose_your_own_adventure.WithPathFunc(pathFn),
	)

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h), mux)
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTmpl = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose your own adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
</body>
</html>`