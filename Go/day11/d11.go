package main

import (
    "fmt"
    "strings"
)

func hasStraight(s string) bool {
    for i := 0; i < len(s) - 2; i++ {
        if s[i] + 1 == s[i+1] && s[i] + 2 == s[i+2] {
            return true
        }
    }
    return false
}

func hasTwoNonOverlappingPairs(s string) bool {
    count := 0
    for i := 0; i < len(s) - 1; i++ {
        if s[i] == s[i+1] {
            count++
            if count == 2 {
                return true
            }
            // Extra increment of i for non-overlapping rule
            i++
        }
    }
    return false
}

func incrementPassword(s string) string {
    runes := []rune(s)
    for i := len(runes) - 1; i >= 0; i-- {
        if runes[i] < 'z' {
            runes[i]++
            break
        } else {
            runes[i] = 'a'
        }
    }
    return string(runes)
}

func findNextPassword(s string) string {
    nextPassword := s
    for {
        nextPassword = incrementPassword(nextPassword)
        if hasStraight(nextPassword) && !strings.ContainsAny(nextPassword, "iol") && hasTwoNonOverlappingPairs(nextPassword) {
            return nextPassword
        }
    }
}

func main() {
    input := "hepxcrrq"
    nextPassword := findNextPassword(input)
    fmt.Printf("NextPassword: %s\n", nextPassword)
    nextNextPassword := findNextPassword(nextPassword)
    fmt.Printf("NextNextPassword: %s\n", nextNextPassword)
}