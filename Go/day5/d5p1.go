package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func hasThreeVowels(s string) bool {
	total := 0
	total += strings.Count(s, "a")
	total += strings.Count(s, "e")
	total += strings.Count(s, "i")
	total += strings.Count(s, "o")
	total += strings.Count(s, "u")
	return total >= 3
}

func hasOneDouble(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasNoBadSubstrings(s string) bool {
	return (!strings.Contains(s, "ab") && 
		!strings.Contains(s, "cd") && 
		!strings.Contains(s, "pq") && 
		!strings.Contains(s, "xy"))
}

func main() {
	// Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    niceStrings := 0

    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        if (hasThreeVowels(trimmed) && hasOneDouble(trimmed) && hasNoBadSubstrings(trimmed)) {
        	niceStrings++
        }
    }

    fmt.Printf("Nice strings: %d\n", niceStrings)

}