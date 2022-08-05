package main

import (
	"flag"
	"fmt"

	"github.com/weakish/xkcd936/xkcd936"
)

const semver = "0.3.0"

func main() {
	var titlized *bool = flag.Bool("t", false, "titlized words (English only)")
	var version *bool = flag.Bool("V", false, "show version")
	var number *int = flag.Int("n", 4, "number of words, 1 <= n <= 12")
	var list *string = flag.String("l", "en", "word list: chs, cht, cs, en, fr, it, jp, ko, es, diceware{1,2,3,4} (split from diceware8k), diceware2k")
	flag.Parse()

	if *version {
		fmt.Printf("xkcd936 %s\n", semver)
	} else {
		fmt.Println(xkcd936.Phrase(xkcd936.Words(*number, *list), *titlized))
	}
}
