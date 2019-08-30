package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "regexp"
)

var regex = regexp.MustCompile(`\\x[\da-f]{2}`)

func calc(str string) int {
    s := str
    length := len(s)
    // Add 2 for new wrapping quotation marks
    count := length + 2
    for i := 0; i < length; i++ {
        if i+3 < length && regex.Match([]byte(s[i:i+4])) {
            count += 1
            i += 3
        } else if s[i] == '"' || s[i] == '\\' {
            count += 1
        }
    }
    return count
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    literalLength := 0
    encodedLength := 0

    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        literalLength += len(trimmed)
        encodedLength += calc(trimmed)        
    }

    fmt.Printf("Difference: %d\n", encodedLength-literalLength)

}