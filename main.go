package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	in, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	list := strings.Split(string(in), "\n")

	counter := 0
	for _, str := range list {
		if checkRules(str) {
			counter++
		}
	}

	log.Println("Execution time:", time.Since(start))
	log.Println(counter)
}

func checkRules(str string) bool {

}
