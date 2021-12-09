package aocget

import (
	"fmt"
	"os"
)

func ExampleAoCClient_DownloadInputString() {
	aoc := NewClient(os.Getenv("SESSION_TOKEN"))
	i, err := aoc.DownloadInputString(2021, 1)
	if err != nil {
		fmt.Printf("unable to download input: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(i)
}
