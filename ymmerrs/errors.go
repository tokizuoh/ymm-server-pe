package ymmerrs

import "fmt"

type NotExistError struct{}

func (e *NotExistError) Error() string {
	return fmt.Sprintf("csv file does not exist")
}

type InvalidElementsCountError struct {
	N int
}

func (e *InvalidElementsCountError) Error() string {
	return fmt.Sprintf("number of elements in the array is not %v", e.N)
}
