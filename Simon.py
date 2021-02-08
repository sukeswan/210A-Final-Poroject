import time 
start_time = time.time()

import os, psutil
process = psutil.Process(os.getpid())

import copy
# Test Vector for Simon
plain_text = "74206e69206d6f6f6d69732061207369"
key = "1f1e1d1c1b1a191817161514131211100f0e0d0c0b0a09080706050403020100"
ciphertext  = "8d2b5579afc8a3a03bf72a87efe7b868"
AND = "AND"

z4 = [1,1,0,1,0,0,0,1,1,1,1,0,0,1,1,0,1,0,1,1,0,1,1,0,0,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,0,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,1,0,1,1,1,1]

# convert hex value to binary list 
def hex_to_binary(input,size):
    binary_input = bin(int(input,16))[2:].zfill(size)
    binary_list = list(map(int,str(binary_input)))
    return binary_list

# convert binary list to hex
def binary_to_hex(input):
    decimal_input = int("".join(str(i) for i in input),2)
    hex_value = hex(decimal_input)[2:]
    return hex_value

# shift left multiple times
def shift_left(input, times):
    result = copy.deepcopy(input)
    for i in range(times):
        newTail = result.pop(0)
        result.append(newTail)
    return result

#shift right function for genrating keys
def shift_right(input,times):
    result = copy.deepcopy(input)
    for i in range(times):
        newHead = result.pop()
        result.insert(0,newHead)
    return result

# split binary list into left and right
def split(input):
    half = len(input)//2
    left = input[:half]
    right = input[half:]
    return left,right

# combine left and right binary lists
def combine(left,right):
    return left + right

# return list of xor-ed bits in input lists
def bit_xor(one,two):
    result = []
    for i in range(len(one)):
        result.append(one[i] ^ two[i])
    return result

# return list of and-ed bits in input lists
def bit_and(one,two):
    result = []
    for i in range(len(one)):
        result.append(one[i] & two[i])
    return result

# invert the bits in k for key expansion 
def invert(input):
    result = copy.deepcopy(input)
    for i in range(len(result)):
        if result[i]==1:
            result[i]=0
        elif result[i]==0:
            result[i]=1
    return result 

# rounds for Simon
def round(left,right,key):
    left_shift_1 = shift_left(left,1)
    left_shift_2 = shift_left(left,2)
    left_shift_8 = shift_left(left,8)

    fx_res = bit_and(left_shift_1,left_shift_8)

    xor_1 = bit_xor(right,fx_res)
    xor_2 = bit_xor(xor_1,left_shift_2)
    xor_3 = bit_xor(xor_2,key)
    
    result_left = xor_3
    result_right  = left
    return result_left,result_right

def simon(binary_plain_text,sub_keys):
    left,right = split(binary_plain_text)
    for i in range(72):
        left,right = round(left,right,sub_keys[i])
    return combine(left,right)

# decrypt by switching left and right in final rounds
def simon_d(binary_plain_text,sub_keys):
    right,left = split(binary_plain_text)
    for i in range(72):
        left,right = round(left,right,sub_keys[i])
    return combine(right,left)

#generate subkeys
def generate_keys(key):
    high_order,low_order = split(key)
    k3,k2 = split(high_order)
    k1,k0 = split(low_order)

    sub_keys = [k0,k1,k2,k3]

    # print("Round {:02d} Key {}".format(0,binary_to_hex(k0)))
    # print("Round {:02d} Key {}".format(1,binary_to_hex(k1)))
    # print("Round {:02d} Key {}".format(2,binary_to_hex(k2)))
    # print("Round {:02d} Key {}".format(3,binary_to_hex(k3)))

    for i in range(4,72):
        temp = shift_right(sub_keys[i-1],3)
        temp = bit_xor(temp,sub_keys[i-3])
        temp = bit_xor(temp,shift_right(temp,1))
        
        invert_kim = invert(sub_keys[i-4])
        temp = bit_xor(invert_kim,temp)
        
        # xor 3 (11) with the two last bits in the key
        z_bit = z4[((i-4) % 62)]
        temp[-1] = temp[-1] ^ z_bit ^ 1
        temp[-2] = temp[-2] ^ 1
        sub_keys.append(temp)

        #print("Round {:02d} Key {}".format(i,binary_to_hex(temp)))
    
    return sub_keys

def main():

    binary_key = hex_to_binary(key,256)
    sub_keys = generate_keys(binary_key)

    # Encryption of Test Vector
    binary_plain_text = hex_to_binary(plain_text,128)
    binary_cipher_text = simon(binary_plain_text,sub_keys)
    final_encrypt = binary_to_hex(binary_cipher_text)
    check = ciphertext == final_encrypt
    print("Did Simon encrypt correctly? {}".format(check))

    # Decryption of Test Vector
    sub_keys.reverse()
    binary_cipher_text = hex_to_binary(ciphertext,128)
    binary_plain_text = simon_d(binary_cipher_text,sub_keys)
    final_decrypt = binary_to_hex(binary_plain_text)
    check = plain_text == final_decrypt
    print("Did Simon decrypt correctly? {}\n".format(check))

if __name__ == "__main__":
    main()

print("Simon in python took %s seconds" % (time.time() - start_time))
print("Simon used {} MB of memory".format(psutil.Process(os.getpid()).memory_info().rss / 1024 ** 2))