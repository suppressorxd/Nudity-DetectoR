package main

import (
	"fmt"

	"github.com/koyachi/go-nude"
)

func main() {
	// List of pictures to be checked for nudity
	pictures := []string{
		"1.png",
	}

	// Loop through each picture in the list
	for _, picture := range pictures {
		// Check if the picture contains nudity
		isNude, _ := nude.IsNude(picture)

		// Print the result of nudity check along with the picture name
		fmt.Println(isNude, "t", picture)
	}
}
