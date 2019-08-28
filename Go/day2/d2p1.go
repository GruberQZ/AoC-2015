
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func min(a int, b int) int {
    if a <= b {
        return a
    } else {
        return b
    }
}

func calcPaper(l int, w int, h int) int {
    var lw int = l*w
    var lh int = l*h
    var wh int = w*h
    var tmp int = min(lw, lh)
    var smallestSide int = min(tmp, wh)
    var total int = 2*lw + 2*lh + 2*wh + smallestSide
    return total
}

func convDims(dims []string) (int, int, int) {
    l, err := strconv.Atoi(dims[0])
    if err != nil {
        panic(err)
    }
    w, err := strconv.Atoi(dims[1])
    if err != nil {
        panic(err)
    }
    h, err := strconv.Atoi(dims[2])
    if err != nil {
        panic(err)
    }
    return l, w, h
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var total int = 0

    // Scan line by line
    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        var dims []string = strings.Split(trimmed, "x")
        l, w, h := convDims(dims)
        total += calcPaper(l, w, h)
    }

    fmt.Printf("Total square feet of paper: %d\n", total)

}