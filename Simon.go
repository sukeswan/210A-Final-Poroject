package main

import (
	"fmt"
  "strconv"
  "strings"
)


func hex_to_binary(input string, size string) string{
  // bits := []int{}
  bit_ui, err := strconv.ParseUint(input, 16, 64)
  if err != nil {
		return "wtf"
	}

  b_input := strings.Join([]string{"%0", size , "b"}, "")
  //iterate through string every 16 digits
  //convert and add to bit string
  //for s in string slice append return int slice



  // for i, c := range input{
  //
  // }
  // for i := 0; i < 24; i++ {
  //       bits = append([]int{val & 0x1}, bits...)
  //       // or
  //       // bits = append(bits, val & 0x1)
  //       // depending on the order you want
  //       val = val >> 1
  //   }
  bitString := fmt.Sprintf(b_input, bit_ui)
  return bitString
  // return bits
}


func main() {
  // plain_text := "74206e69206d6f6f6d69732061207369"
  key := "1f1e1d1c1b1a191817161514131211100f0e0d0c0b0a09080706050403020100"
  // ciphertext  := "8d2b5579afc8a3a03bf72a87efe7b868"
	// z4 := []int{1,1,0,1,0,0,0,1,1,1,1,0,0,1,1,0,1,0,1,1,0,1,1,0,0,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,0,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,1,0,1,1,1,1}
  binary_key := hex_to_binary(key,"256")
  fmt.Printf(binary_key)
  // s.constructor("E", slc)
	// fmt.Printf("en %s, %s", s.en_de,s.key )
}
