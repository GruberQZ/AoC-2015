
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

    for index, char := range steps {
        cur := string(char)
        if cur == up {
            floor++
        } else if cur == down {
            floor--
        } else {
            fmt.Println("Unexpected character")
        }
        if floor < 0 {
            fmt.Printf("Position where basement first entered: %d\n", index+1)
            break
        }
    }

}
