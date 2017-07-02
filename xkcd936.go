package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func main() {

	word_list_length := 234937
	var random_numbers [4]int = get_random_numbers(word_list_length)

	for _, n := range random_numbers {

		fmt.Print(strings.Title(words[n]))
	}
}

func get_random_numbers(word_list_length int) [4]int {

	max := big.NewInt(int64(word_list_length))
	var random_numbers [4]int

	for i := 0; i < 4; i++ {

		random_numbers[i] = int(rand_int(max).Uint64())
	}

	return random_numbers
}

func rand_int(max *big.Int) *big.Int {

	result, _ := rand.Int(rand.Reader, max)

	return result
}
