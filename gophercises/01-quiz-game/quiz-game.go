package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	questionCol = 0
	answerCol   = 1
)

type QuizRecord struct {
	question string
	answer   string
}

func main() {
	records, err := getCSVRecords("quiz.csv")
	if err != nil {
		log.Fatal(err)
	}

	quiz := getQuiz(records)
	score, err := executeQuiz(quiz)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Score: %d/%d", score, len(quiz))
}

func getCSVRecords(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file %s: %w", fileName, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("couldn't read csv: %w", err)
	}

	return records, nil
}

func executeQuiz(quiz []QuizRecord) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	for i, quiz := range quiz {
		fmt.Printf("Problem %d: %s = ", i+1, quiz.question)
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("couldn't read line: %w", err)
		}

		input = strings.TrimSpace(input)
		if input == quiz.answer {
			score++
		}
	}

	return score, nil
}

func getQuiz(data [][]string) []QuizRecord {
	var quiz []QuizRecord
	for _, line := range data[1:] {
		quiz = append(quiz, QuizRecord{
			question: line[questionCol],
			answer:   line[answerCol],
		})
	}

	return quiz
}
