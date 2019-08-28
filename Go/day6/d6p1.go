package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func getXYCoords(s string) (int, int) {
    a := strings.Split(s, ",")
    x, err := strconv.Atoi(a[0])
    if err != nil {
        panic(err)
    }
    y, err := strconv.Atoi(a[1])
    if err != nil {
        panic(err)
    }
    return x, y
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // Access grid with grid[y][x]
    // y0 is top, y999 is bottom
    // x0 is left, x999 is right
    grid := make([][1000]bool, 1000)

    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        var s []string = strings.Split(trimmed, " ")
        var x1, y1, x2, y2 int
        // Get coords
        if len(s) == 4 {
            // Toggle
            x1, y1 = getXYCoords(s[1])
            x2, y2 = getXYCoords(s[3])
        } else {
            // On/off
            x1, y1 = getXYCoords(s[2])
            x2, y2 = getXYCoords(s[4])
        }
        // Change lights
        for y := y1; y <= y2; y++ {
            for x := x1; x <= x2; x++ {
                if s[0] == "toggle" {
                    grid[y][x] = !grid[y][x]
                } else if s[1] == "on" {
                    grid[y][x] = true
                } else if s[1] == "off" {
                    grid[y][x] = false
                } else {
                    panic("Unexpected action")
                }
            }
        }

    }

    // Count
    lit := 0
    for y := 0; y <= 999; y++ {
        for x := 0; x <= 999; x++ {
            if grid[y][x] {
                lit++
            }
        }
    }

    fmt.Printf("Lights on: %d\n", lit)

}