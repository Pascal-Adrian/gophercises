package quiz

import (
	"fmt"
	"strings"
)

func askQuestion(index int, pair []string) bool {
	fmt.Printf("%d. %s = ", index+1, pair[0])

	var answer string

	fmt.Scanln(&answer)

	answer = strings.TrimSpace(answer)

	return strings.EqualFold(answer, strings.TrimSpace(pair[1]))
}

func askQuestions(pairs [][]string) []bool {
	var responses []bool

	for i, pair := range pairs {
		response := askQuestion(i, pair)

		responses = append(responses, response)
	}

	return responses
}

func countAnswers(answers []bool) (correct int, total int) {
	correct = 0

	for _, answer := range answers {
		if answer {
			correct++
		}
	}

	return correct, len(answers)
}

func Run(pairs [][]string) {
	correct, total := countAnswers(askQuestions(pairs))

	fmt.Println()
	fmt.Printf("You got %d correct answer(s) out of %d.\n", correct, total)
}
