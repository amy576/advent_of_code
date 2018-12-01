package main

import (
    "encoding/csv"
    "strconv"
    "fmt"
    "os"
)

func main() {
    val := 0
    frequencies := make(map[int]int)
    frequencies[val] = 1

    Loop:
      for frequencies[val] < 2 {
          filename := "day1_input.csv"
          f, _ := os.Open(filename)
          defer f.Close()

          r := csv.NewReader(f)

          records, _ := r.ReadAll()
          for _, record := range records {
              for _, field := range record {
                  i, _ := strconv.Atoi(field)
                  val = val + i
                  frequencies[val]++
                  if frequencies[val] == 2 {
                      fmt.Printf("%v",val)
                      break Loop
                  }
              }
          }
      }
}
