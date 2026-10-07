package main

import "fmt"

func main() {
	const freezingWater = 32.0
	celsius := 25.0
	fahrenheit := (celsius * 9 / 5) + freezingWater
	fmt.Printf("%.1f C° is %.1f F°\n", celsius, fahrenheit)

}
