package main

import (
	"flag"
	"fmt"
	choose_your_own_adventure "github.com/IvanSharovarov/choose-your-own-adventure"
	"log"
	"net/http"
	"os"
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

	h := choose_your_own_adventure.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
