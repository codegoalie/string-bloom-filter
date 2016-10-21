package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/codegoalie/string-bloom-filter/bloom"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	filter := bloom.Filter{}

	filename := "/usr/share/dict/words"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", filename, err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		filter.Add(line)
	}

	log.Printf("%+v", filter)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		rando := randSeq(5)
		log.Printf("%s: %v", rando, filter.Check(string(rando)))
	}

}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
