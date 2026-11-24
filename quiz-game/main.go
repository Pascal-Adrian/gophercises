package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
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

func readToChan(channel chan string, reader *bufio.Reader) {
	answer, err := reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			channel <- answer
		} else {
			channel <- ""
		}
	} else {
		channel <- answer
	}
}

func ask(reader *bufio.Reader, problems [][]string, timeout time.Duration) (int, error) {
	timeup := time.After(timeout)

	channel := make(chan string, 1)

	var correct int

	for i, row := range problems {
		fmt.Printf("%d: %s = ", i, row[0])

		go readToChan(channel, reader)

		select {
		case <-timeup:
			return correct, errors.New("Time is up!")
		case value, _ := <-channel:
			if strings.EqualFold(strings.TrimSpace(row[1]), strings.TrimSpace(value)) {
				correct++
			}
		}
	}

	return correct, nil
}

func main() {
	pathnamePtr := flag.String("filepath", "problems.csv", "filepath to the CSV file with the problems")
	timeoutPtr := flag.Int("timeout", 30, "time, in seconds, to solve the problems")
	shufflePtr := flag.Bool("shuffle", false, "if the problems are shuffled before start or not")

	flag.Parse()

	problems, err := getProblems(*pathnamePtr)
	check(err)

	total := len(problems)

	if *shufflePtr {
		rand.Shuffle(total, func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press enter to start!")
	reader.ReadString('\n')

	correct, err := ask(reader, problems, time.Second*time.Duration(*timeoutPtr))

	if err != nil {
		fmt.Println("\n\n", err.Error())
	}

	fmt.Printf("\nYou answered %d out of %d question(s) correct.\n", correct, total)
}
