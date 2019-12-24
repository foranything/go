package main

import (
	"fmt"
	"math"
)

func mySqrt(x float64) (z float64) {
	z = x
	r, i := 0.00001, 0

	for temp := 0.0; math.Abs(z-temp) > r; {
		temp = z
		i++
		z = z - (z*z-x)/(2*z)
	}
	fmt.Println("반복 횟수: ", i)
	return
}

func main() {
	target := 117.0
	myResult := mySqrt(target)
	fmt.Println("결과: ", myResult)
	fmt.Println("실제: ", math.Sqrt(target))
	fmt.Println("오차: ", myResult-math.Sqrt(target))
}
