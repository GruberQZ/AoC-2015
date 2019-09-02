package main

import (
    "fmt"
    "strconv"
)

func lookAndSay(s string) string {
    output := ""
    currentChar := rune(s[0])
    count := 0
    for _, c := range s {
        if c == currentChar {
            count++
        } else {
            if count > 0 {
                output = output + strconv.Itoa(count) + string(currentChar)
            }
            currentChar = c
            count = 1
        }
    }
    if count > 0 {
        output = output + strconv.Itoa(count) + string(currentChar)
    }
    return output
}

func main() {
    
    input := "1113122113"

    for i := 0; i < 50; i++ {
        input = lookAndSay(input)
    }

    fmt.Printf("Length: %d\n", len(input))

}