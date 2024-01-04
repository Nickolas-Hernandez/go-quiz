package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, please prep that ass for this quiz.!")

	quizFile, err := os.Open("./data/problems.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer quizFile.Close()

	quizReader := csv.NewReader(quizFile)

	correct := 0
	incorrect := 0

	for {
		questionSet, err := quizReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var userAnswer string
		correctAnswer := questionSet[1]

		fmt.Println(questionSet[0])

		fmt.Scanln(&userAnswer)

		if userAnswer == correctAnswer {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Printf("Correct: %v. Incorrect: %v. ", correct, incorrect)
}
