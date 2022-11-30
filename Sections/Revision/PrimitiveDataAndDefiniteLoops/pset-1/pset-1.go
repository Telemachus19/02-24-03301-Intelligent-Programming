package main

import (
	"fmt"
)

func position(s_0, v_0, t, a float64) float64 {
	return (s_0 + v_0*t + (0.5 * a * t * t))
}
func main() {
	var s_0, v_0, t, a float64
	fmt.Print("Enter s_0: ")
	fmt.Scanln(&s_0)
	fmt.Print("Enter v_0: ")
	fmt.Scanln(&v_0)
	fmt.Print("Enter t: ")
	fmt.Scanln(&t)
	fmt.Print("Enter a: ")
	fmt.Scanln(&a)
	fmt.Print("The s position is ", position(s_0, v_0, t, a))
}
