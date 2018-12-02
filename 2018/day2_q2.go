package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    filename := "day2_input.csv"
    f, _ := os.Open(filename)
    defer f.Close()

    r := csv.NewReader(f)

    var id1 string
    var id2 string
    diff_pos := 0

    records, _ := r.ReadAll()
    Loop:
      for i, record := range records {
        for j := i+1; j < len(records); j++ {
          diff := 0
          for k := 0; k < len(record[0]); k++ {
            if record[0][k] != records[j][0][k] {
              diff ++
              diff_pos = k
            }
          }
          if diff == 1 {
            id1 = record[0]
            id2 = records[j][0]
            break Loop
          }
        }
      }

    same_str := id1[:diff_pos] + id1[diff_pos+1:]

    // fmt.Println(id1)
    // fmt.Println(id2)
    // fmt.Println(diff_pos)
    fmt.Println(same_str)

}
