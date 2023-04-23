package main

import "fmt"

func main() {
	var a, b, c float64
	var delta float64
	var x1, x2 float64

	fmt.Println("This program solves a quadratic equation of the form ax^2 + bx + c = 0")
	fmt.Println("Enter a:")
	fmt.Scan(&a)
	fmt.Println("Enter b:")
	fmt.Scan(&b)
	fmt.Println("Enter c:")
	fmt.Scan(&c)

	delta = b*b - 4*a*c
	fmt.Println("delta =", delta)

	if delta < 0 {
		fmt.Println("The equation has no real roots")
	} else if delta == 0 {
		fmt.Println("The equation has one real root")
		x1 = -b / (2 * a)
		fmt.Println("x1 =", x1)
	} else {
		fmt.Println("The equation has two real roots")
		x1 = (-b + (b*b-4*a*c)) / (2 * a)
		x2 = (-b - (b*b-4*a*c)) / (2 * a)

		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}

	
}