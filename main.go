package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/peterh/liner"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	quizes := []func() error{
		square2digit,
		sqrt,
		cuberoot,
		mult2x2,
		mult3x1,
		dayOfWeek,
	}

	q := rand.Intn(len(quizes))
	if err := quizes[q](); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ndigit(d int) int {
	switch d {
	case 1:
		return 1 + rand.Intn(9)
	case 2:
		return 10 + rand.Intn(90)
	case 3:
		return 100 + rand.Intn(900)
	}

	panic("unreached")
}

func square2digit() error {
	n := ndigit(2)
	return ask(n*n, "square of %v", n)
}

func mult2x2() error {
	m := ndigit(2)
	n := ndigit(2)
	return ask(n*m, "multiply %v x %v", m, n)
}

func mult3x1() error {
	m := ndigit(3)
	n := ndigit(1)
	return ask(n*m, "multiply %v x %v", m, n)
}

func sqrt() error {
	n := ndigit(2)
	return ask(n, "sqrt of %v", n*n)
}

func cuberoot() error {
	n := ndigit(2)
	return ask(n, "cube root of %v", n*n*n)
}

func dayOfWeek() error {
	// year := 1700 + rand.Intn(799)
	year := 1900 + rand.Intn(199)
	month := 1 + rand.Intn(12)
	// 0th day of next month == last day of chosen month
	days := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
	day := 1 + rand.Intn(days)
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return ask(int(date.Weekday()), "day of week for %v", date.Format("Jan _2 2006"))
}

func ask(answer int, s string, args ...interface{}) error {

	const prompt = "> "

	line := liner.NewLiner()
	defer line.Close()

	fmt.Printf(s, args...)
	fmt.Println()

	var guess int
	for {
		s, err := line.Prompt(prompt)
		if err == io.EOF {
			return err

		}
		if err != nil {
			return err
		}

		if s == "" {
			continue
		}

		guess, err = strconv.Atoi(s)
		if err != nil {
			fmt.Printf("error converting %q: %v\n", s, err)
			continue
		}

		break
	}

	if guess != answer {
		return fmt.Errorf("wanted: %v", answer)
	}

	return nil
}
