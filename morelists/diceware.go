package morelists

import _ "embed"
import "strings"

var Diceware1, Diceware2, Diceware3, Diceware4 []string = diceware[:2048], diceware[2048 : 2048*2], diceware[2048*2 : 2048*3], diceware[2048*3:]

// https://theworld.com/~reinhold/diceware8k.txt
var diceware = strings.Split(strings.TrimSpace(diceware8k), "\n")

//go:embed diceware8k.txt
var diceware8k string
