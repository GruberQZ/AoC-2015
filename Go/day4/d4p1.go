package main

import (
    "fmt"
    "crypto/md5"
    "strconv"
    "strings"
    "encoding/hex"
)

type Coord struct {
    x, y int
}

type Void struct {}

func main() {
    // Get input
    input := "iwrupvqb"

    number := 1

    fiveZeros := "00000"

    for {
        numStr := strconv.Itoa(number)
        data := []byte(input + numStr)
        sum := md5.Sum(data)
        encoded := hex.EncodeToString(sum[:])
        // fmt.Printf("%s %s\n", string(data), encoded)
        if strings.HasPrefix(encoded, fiveZeros) {
            break
        }
        number++
    }

    fmt.Printf("Number: %d\n", number)

}