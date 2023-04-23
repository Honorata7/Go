package main

import "fmt"

func main() {
	var age int
	var planet string

	var planets = map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	fmt.Println("Enter the name of a planet:")
	fmt.Scan(&planet)

	planetAge := float64(age) / planets[planet]

	fmt.Println("Your age on", planet, "is", planetAge)

}