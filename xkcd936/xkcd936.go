package xkcd936

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
	"github.com/weakish/xkcd936/morelists"
)

func Words(n int, lang string) []string {
	if n <= 0 {
		n = 4 // xkcd936 uses four words
	} else if n > 12 {
		n = 12 // BIP32 uses 12 words
	}

	switch lang {
	case "chs":
		bip39.SetWordList(wordlists.ChineseSimplified)
	case "cht":
		bip39.SetWordList(wordlists.ChineseTraditional)
	case "cs":
		bip39.SetWordList(wordlists.Czech)
	case "en":
		bip39.SetWordList(wordlists.English)
	case "fr":
		bip39.SetWordList(wordlists.French)
	case "it":
		bip39.SetWordList(wordlists.Italian)
	case "jp":
		bip39.SetWordList(wordlists.Japanese)
	case "ko":
		bip39.SetWordList(wordlists.Korean)
	case "es":
		bip39.SetWordList(wordlists.Spanish)
	case "diceware1":
		bip39.SetWordList(morelists.Diceware1)
	case "diceware2":
		bip39.SetWordList(morelists.Diceware2)
	case "diceware3":
		bip39.SetWordList(morelists.Diceware3)
	case "diceware4":
		bip39.SetWordList(morelists.Diceware4)
	case "dicewareshort", "diceware2k":
		bip39.SetWordList(morelists.DicewareShort)
	default:
		bip39.SetWordList(wordlists.English)
	}

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	words := strings.Split(mnemonic, " ")
	return words[0:n]
}

func Phrase(words []string, titlized bool) string {
	if titlized {
		var phrase string = ""
		for _, word := range words {
			phrase += cases.Title(language.English).String(word)
		}
		return phrase
	} else {
		return strings.Join(words, " ")
	}
}
