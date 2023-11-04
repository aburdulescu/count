package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if err := mainErr(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func mainErr() error {
	args := os.Args[1:]

	if len(args) < 1 {
		return fmt.Errorf("need direction: up or down")
	}

	direction := args[0]

	if len(args) < 2 {
		return fmt.Errorf("need duration")
	}

	duration, err := time.ParseDuration(args[1])
	if err != nil {
		return err
	}

	switch direction {
	case "down":
		down(duration)
	case "up":
		up(duration)
	default:
		return fmt.Errorf("invalid direction: %s", direction)
	}

	return nil
}

const step = 1 * time.Second

func down(duration time.Duration) {
	c := time.Tick(step)

	s := duration.String()
	prevLen := len(s)

	fmt.Printf("\t%s\r", s)

	for range c {
		duration -= step

		s = duration.String()

		spaceCount := prevLen - len(s)

		if spaceCount > 0 {
			fmt.Printf("\t%s%s\r", s, strings.Repeat(" ", prevLen-len(s)))
		} else {
			fmt.Printf("\t%s\r", s)
		}

		prevLen = len(s)

		if duration == 0 {
			fmt.Println("")
			break
		}
	}
}

func up(duration time.Duration) {
	c := time.Tick(step)

	count := 1 * time.Second

	for range c {
		fmt.Printf("\t%s\r", count)

		count += step

		if count > duration {
			fmt.Println("")
			break
		}
	}
}
