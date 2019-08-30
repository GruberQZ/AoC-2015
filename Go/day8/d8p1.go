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
    s := str[1:len(str)-1]
    length := len(s)
    count := length
    for i := 0; i < length; i++ {
        if s[i] == '\\' {
            if i+1 < length && (s[i+1] == '"' || s[i+1] == '\\') {
                count -= 1
                i += 1
            } else if i+3 < length && regex.Match([]byte(s[i:i+4])) {
                count -= 3
                i += 3
            }
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
    inMemoryLength := 0

    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        literalLength += len(trimmed)
        inMemoryLength += calc(trimmed)        
    }

    fmt.Printf("Difference: %d\n", literalLength-inMemoryLength)

}