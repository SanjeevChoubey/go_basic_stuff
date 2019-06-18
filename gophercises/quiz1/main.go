package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type QuestionAnswer struct {
	question string `json: "question"`
	answer   string `json:"answer"`
}

func main() {
	csvReader()
}

func csvReader() {
	// Open file
	f, err := os.Open("./problem.csv")
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize Reader
	reader := csv.NewReader(f)

	//Seprator
	reader.Comma = ','
	reader.Comment = '#'
	var que []QuestionAnswer
	for {
		records, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		// Keep all question and answer in struct
		que = append(que, QuestionAnswer{
			question: records[0],
			answer:   records[1],
		})
	}
	// read struct
	var totalQuestions, correctAnswers int
	for i := 0; i < len(que); i++ {
		var ans string
		totalQuestions += 1
		fmt.Println("problem #", totalQuestions, " :", que[i].question)
		fmt.Scan(&ans)
		if ans == que[i].answer {
			correctAnswers += 1
		}
	}

	fmt.Println("Total question were ", totalQuestions, " and correct answers were ", correctAnswers)
}
