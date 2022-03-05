package main

import (
	"errors"
	"fmt"
)

func perimeter(length, width float64) (float64, error) {
	if length < 0 {
        return 0, errors.New("Length is negative") 
    }
    if width < 0 {
        return 0, errors.New("Width is negative")
    }
	return length*2 + width*2, nil
}

func main() {
	total, err := perimeter(8.2, -10)
	if(err != nil) {
		fmt.Println("INVALID")
		return 
	}
	fmt.Println("Total fencing: ", total)
}