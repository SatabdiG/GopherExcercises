package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	question string
	answer   string
}

type Player struct {
	name  string
	score int
}

func (p *Player) changeName() {
	p.name = "Taka"
}

func read_problems(reader [][]string) []Problem {
	result := []Problem{}
	for _, lines := range reader {
		result = append(result, Problem{lines[0], lines[1]})
	}
	return result
}

func isTimeOver(startTime time.Time, defaultPeriod int) bool {

	currentTime := time.Now()
	if time.Duration(currentTime.Sub(startTime).Seconds()) > time.Duration(defaultPeriod) {
		return true
	}
	return false
}

func main() {
	fileName := flag.String("file", "./data/problems.csv", "Provide an absolute file name")
	defaultPeriod := flag.Int("time", 10, "Enter the time period to run")

	flag.Parse()

	file, err := os.Open(*fileName)

	if err != nil {
		fmt.Printf("Unable to open file because of %s", err)
	}

	reader := csv.NewReader(file)

	fmt.Println("Welcome to the QUIZ")
	player := Player{"", 0}

	fmt.Println("Enter your name ...")
	fmt.Scan(&player.name)

	fmt.Printf("%s your time starts now ...\n", player.name)
	startTime := time.Now()

	csvLines, err := reader.ReadAll()
	problems := read_problems(csvLines)

	for i, problem := range problems {

		timeUp := isTimeOver(startTime, *defaultPeriod)

		if timeUp {
			fmt.Println("Sorry Time's up!")
			break
		}
		fmt.Printf("%d/%d Question is %s : \n", i+1, len(problems), problem.question)
		enteredAns := ""
		fmt.Scan(&enteredAns)
		if problem.answer == enteredAns {
			fmt.Println("Correct Ans")
			player.score++
		} else {
			fmt.Println("Incorrect Ans, lets try again")
		}

	}
	fmt.Printf(" %s,  Score is %d/%d \n", player.name, player.score, len(problems))

}
