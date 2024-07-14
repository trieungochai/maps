package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// var colors map[string]string

	colors := make(map[int]string)

	colors[10] = "#ffffff"
	delete(colors, 10)

	fmt.Println(colors)
}
