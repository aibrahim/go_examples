package main

import "fmt"

func abs(x float64) float64 {
	if x < 0 {
		return -1*x
	}
	return x
}

func Sqrt(x float64) float64 {
	guess := 1.0
	delta := 1.0
	
	for {
		guess = (guess + x/guess)/2
		delta = abs(x - guess*guess)
		if delta <= .000000001 {
			return guess
		}
	}
}

func main() {
	fmt.Println(Sqrt(4))
}

