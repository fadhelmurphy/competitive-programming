// Contoh Soal Competitive Programming dengan Python, Golang, dan JavaScript
// Soal:

// Diberikan sebuah string yang terdiri dari karakter 0 dan 1. Hitunglah jumlah substring yang memiliki jumlah karakter 0 dan 1 yang sama.

// Contoh:

// Input: 1101
// Output: 3
// Penjelasan:

// Substring dengan jumlah 0 dan 1 sama: 11, 01, dan 101.

package main

import (
  "fmt"
)

func countSubstrings(s string) int {
  count := 0
  zeros := 0
  ones := 0
  prevChar := rune(' ') // Ganti byte menjadi rune
  for _, char := range s {
    if char == '0' {
      zeros++
    } else {
      ones++
    }
    if zeros <= ones && (prevChar != '1' || char != '0') {
      count++
    }
    prevChar = char
  }
  return count
}

func main() {
  s := "1101"
  fmt.Println(countSubstrings(s))
}
