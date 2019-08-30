package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Instruction struct {
    in1 string
    in2 string
    op string
    out string
}

func NewInstruction(s []string) Instruction {
    i := Instruction {}
    lenS := len(s)
    if lenS == 3 {
        i.in1 = s[0]
        i.in2 = ""
        i.op = ""
        i.out = s[2]
    } else if lenS == 4 {
        i.in1 = s[1]
        i.in2 = ""
        i.op = s[0]
        i.out = s[3]
    } else {
        i.in1 = s[0]
        i.in2 = s[2]
        i.op = s[1]
        i.out = s[4]
    }
    return i
}

// Try to evaluate instruction
func EvalInstruction(i Instruction, m map[string]uint16) (uint16, bool) {
    var in1val uint16
    var in2val uint16
    // Check first input
    if val, ok := m[i.in1]; ok {
        in1val = val
    } else {
        val, err := strconv.Atoi(i.in1)
        // Check if first input is a number
        if err != nil {
            return 0, false
        } else {
            in1val = uint16(val)
        }
    }
    // Check second input
    if i.in2 == "" {
        // Do nothing, won't be used
    } else if val, ok := m[i.in2]; ok {
        in2val = val
    } else {
        val, err := strconv.Atoi(i.in2)
        // Check if second input is a number
        if err != nil {
            return 0, false
        } else {
            in2val = uint16(val)
        }
    }
    // Evaluate
    var result uint16
    switch i.op {
    case "":
        result = in1val
    case "NOT":
        result = ^in1val
    case "AND":
        result = in1val & in2val
    case "OR":
        result = in1val | in2val
    case "LSHIFT":
        result = in1val << in2val
    case "RSHIFT":
        result = in1val >> in2val
    default:
        panic("Unexpected op: " + i.op)
    }


    return result, true
}

func main() {
    // Get input
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var instructions []Instruction

    for scanner.Scan() {
        // Get line text
        var line string = scanner.Text()
        var trimmed string = strings.TrimSpace(line)
        var s []string = strings.Split(trimmed, " ")
        instructions = append(instructions, NewInstruction(s))
    }

    known := make(map[string]uint16)

    for len(instructions) > 0 {
        var tempInstructions []Instruction
        for i := 0; i < len(instructions); i++ {
            inst := instructions[i]
            result, valid := EvalInstruction(inst, known)
            if valid {
                known[inst.out] = result
            } else {
                tempInstructions = append(tempInstructions, inst)
            }
        }
        instructions = tempInstructions
    }

    fmt.Printf("Value on wire a: %d\n", known["a"])

}