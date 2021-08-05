package main

import (
	"fmt"
	"log"
)

type notExistError struct{}

func (e *notExistError) Error() string {
	return fmt.Sprintf("csv file does not exist")
}

func main() {
	log.Println("init")
}
