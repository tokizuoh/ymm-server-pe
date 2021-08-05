package main

import "fmt"

type notExistError struct{}

func (e *notExistError) Error() string {
	return fmt.Sprintf("csv file does not exist")
}

type invalidElementsCount struct {
	n int
}

func (e *invalidElementsCount) Error() string {
	return fmt.Sprintf("number of elements in the array is not %v", e.n)
}
