package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Reindeer struct {
    pos, speed, phaseTime, runTime, restTime, points int
    resting bool
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var reindeer []*Reindeer

    for scanner.Scan() {
        // Get line text
        var text string = scanner.Text()
        var trimmed string = strings.TrimSpace(text)
        var words []string = strings.Split(trimmed, " ")
        // Parse
        speed, _ := strconv.Atoi(words[3])
        runTime, _ := strconv.Atoi(words[6])
        restTime, _ := strconv.Atoi(words[13])
        // Make reindeer
        r := &Reindeer{
            pos: 0,
            speed: speed,
            runTime: runTime,
            restTime: restTime,
            phaseTime: runTime,
            points: 0,
            resting: false,
        }
        // Add reindeer
        reindeer = append(reindeer, r)
    }

    for i := 0; i < 2503; i++ {
        // Tick each reindeer
        maxDistance := 0
        for _, r := range reindeer {
            // Tick
            r.phaseTime--
            // Advance if not resting
            if !r.resting {
                r.pos += r.speed
            }
            // Look for max distance
            if r.pos > maxDistance {
                maxDistance = r.pos
            }
            // Run/rest transition
            if r.phaseTime == 0 {
                if r.resting {
                    r.phaseTime = r.runTime
                    r.resting = false
                } else {
                    r.phaseTime = r.restTime
                    r.resting = true
                }
            }
        }
        // Award points
        for _, r := range reindeer {
            if r.pos == maxDistance {
                r.points++
            }
        }
    }

    // Find max points
    maxPoints := 0
    for _, r := range reindeer {
        if r.points > maxPoints {
            maxPoints = r.points
        }
    }

    fmt.Printf("Maximum points: %d\n", maxPoints)

}