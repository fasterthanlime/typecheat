package main

import (
	"bufio"
	"fmt"
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
		letters: []Letter{},
	}
	if len(os.Args[1:]) < 1 {
		fmt.Fprintf(os.Stderr, "Place letter groups as subsequent arguments, like this:\n\t%s rsalp eihon wxast tyrdi pioe v gnyde\n", os.Args[0])
		os.Exit(1)
	}
	for _, letter := range os.Args[1:] {
		l := []string{}
		for _, char := range letter {
			l = append(l, string(char))
		}
		spec.letters = append(spec.letters, l)
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
		for j := len(indices) - 1; j >= 0; j-- {
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
