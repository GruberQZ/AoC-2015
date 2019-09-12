package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
)

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
        r := regex.FindAll(trimmedBytes, -1)
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