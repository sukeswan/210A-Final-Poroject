const { cpus } = require("os");
var performance = require("performance-now");
var process = require('process'); 
// Test Vector for Simon
const plaintext = "74206e69206d6f6f6d69732061207369"
const key = "1f1e1d1c1b1a191817161514131211100f0e0d0c0b0a09080706050403020100"
const ciphertext  = "8d2b5579afc8a3a03bf72a87efe7b868"
// z4 as per Simon Spec
const z4 = [1,1,0,1,0,0,0,1,1,1,1,0,0,1,1,0,1,0,1,1,0,1,1,0,0,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,0,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,1,0,1,1,1,1]

// get difference of elements in list 
function diff(ary) {
    var newA = [];
    for (var i = 1; i < ary.length; i++)  newA.push(ary[i] - ary[i - 1])
    return newA;
}

// check if two arrays have equal content
function equalA(a1, a2){
    if(a1.length != a2.length){
        console.log("equalA Error")
    }
    else{
        for(var i = 0; i < a1.length; i++){
            if(a1[i]!=a2[i]){
                return false
            }
        }
        return true
    }
}

// hex string to binary list
function hex_to_binary(hex){
    var out = "";
    for(var c of hex) {
        switch(c) {
            case '0': out += "0000"; break;
            case '1': out += "0001"; break;
            case '2': out += "0010"; break;
            case '3': out += "0011"; break;
            case '4': out += "0100"; break;
            case '5': out += "0101"; break;
            case '6': out += "0110"; break;
            case '7': out += "0111"; break;
            case '8': out += "1000"; break;
            case '9': out += "1001"; break;
            case 'a': out += "1010"; break;
            case 'b': out += "1011"; break;
            case 'c': out += "1100"; break;
            case 'd': out += "1101"; break;
            case 'e': out += "1110"; break;
            case 'f': out += "1111"; break;
            default: return "";
        }
    }

    var list = []
    for(var i = 0; i < out.length;i++){
        if(out[i] == "1"){
            list.push(1)
        }
        else if(out[i] =="0"){
            list.push(0)
        }
        else{
            console.log("hex_to_binary error")
            return
        }
    }
    return list
}

function binary_to_hex(input){
    bin_string = input.toString()
    bin_string = bin_string.replace(/,/g,"");
    var hex = parseInt(bin_string, 2).toString(16);
    return hex
}

// shift array left 
function shift_left(input,times){
    result = Array.from(input)
    for(var i = 0; i < times; i++){
        first = result.shift()
        result.push(first)
    }
    return result 
}

// shift array right
function shift_right(input,times){
    result = Array.from(input)
    for (var i = 0; i < times; i++) {
        result.unshift(result.pop())
    }
    return result
}

// returned xor-ed arrays
function bit_xor(a1,a2){
    result = []
    for(var i = 0; i < a1.length; i++){
        result.push(a1[i]^a2[i])
    }
    return result
}

// returned and-ed arrays
function bit_and(a1,a2){
    result = []
    for(var i = 0; i < a1.length; i++){
        result.push(a1[i] & a2[i])
    }
    return result
}

function invert(a1){
    result = []
    for (var i = 0; i < a1.length; i++) {
        if(a1[i]==1){
            result.push(0)
        }
        else if(a1[i]==0){
            result.push(1)
        }
        else{
            console.log("invert error")
            break
        }
    }
    return result
}

function split(a1){
    half = a1.length/2
    high = Array.from(a1.slice(0,half))
    low = Array.from(a1.slice(half))
    return [high,low]
}

function generate_keys(key){
    big_halves = split(key)
    k_3_2 = split(big_halves[0])
    k_1_0 = split(big_halves[1])

    sub_keys = [k_1_0[1],k_1_0[0],k_3_2[1],k_3_2[0]]

    for(var i = 4; i < 72; i++){
        temp = shift_right(sub_keys[i-1],3)
        temp = bit_xor(temp,sub_keys[i-3])
        temp = bit_xor(temp,shift_right(temp,1))

        invert_kim = invert(sub_keys[i-4])
        temp = bit_xor(invert_kim,temp)

        z_bit = z4[((i-4) % 62)]            // xor 3 (11) with the two last bits in the key
        temp[temp.length - 1] = temp[temp.length - 1] ^ z_bit ^ 1
        temp[temp.length - 2] = temp[temp.length - 2] ^ 1
        sub_keys.push(temp)
    }
    return sub_keys
}

// rounds for Simon
function round(halves,key){
    left = halves[0]
    right = halves[1]

    left_shift_1 = shift_left(left,1)
    left_shift_2 = shift_left(left,2)
    left_shift_8 = shift_left(left,8)

    fx_res = bit_and(left_shift_1,left_shift_8)

    xor_1 = bit_xor(right,fx_res)
    xor_2 = bit_xor(xor_1,left_shift_2)
    xor_3 = bit_xor(xor_2,key)
    
    result_left = xor_3
    result_right  = left
    return [result_left,result_right]
}

function simon(binary_plain_text,sub_keys){
    halves = split(binary_plain_text)

    for(var i = 0; i < 72; i++){
        halves = round(halves,sub_keys[i])
        //console.log("Round " + i + ": " + binary_to_hex(halves[0]) + " " + binary_to_hex(halves[1]))
    }
    total = (halves[0]).concat(halves[1])
    return total
    
}

// decrypt by switching left and right in final rounds
function simon_d(binary_cipher_text,sub_keys){
    halves = split(binary_cipher_text)
    halves = halves.reverse()
    for (var i = 0; i < 72;i++){
        halves = round(halves,sub_keys[i])
    }
    return (halves[1]).concat(halves[0])
}

function average(a1){
    if(a1.length==0){
        return 0
    }
    else{
        sum = 0
        for(var i = 0; i < a1.length;i++){
            sum = sum + a1[i]
        }
        return (sum/a1.length)
    }
}

function main(){

    start_time = performance()

    binary_key = hex_to_binary(key)
    sub_keys = generate_keys(binary_key)
    
    // Encryption of Test Vector
    binary_plain_text = hex_to_binary(plaintext)
    binary_cipher_text = simon(binary_plain_text,sub_keys)
    encrypt_check = equalA(binary_cipher_text,hex_to_binary(ciphertext))

    // Decryption of Test Vector
    sub_keys = sub_keys.reverse()
    binary_cipher_text = hex_to_binary(ciphertext)
    binary_plain_text = simon_d(binary_cipher_text,sub_keys)
    decrypt_check = equalA(binary_plain_text,hex_to_binary(plaintext))

    end_time = performance()
    net_time = end_time-start_time

    mem = process.memoryUsage().heapUsed / 1024 / 1024; 
    mbs = Math.round(mem * 100) / 100
    cpu_stats = process.cpuUsage()// next line converts from microseconds to seconds
    return [encrypt_check,decrypt_check,net_time,mbs,cpu_stats.user/1000000,cpu_stats.system/1000000]
}

user_system_init = process.cpuUsage()

times = []
mems = []
cpus_user = []
cpus_system = []

cpus_user.push(user_system_init.user/1000000)
cpus_system.push(user_system_init.system/1000000)

for(var i = 0; i < 10; i++){
    results = main()
    times.push(results[2])
    mems.push(results[3])
    cpus_user.push(results[4])
    cpus_system.push(results[5])
}

console.log("--- Results of Simon in JavaScript --- \n")
console.log("Did Simon encrypt correctly? " + results[0])
console.log("Did Simon decrypt correctly? " + results[1])
console.log("Simon took an average of " + average(times) +  " seconds over 10 runs")
console.log("Simon used an average of " + average(mems) +  " MB of memory over 10 runs")
console.log("Simon spent an average of " + average(diff(cpus_user)) +  " seconds in the user CPU over 10 runs")
console.log("Simon spent an average of " + average(diff(cpus_system)) +  " seconds in the system CPU over 10 runs\n")

console.log("Times: " + times)
console.log("Space: " + mems )
console.log("CPU User: " + cpus_user)
console.log("CPU System: " + cpus_system + "\n")