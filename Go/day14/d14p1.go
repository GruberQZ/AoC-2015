package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
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

    maxDistance := 0

    for scanner.Scan() {
        // Get line text
        var text string = scanner.Text()
        var trimmed string = strings.TrimSpace(text)
        var words []string = strings.Split(trimmed, " ")
        // Parse
        speed, _ := strconv.Atoi(words[3])
        runTime, _ := strconv.Atoi(words[6])
        restTime, _ := strconv.Atoi(words[13])
        // Run
        time := 2503
        distance := 0
        for time > 0 {
            if time >= runTime {
                distance += runTime * speed
                time -= (runTime + restTime)
            } else {
                distance += time * speed
                time = 0
            }
        }
        if distance > maxDistance {
            maxDistance = distance
        }
    }

    fmt.Printf("Maximum distance: %d\n", maxDistance)

}