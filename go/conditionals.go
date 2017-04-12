package main

import (
	"fmt"
	"math"
)

func sqrt_func(x float64) string {
	if x < 0 {
		return sqrt_func(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt_func(2), sqrt_func(-9))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
