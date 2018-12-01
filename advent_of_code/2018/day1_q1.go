package main

import (
    "encoding/csv"
    "strconv"
    "fmt"
    "os"
)

func main() {
    filename := "day1_input.csv"
    f, _ := os.Open(filename)
    defer f.Close()

    r := csv.NewReader(f)
    val := 0

    records, _ := r.ReadAll()
    for _, record := range records {
        for _, field := range record {
            i, _ := strconv.Atoi(field)
            val = val + i
        }
    }

    fmt.Printf("%v",val)
}
