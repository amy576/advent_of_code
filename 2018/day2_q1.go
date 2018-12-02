package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
)

func main() {
    filename := "day2_input.csv"
    f, _ := os.Open(filename)
    defer f.Close()

    r := csv.NewReader(f)
    boxes_with_two := 0
    boxes_with_three := 0

    records, _ := r.ReadAll()
    for _, record := range records {
        for _, field := range record {
            seen := make(map[rune]int)
            for _, c := range field {
                seen[c] = strings.Count(field,string(c))
            }
            TwoLoop:
              for r, _ := range seen {
                  if seen[r] == 2 {
                    boxes_with_two += 1
                    break TwoLoop
                  }
              }
            ThreeLoop:
              for r, _ := range seen {
                  if seen[r] == 3 {
                    boxes_with_three += 1
                    break ThreeLoop
                  }
              }
        }
    }

    checksum := boxes_with_two * boxes_with_three

    fmt.Println(checksum)
}
