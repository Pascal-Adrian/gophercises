package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getProblems(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func ask(problems [][]string) (int, int) {
	var answer string
	var correct int

	for i, row := range problems {
		fmt.Printf("%d: %s = ", i, row[0])

		_, err := fmt.Scanf("%s", &answer)

		if err != nil {
			continue
		}

		if row[1] == answer {
			correct++
		}
	}

	return correct, len(problems)
}

func main() {
	pathnamePtr := flag.String("filepath", "problems.csv", "filepath to the CSV file with the problems")

	flag.Parse()

	problems, err := getProblems(*pathnamePtr)
	check(err)

	correct, total := ask(problems)

	fmt.Printf("\nYou answered %d out of %d question(s) correct.\n", correct, total)
}
