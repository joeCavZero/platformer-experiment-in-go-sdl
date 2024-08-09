package main

import (
	"fmt"
	"os"
)

func main() {
	var tilemap []uint8 = make([]uint8, 0)
	numbers := map[string]uint8{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	data, err := os.ReadFile("data/level.data")
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers["60"])
	for _, v := range data {
		tilemap = append(tilemap, numbers[string(v)])
	}

	fmt.Println(tilemap)
}
