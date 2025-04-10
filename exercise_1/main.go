package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	//implement change duration with flag

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
	correct := 0

	duration := 5

	//implement start timer and quiz by pressing return
	fmt.Println("Press Return to begin quiz.")
	var dummy string
	fmt.Scanln(&dummy)

	fmt.Println("Starting timer")
	time.AfterFunc(time.Duration(duration)*time.Second, func() {
		fmt.Printf("\nTime is up. You scored %d out of %d.\n", correct, len(problems))
		os.Exit(0)
	})

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
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

//create a timer with a default duration of 30 seconds (timer will start once created)
//	timer duration is customizable with a flag
//stop the quiz when the timer is done
//stop the timer if the quiz is completed before timer

//user should be prompted to hit the enter (or some key) to start the timer and the quiz: "Press enter to begin."
