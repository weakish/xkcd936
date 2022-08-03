package xkcd936

import (
	"fmt"
	"testing"
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
	words := Words(4)
	if l := len(words); l != 4 {
		t.Errorf("expect 4 words, got %d\n", l)
	}
}
