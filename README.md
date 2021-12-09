# Advent of Code Get (aocget)
A small lib to download your puzzle input for a given day.

Uses your session token to authenticate to obtain your personalized input.

## Getting your Session Token
1. Login to [Advent of Code](https://adventofcode.com)
2. Open your browser's DevTools (usually F12)
3. Head over to your cookies
4. Copy the `value` of the `session` cookie


## Example

```go
package main

import (
	"fmt"
	"github.com/Zyian/aocget"
	"os"
)

func main() {
	aoc := aocget.NewClient(os.Getenv("SESSION_TOKEN"))
	i, err := aoc.DownloadInputString(2021, 1)
	if err != nil {
        fmt.Printf("unable to download input: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(i)
}
```
