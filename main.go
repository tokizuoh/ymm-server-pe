package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

type notExistError struct{}

func (e *notExistError) Error() string {
	return fmt.Sprintf("csv file does not exist")
}

func extractCSV(args []string) (string, error) {
	if len(args) != 1 {
		return "", &notExistError{}
	}

	arg := args[0]

	if !strings.HasSuffix(arg, ".csv") {
		return "", &notExistError{}
	}

	return arg, nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	csvfile, err := extractCSV(args)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(csvfile)
}
