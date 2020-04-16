package FunctionsMethodsInterfaces

import (
	"fmt"
	"math"
)

var funcVar func(int) int

func mai1n() {
	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64

	fmt.Print("Please Enter Acceleration: ")
	fmt.Scan(&acceleration)
	fmt.Println()
	fmt.Print("Please Enter initial Velocity: ")
	fmt.Scan(&initialVelocity)
	fmt.Println()
	fmt.Print("Please Enter initial Displacement: ")
	fmt.Scan(&initialDisplacement)

	fn := GenDisplayFn(10, 2, 1)
	fmt.Print("Please Enter Time: ")
	fmt.Scan(&time)
	fmt.Printf("The Final Displacement after %v seconds is %v", (time), fn(time))
	// fmt.Println(fn(time))
}

func GenDisplayFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*a*math.Pow(time, 2) + v0*time + s0
	}

	return fn
}
