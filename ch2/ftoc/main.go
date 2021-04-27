package main

import "fmt"

const (
	boilingF = 212.0
	symbolF  = "℉"
	symbolC  = "℃"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g%s = %g%s\n", freezingF, symbolF, fToC(freezingF), symbolC) // "32°F = 0°C"
	fmt.Printf("%g%s = %g%s\n", boilingF, symbolF, fToC(boilingF), symbolC)   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
