package main

import (
	"fmt"
  "strconv"
  // "strings"
)


func hex_to_binary(input string) []int{
  var bits []int
  bitString := ""
  loop_range := len(input)/16

  for i := 0; i < loop_range; i++ {
    idx := i*16
    bit_ui, _ := strconv.ParseUint(input[idx:idx+16], 16, 64)
    bitString = bitString + fmt.Sprintf("%064b", bit_ui)
  }
  for _ , bit := range bitString {
    b, _ := strconv.Atoi(string(bit))
		bits = append(bits, b)
	}

  return bits
}


func main() {
  // plain_text := "74206e69206d6f6f6d69732061207369"
  key := "1f1e1d1c1b1a191817161514131211100f0e0d0c0b0a09080706050403020100" //
  // ciphertext  := "8d2b5579afc8a3a03bf72a87efe7b868"
	// z4 := []int{1,1,0,1,0,0,0,1,1,1,1,0,0,1,1,0,1,0,1,1,0,1,1,0,0,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,0,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,1,0,1,1,1,1}
  binary_key := hex_to_binary(key)
  fmt.Printf("%v\n", binary_key)

}


















//
