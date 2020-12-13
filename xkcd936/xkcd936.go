package xkcd936

import (
	"github.com/tyler-smith/go-bip39"
	"strings"
)

func Words(n int) []string {
	if n <= 0 {
		n = 4 // xkcd936 uses four words
	} else if n > 12 {
		n = 12 // BIP32 uses 12 words
	}

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	words := strings.Split(mnemonic, " ")
	return words[0:n]
}

func Phrase(words []string) string {
	phrase := ""
	for _, word := range words {
		phrase += strings.Title(word)
	}
	return phrase
}
