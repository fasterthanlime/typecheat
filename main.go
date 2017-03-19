package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

type Letter []string

type Spec struct {
	letters []Letter
}

func main() {
	spec := Spec{
		letters: []Letter{
			[]string{"r", "s", "a", "l", "p"},
			[]string{"e", "i", "h", "o", "n"},
			[]string{"w", "x", "a", "s", "t"},
			[]string{"t", "y", "r", "d", "i"},
			[]string{"p", "i", "o", "e"},
			[]string{"v"},
			[]string{"g", "n", "y", "d", "e"},
		},
	}

	t1 := time.Now()
	log.Println("loading word database...")
	words := make(map[string]bool)

	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	t2 := time.Now()

	log.Printf("loaded in %s", t2.Sub(t1))

	indices := make([]int, len(spec.letters))
	for i := 0; i < len(indices); i++ {
		indices[i] = 0
	}

	for {
		s := ""
		for i := 0; i < len(indices); i++ {
			s += spec.letters[i][indices[i]]
		}

		if words[s] {
			log.Printf("%s is a word", s)
		}

		progressed := false
		for j := len(indices) - 1; j > 0; j-- {
			if indices[j]+1 < len(spec.letters[j]) {
				indices[j]++
				for k := j + 1; k < len(indices); k++ {
					indices[k] = 0
				}

				progressed = true
				break
			}
		}

		if !progressed {
			break
		}
	}
}
