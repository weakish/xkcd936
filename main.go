package main

import (
	"fmt"
	"github.com/weakish/xkcd936/xkcd936"
	"os"
	"strconv"
)

const version = "0.2.0"

func usage(exitCode int) {
	fmt.Print("xkcd936 [WORD_LIST_LENGTH]\n\nIf not specified, it will generate four words.\n")
	os.Exit(exitCode)
}

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println(xkcd936.Phrase(xkcd936.Words(4)))
	case 2:
		switch arg := os.Args[1]; arg {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12":
			n, _ := strconv.ParseInt(arg, 10, 0)
			fmt.Println(xkcd936.Phrase(xkcd936.Words(int(n))))
		case "-h", "--help", "help":
			usage(0)
		case "--version", "version":
			fmt.Printf("xkcd936 %s\n", version)
		default:
			usage(64)
		}
	default:
		usage(64)
	}
}
