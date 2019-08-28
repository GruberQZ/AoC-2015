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
    roboCoord := Coord{
        x: 0,
        y: 0,
    }

    pSantaCoord := &santaCoord
    pRoboCoord := &roboCoord

    // Deliver present at starting location
    visited[santaCoord] = Void{}

    var pMover *Coord
    for i, char := range sequence {
        // Alternate reference to who is moving this turn
        if i % 2 == 0 {
            pMover = pSantaCoord
        } else {
            pMover = pRoboCoord
        }
        // Move
        if char == '^' {
            (*pMover).y += 1
        } else if char == 'v' {
            (*pMover).y -= 1
        } else if char == '<' {
            (*pMover).x -= 1
        } else if char == '>' {
            (*pMover).x += 1
        }
        // Check to see if visited this coord already
        _, alreadyVisited := visited[(*pMover)]
        if !alreadyVisited {
            visited[(*pMover)] = Void{}
        }

    }

    fmt.Printf("Visited houses: %d\n", len(visited))

}