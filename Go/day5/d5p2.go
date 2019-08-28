package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func hasPairAppearingTwice(s string) bool {
    lenS := len(s)
    for i := 0; i < lenS-3; i++ {
        for j := i+2; j < lenS-1; j++ {
            if s[i:i+2] == s[j:j+2] {
                return true
            }
        }
    }
    return false
}

func hasRepeater(s string) bool {
    for i := 0; i < len(s)-2; i++ {
        if s[i] == s[i+2] {
            return true
        }
    }
    return false
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
        if (hasPairAppearingTwice(trimmed) && hasRepeater(trimmed)) {
            niceStrings++
        }
    }

    fmt.Printf("Nice strings: %d\n", niceStrings)

}