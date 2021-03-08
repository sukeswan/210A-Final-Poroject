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
func bit_and(one []int, two []int)[]int{
  result := make([]int, len(one))
  for i := 0; i < len(one); i++ {
    result[i] = (one[i] & two[i])
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

func round(left,right,key []int) ([]int,[]int){
  left_shift_1 := shift_left(left,1)
  left_shift_2 := shift_left(left,2)
  left_shift_8 := shift_left(left,8)

  fx_res := bit_and(left_shift_1,left_shift_8)

  xor_1 := bit_xor(right,fx_res)
  xor_2 := bit_xor(xor_1,left_shift_2)
  xor_3 := bit_xor(xor_2,key)

  result_left := xor_3
  result_right  := left
  return result_left, result_right

}

func simon(binary_plain_text []int, sub_keys [][]int) []int{
  left,right := split(binary_plain_text)
  for i := 0; i < 72; i++ {
    left,right = round(left, right, sub_keys[i])
  }
  return append(left, right...)
}
func simon_d(binary_plain_text []int, sub_keys [][]int) []int{
  right,left := split(binary_plain_text)
  for i := 0; i < 72; i++ {
    left,right = round(left, right, sub_keys[i])
  }
  return append(right, left...)
}

func binary_to_hex(input []int) string{

  binaryString := ""
  for _ , bit := range input {
    b := strconv.Itoa(int(bit))
    binaryString = binaryString + b
  }

  hexString := ""
  loop_range := len(input)/4

  for i := 0; i < loop_range; i++ {
    idx := i*4
    bit_ui, _ := strconv.ParseUint(binaryString[idx:idx+4], 2, 64)
    hexString = hexString + fmt.Sprintf("%x", bit_ui)
  }


  return hexString
  // ui, err := strconv.ParseUint(s, 2, 64)
  //   if err != nil {
  //       return "error"
  //   }
}

func main() {
  plain_text := "74206e69206d6f6f6d69732061207369"
  key := "1f1e1d1c1b1a191817161514131211100f0e0d0c0b0a09080706050403020100" //
  ciphertext  := "8d2b5579afc8a3a03bf72a87efe7b868"
	z4 := []int{1,1,0,1,0,0,0,1,1,1,1,0,0,1,1,0,1,0,1,1,0,1,1,0,0,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,0,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,1,0,1,1,1,1}
  binary_key := hex_to_binary(key)
  // fmt.Printf("%v\n", binary_key)
  sub_keys := generate_keys(binary_key, z4)
  // fmt.Printf("%v\n", sub_keys[len(sub_keys)-1])
  binary_plain_text := hex_to_binary(plain_text)
  // test := []int{1,3,5,7,9,11}
  // fmt.Printf("%v\n", binary_plain_text)
  binary_cipher_text := simon(binary_plain_text,sub_keys)
  // fmt.Printf("%v\n", binary_cipher_text)
  final_encrypt := binary_to_hex(binary_cipher_text)
  // test = shift_left(test, 3)
  encrypt_check := ciphertext == final_encrypt
  fmt.Printf("%v\n", encrypt_check)
  // test := bit_and([]int{1,0,0,1}, []int{0,1,0,1})
  // fmt.Printf("%v\n", test)
  for i, j := 0, len(sub_keys)-1; i < j; i, j = i+1, j-1 {
       sub_keys[i], sub_keys[j] = sub_keys[j], sub_keys[i]
   }
// fmt.Printf("%v\n", sub_keys[len(sub_keys)-1])
  recovered_plain_text := simon_d(binary_cipher_text, sub_keys)
  final_decrypt := binary_to_hex(recovered_plain_text)
  fmt.Printf("%s\n", final_decrypt)

}


















//
