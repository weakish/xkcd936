package main

import (
	"flag"
	"fmt"

	"github.com/weakish/xkcd936/xkcd936"
)

const semver = "0.3.0"

func main() {
	var titlized *bool = flag.Bool("t", false, "join titlized words")
	var version *bool = flag.Bool("V", false, "show version")
	var number *int = flag.Int("n", 4, "number of words, 1 <= n <= 12")
	flag.Parse()

	if *version {
		fmt.Printf("xkcd936 %s\n", semver)
	} else {
		fmt.Println(xkcd936.Phrase(xkcd936.Words(*number), *titlized))
	}
}
