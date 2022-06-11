package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	records := getCSVRecords("quiz.csv")
	quiz := getQuiz(records)
	score := executeQuiz(quiz)
	fmt.Printf("Score: %d/%d", score, len(quiz))
}

func getCSVRecords(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func executeQuiz(quiz []Quiz) int {
	reader := bufio.NewReader(os.Stdin)
	index := 1
	score := 0
	for _, quiz := range quiz {
		fmt.Printf("Problem %d: %s = ", index, quiz.question)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == quiz.answer {
			score++
		}
		index++
	}

	return score
}

func getQuiz(data [][]string) []Quiz {
	var quiz []Quiz
	for i, line := range data {
		if i == 0 {
			continue
		}

		var record Quiz
		for col, value := range line {
			if col == 0 {
				record.question = value
			} else if col == 1 {
				record.answer = value
			}
		}
		quiz = append(quiz, record)
	}

	return quiz
}
