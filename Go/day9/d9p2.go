package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Void struct {}

func findLongest(distances map[string]map[string]int, prevCity string, destinations []string, currentDistance int) int {
    longest := 0
    if len(destinations) == 1 {
        return currentDistance + distances[prevCity][destinations[0]]
    } else {
        for _, currentCity := range destinations {
            otherCities := make([]string, 0, len(destinations)-1)
            for _, c := range destinations {
                if c != currentCity {
                    otherCities = append(otherCities, c)
                }
            }
            var tmp int = 0
            if prevCity != "" {
                tmp = findLongest(distances, currentCity, otherCities, currentDistance + distances[prevCity][currentCity])
            } else {
                tmp = findLongest(distances, currentCity, otherCities, 0)
            }
            if tmp > longest {
                longest = tmp
            }
        }
        return longest
    }
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    distances := make(map[string]map[string]int)
    cities := make(map[string]Void)

    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        var s []string = strings.Split(trimmed, " ")
        d, err := strconv.Atoi(s[4])
        if err != nil {
            panic(err)
        }
        if _, ok := distances[s[0]]; !ok {
            distances[s[0]] = make(map[string]int)
        }
        if _, ok := distances[s[2]]; !ok {
            distances[s[2]] = make(map[string]int)
        }
        distances[s[0]][s[2]] = d
        distances[s[2]][s[0]] = d
        cities[s[0]] = Void{}
        cities[s[2]] = Void{}
    }

    destinations := make([]string, 0, len(cities))
    for k := range cities {
        destinations = append(destinations, k)
    }

    // for source, m := range distances {
    //     for dest, dist := range m {
    //         fmt.Printf("Distance from %s to %s = %d\n", source, dest, dist)
    //     }
    // }

    fmt.Printf("Longest: %d\n", findLongest(distances, "", destinations, 0))

}