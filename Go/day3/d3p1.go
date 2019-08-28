package main

import (
    "bufio"
    "fmt"
    "os"
)

type Coord struct {
    x, y int
}

type Void struct {}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    if !scanner.Scan() {
        panic("Scanner didn't find anything")
    }
    var sequence string = scanner.Text()

    visited := make(map[Coord]Void)
    santaCoord := Coord{
        x: 0,
        y: 0,
    }

    // Deliver present at starting location
    visited[santaCoord] = Void{}

    for _, char := range sequence {
        // Move
        if char == '^' {
            santaCoord.y += 1
        } else if char == 'v' {
            santaCoord.y -= 1
        } else if char == '<' {
            santaCoord.x -= 1
        } else if char == '>' {
            santaCoord.x += 1
        }
        // Check to see if visited this coord already
        _, alreadyVisited := visited[santaCoord]
        if !alreadyVisited {
            visited[santaCoord] = Void{}
        }

    }

    fmt.Printf("Visited houses: %d\n", len(visited))

}