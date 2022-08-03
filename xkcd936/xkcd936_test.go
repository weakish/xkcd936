package xkcd936

import (
	"fmt"
	"testing"

	"github.com/weakish/xkcd936/morelists"
)

func ExamplePhrase() {
	fmt.Println(Phrase([]string{"hello", "world"}, false))
	// Output: hello world
}

func TestPhraseTitlized(t *testing.T) {
	if Phrase([]string{"hello", "world"}, true) != "HelloWorld" {
		t.Fail()
	}
}

func TestWords(t *testing.T) {
	words := Words(4, "en")
	if l := len(words); l != 4 {
		t.Errorf("expect 4 words, got %d\n", l)
	}
}

func TestWordsWithDicewareList(t *testing.T) {
	if len(morelists.Diceware1) != 2048 || len(morelists.Diceware2) != 2048 || len(morelists.Diceware3) != 2048 || len(morelists.Diceware4) != 2048 {
		t.Fail()
	}
}
