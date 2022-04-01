package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#fff1111",
	}

	// M2
	// var colors2 map[string]string
	// colors2["white"] = "#ffffff"

	// M3
	// colors3 := make(map[string]string)
	// colors3["black"] = "#000000"

	delete(colors, "red")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
