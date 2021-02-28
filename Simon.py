import time 
from statistics import mean 
import os, psutil
process = psutil.Process(os.getpid())

import copy, decimal
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

    for i in range(4,72):
        temp = shift_right(sub_keys[i-1],3)
        temp = bit_xor(temp,sub_keys[i-3])
        temp = bit_xor(temp,shift_right(temp,1))
        
        invert_kim = invert(sub_keys[i-4])
        temp = bit_xor(invert_kim,temp)
        
        z_bit = z4[((i-4) % 62)]            # xor 3 (11) with the two last bits in the key
        temp[-1] = temp[-1] ^ z_bit ^ 1
        temp[-2] = temp[-2] ^ 1
        sub_keys.append(temp)
    
    return sub_keys

def main():
    start_time = time.time()
    binary_key = hex_to_binary(key,256)
    sub_keys = generate_keys(binary_key)
    # Encryption of Test Vector
    binary_plain_text = hex_to_binary(plain_text,128)
    binary_cipher_text = simon(binary_plain_text,sub_keys)
    final_encrypt = binary_to_hex(binary_cipher_text)
    encrypt_check = ciphertext == final_encrypt

    # Decryption of Test Vector
    sub_keys.reverse()
    binary_cipher_text = hex_to_binary(ciphertext,128)
    binary_plain_text = simon_d(binary_cipher_text,sub_keys)
    final_decrypt = binary_to_hex(binary_plain_text)
    decrypt_check = plain_text == final_decrypt
    
    net_time = time.time() - start_time
    memory = psutil.Process(os.getpid()).memory_info().rss / 1024 ** 2
    cpu_stats = psutil.cpu_times(percpu=False)
    return encrypt_check,decrypt_check,net_time,memory,decimal.Decimal(cpu_stats[0]),decimal.Decimal(cpu_stats[2])
    
if __name__ == "__main__":
    cpu_stats = psutil.cpu_times(percpu=False)

    init_user = decimal.Decimal(cpu_stats[0])
    init_system = decimal.Decimal(cpu_stats[2])

    times = []
    mems = []
    cpus_user = []
    cpus_system = []

    cpus_user.append(init_user)
    cpus_system.append(init_system)

    for i in range(10):
        e_check,d_check,net_time,memory,cpu_user,cpu_system = main()
        times.append(net_time)
        mems.append(memory)
        cpus_user.append(cpu_user)
        cpus_system.append(cpu_system)

avg_time = mean(times)
avg_space = mean(mems)
avg_cpu_user = mean([cpus_user[i + 1] - cpus_user[i] for i in range(len(cpus_user)-1)])
avg_cpu_system = mean([cpus_system[i + 1] - cpus_system[i] for i in range(len(cpus_system)-1)])

print("--- Results of Simon in Python --- \n ")
print("Did Simon encrypt correctly? {}".format(e_check))
print("Did Simon decrypt correctly? {}".format(d_check))
print("Simon took an average of {} seconds over 10 runs".format(avg_time))
print("Simon used an average of {} MB of memory over 10 runs".format(avg_space))
print("Simon spent an average of {} seconds in the user CPU over 10 runs".format(avg_cpu_user))
print("Simon spent an average of {} seconds in the system CPU over 10 runs\n".format(avg_cpu_system))

print("Times:      {}".format(times))
print("Space:      {}".format(mems))
print("CPU User:   {}".format([str(i-init_user) for i in cpus_user]))     # user: time spent by normal processes executing in user modest time
print("CPU System: {}\n".format([str(i-init_system) for i in cpus_system]))   # system: time spent by processes executing in kernel mode
