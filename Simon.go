package main

import (
	"fmt"
  "strconv"
  // "strings"
)

func split(input []int) ([]int, []int){
  half := len(input)/2
  slc1 := make([]int, half)
  copy(slc1, input[:half])
  slc2 := make([]int, half)
  copy(slc2, input[half:])
  return slc1, slc2
}
func shift_left(input []int, times int) []int{
  result := make([]int, len(input))
  copy(result, input)
  for i := 0; i < times; i++ {
    newTail  := result[0]
    result = result[1:len(result)]
    result = append( result, newTail)
  }
  return result
}
func shift_right(input []int, times int) []int{
  result := make([]int, len(input))
  copy(result, input)
  for i := 0; i < times; i++ {
    newHead  := result[len(result)-1]
    result = result[:len(result)-1]
    result = append([]int{newHead}, result...)
  }
  return result
}
func bit_xor(one []int, two []int)[]int{
  result := make([]int, len(one))
  for i := 0; i < len(one); i++ {
    result[i] = (one[i] ^ two[i])
  }
  return result
}
func invert(input []int) []int{
  result := make([]int, len(input))
  copy(result, input)
  for i, b := range result{
    if b == 1{
      result[i] = 0
    }else{
      result[i] = 1
    }
  }
  return result
}

func generate_keys(input []int, z4 []int) ([][]int){
  sub_keys := make([][]int, 72)
  high_order, low_order := split(input)
  k3,k2 := split(high_order)
  k1,k0 := split(low_order)
  sub_keys[0], sub_keys[1], sub_keys[2], sub_keys[3] = k0,k1,k2,k3
  for i := 4; i < 72; i++ {
    temp_s3 := shift_right(sub_keys[i-1],3)
    temp_x1 := bit_xor(temp_s3, sub_keys[i-3])
    temp_x2 := bit_xor(temp_x1, shift_right(temp_x1,1))
    invert_kim := invert(sub_keys[i-4])
    temp_x3 := bit_xor(invert_kim,temp_x2)

    z_bit := z4[((i-4) % 62)]
    temp_x3[len(temp_x3)-1] = temp_x3[len(temp_x3)-1] ^ z_bit ^ 1
    temp_x3[len(temp_x3)-2] = temp_x3[len(temp_x3)-2] ^ 1
    sub_keys[i] = temp_x3
  }

  return sub_keys
}

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
	z4 := []int{1,1,0,1,0,0,0,1,1,1,1,0,0,1,1,0,1,0,1,1,0,1,1,0,0,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,0,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,1,0,1,1,1,1}
  binary_key := hex_to_binary(key)
  // fmt.Printf("%v\n", binary_key)
  sub_keys := generate_keys(binary_key, z4)
  fmt.Printf("%v\n", sub_keys[len(sub_keys)-1])
  // test := []int{1,3,5,7,9,11}
  // fmt.Printf("%v\n", test)
  // test = shift_left(test, 3)
  // fmt.Printf("%v\n", test)
  // test := invert([]int{1,1,0,1})
  // fmt.Printf("%v\n", test)


}


















//
