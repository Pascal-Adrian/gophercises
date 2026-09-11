package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pascal-Adrian/gophercises/quiz-game/internal/quiz"
	"github.com/Pascal-Adrian/gophercises/quiz-game/internal/read"
)

func main() {
	pathPtr := flag.String("path", "./problems.csv", "a path")

	flag.Parse()

	problems, err := read.ReadCsv(*pathPtr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(65)
	}

	quiz.Run(problems)
}
