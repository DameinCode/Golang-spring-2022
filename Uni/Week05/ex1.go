package main

import (
	"fmt"
	"strings"
)

func scoreSummary(name string, score1, score2, score3 float64) {
	avg := (score1+score2+score3)/3
	fmt.Printf("%10s|%8.2f|%8.2f|%8.2f|%8.2f\n", name, score1, score2, score3, avg)
}

// "%.2f" .n3

func main() {
	fmt.Printf("%10s|%8s|%8s|%8s|%8s\n", "Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println(strings.Repeat("-", 46))

	
	scoreSummary("Gulbina", 90.70, 70.80, 89.90)
	scoreSummary("Aizhan", 90.0, 100.0, 90.90)
	scoreSummary("Asyl", 70.90, 50.0, 0.0)

}