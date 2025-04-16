package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	limit := flag.Int("limit", 30, "allows a custom duration")
	shuffle := flag.Bool("shuffle", false, "allows random shuffling of problems")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	if *shuffle {
		randGen.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	correct := 0

	fmt.Println("Press Return to begin quiz.")
	var dummy string
	fmt.Scanln(&dummy)

	fmt.Println("Begin!")

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	// time.AfterFunc(time.Duration(*limit)*time.Second, func() {
	// 	fmt.Printf("\nTime is up. You scored %d out of %d.\n", correct, len(problems))
	// 	os.Exit(0)
	// })

	//create a go routine to listen for answers
	//create a go routine to return upon timer running out

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime is up. You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("\nQuiz completed. You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
