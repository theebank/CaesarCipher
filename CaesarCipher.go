package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CaesarCipher("I LOVE CS!", 5))
	fmt.Println("-------------------End of exercise 1-------------------")
	messages := []string{"Csi2520", "CSI2120", "3 Paradigms",
		"Go is 1st", "Prolog is 2nd", "Scheme is 3rd",
		"uottawa.ca", "csi/elg/ceg/seg", "800 King Edward"}
	//Exercise 02
	//
	//
	messagechan := make(chan int, len(messages)+1)
	encrypt := make(chan string, len(messages)+1)

	go CaesarCipherList(messages[:], 2, messagechan, encrypt)
	for i := 0; i < len(messages); i++ {
		messagechan <- i
	}
	close(messagechan)

	for j := 0; j < len(messages); j++ {
		fmt.Println(<-encrypt)
	}
	fmt.Println("-------------------End of exercise 2-------------------")
	//Exercise 03
	//
	//
	messagechan3 := make(chan int, len(messages)+1)
	encrypt3 := make(chan string, len(messages)+1)

	go CaesarCipherList(messages[:], 2, messagechan3, encrypt3)
	go CaesarCipherList(messages[:], 2, messagechan3, encrypt3)
	go CaesarCipherList(messages[:], 2, messagechan3, encrypt3)

	for i := 0; i < len(messages); i++ {
		messagechan3 <- i
	}
	close(messagechan3)

	for j := 0; j < len(messages); j++ {
		fmt.Println(<-encrypt3)
	}

}

func CaesarCipherList(s []string, shift int, jobs chan int, results chan string) {
	for n := range jobs {
		results <- CaesarCipher(s[n], shift)

	}

}

//exercise01
func CaesarCipher(m string, shift int) string {
	var r []rune
	m = strings.ToUpper(m)
	for _, c := range m {
		if unicode.IsLetter(c) {
			r = append(r, c)
		}
	}
	var buffer bytes.Buffer
	for _, c := range r {
		s := int(math.Mod(float64((int(c)-65)+shift), float64(26))) + 65

		buffer.WriteRune(rune(s))
	}
	newString := buffer.String()

	return newString
}
