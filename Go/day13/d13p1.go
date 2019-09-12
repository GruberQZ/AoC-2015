package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func evaluateTable(order []string, happiness map[string]map[string]int) int {
    total := 0
    lastPos := len(order)-1
    for i, guest := range order {
        var left string
        var right string
        // Left
        if i == 0 {
            left = order[lastPos]
        } else {
            left = order[i-1]
        }
        // Right
        if i == lastPos {
            right = order[0]
        } else {
            right = order[i+1]
        }
        // Calc
        total += happiness[guest][left]
        total += happiness[guest][right]
    }
    return total
}

func findHappiest(seated []string, standing []string, happiness map[string]map[string]int) int {
    happiest := -1000
    if len(standing) == 1 {
        return evaluateTable(append(seated, standing[0]), happiness)
    } else {
        for i, _ := range standing {
            // Copy arrays because of weird append() things
            oldSeated := make([]string, len(seated))
            copy(oldSeated, seated)
            oldStanding := make([]string, len(standing))
            copy(oldStanding, standing)
            // Next arrangement
            newSeated := append(oldSeated, oldStanding[i])
            newStanding := append(oldStanding[:i], oldStanding[i+1:]...)
            tmp := findHappiest(newSeated, newStanding, happiness)
            if tmp > happiest {
                happiest = tmp
            }
        }
    }
    return happiest
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    happiness := make(map[string]map[string]int)
    var guestlist []string

    for scanner.Scan() {
        // Get line text
        var text string = scanner.Text()
        var trimmed string = strings.TrimSpace(text)
        var words []string = strings.Split(trimmed, " ")
        // Parse
        guest1 := words[0]
        guest2 := string(words[10][:len(words[10])-1])
        val, err := strconv.Atoi(words[3])
        if err != nil {
            panic(err)
        }
        if words[2] == "lose" {
            val = -val
        }
        // Build guest list
        if _, ok := happiness[guest1]; !ok {
            happiness[guest1] = make(map[string]int)
            guestlist = append(guestlist, guest1)
        }
        happiness[guest1][guest2] = val
    }

    emptylist := make([]string, 0)
    // Everyone standing, no one seated at start
    h := findHappiest(emptylist, guestlist, happiness)

    fmt.Printf("Maximum happiness: %d\n", h)

}