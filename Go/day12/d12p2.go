package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
    "bytes"
)

func removeRed(start []byte) []byte {
    red := []byte(`:"red"`)
    b := start
    curr := 0
    redsLeft := bytes.Count(b, red)
    for redsLeft > 0 {
        i := bytes.Index(b, red)
        // Go backward looking for brackets
        openingToFind := 1
        curr = i
        for openingToFind != 0 {
            curr--
            if b[curr] == '{' {
                openingToFind--
            } else if b[curr] == '}' {
                openingToFind++
            }
        }
        start := curr
        // Go forward looking for brackets
        closingToFind := 1
        curr = i
        for closingToFind != 0 {
            curr++
            if b[curr] == '}' {
                closingToFind--
            } else if b[curr] == '{' {
                closingToFind++
            }
        }
        end := curr
        b = append(b[:start], b[end+1:]...)
        redsLeft = bytes.Count(b, red)
    }
    return b
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    regex := regexp.MustCompile(`[\-\d]+`)
    total := 0

    for scanner.Scan() {
        // Get line text
        var text string = scanner.Text()
        var trimmed string = strings.TrimSpace(text)
        var trimmedBytes = []byte(trimmed)
        // Remove red
        nored := removeRed(trimmedBytes)
        // Sum
        r := regex.FindAll(nored, -1)
        for _, m := range r {
            val, err := strconv.Atoi(string(m))
            if err != nil {
                panic(err)
            }
            total += val
        }
    }

    fmt.Printf("Total: %d\n", total)

}