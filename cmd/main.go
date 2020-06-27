package main

import (
	"flag"
	"fmt"
	choose_your_own_adventure "github.com/IvanSharovarov/choose-your-own-adventure"
	"os"
)

func main() {
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
	fmt.Printf("%+v.\n", story)
}
