
package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
    // Get input
    steps, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        panic(err)
    }

    floor := 0

    up := string("(")
    down := string(")")

    for _, char := range steps {
        cur := string(char)
        if cur == up {
            floor++
        } else if cur == down {
            floor--
        } else {
            fmt.Println("Unexpected character")
        }
    }

    fmt.Printf("Final floor: %d\n", floor)

}