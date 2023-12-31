package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
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
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage:

  count up DURATION [MESSAGE]

    Count one second at a time up to the given DURATION.

  count down DURATION [MESSAGE]

    Count one second at a time down from the given DURATION.

DURATION format:
  A duration is a sequence of decimal numbers,
  each with optional fraction and a unit
  suffix, such as "300ms", "1.5h" or "2h45m".
  Valid time units are "ns", "us", "ms", "s", "m", "h".

Flags:
`)
		flag.PrintDefaults()
	}
	printVersion := flag.Bool("version", false, "Print version")
	flag.Parse()

	if *printVersion {
		bi, _ := debug.ReadBuildInfo()
		fmt.Print(bi)
		return nil
	}

	args := flag.Args()

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

	msg := ""
	if len(args) > 2 {
		msg = flag.Arg(2)
	}

	switch direction {
	case "down":
		down(duration, msg)
	case "up":
		up(duration, msg)
	default:
		return fmt.Errorf("invalid direction: %s", direction)
	}

	return nil
}

const step = 1 * time.Second

func down(duration time.Duration, msg string) {
	c := time.Tick(step)

	s := duration.String()
	prevLen := len(s)

	if msg != "" {
		fmt.Printf("\t%s\n", msg)
	}

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

func up(duration time.Duration, msg string) {
	c := time.Tick(step)

	count := 1 * time.Second

	if msg != "" {
		fmt.Printf("\t%s\n", msg)
	}

	for range c {
		fmt.Printf("\t%s\r", count)

		count += step

		if count > duration {
			fmt.Println("")
			break
		}
	}
}
