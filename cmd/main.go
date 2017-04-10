package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/codegoalie/string-bloom-filter/bloom"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const randos = 10000

func main() {
	filter := bloom.Filter{}

	filename := "/usr/share/dict/words"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", filename, err)
	}

	fmt.Print("Loading...")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		filter.Add(line)
	}
	fmt.Print("done.\n")

	// log.Printf("%+v", filter)
	rand.Seed(time.Now().UnixNano())

	found := 0
	for i := 0; i < randos; i++ {
		rando := randSeq(5)
		if filter.Check(string(rando)) {
			found++
			fmt.Printf("found = %+v\n", rando)
		}
		if i%100 == 0 {
			fmt.Printf("i = %+v\n", i)
		}
	}
	log.Printf("Found %d of %d: %d%%", found, randos, found/10000)

}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
