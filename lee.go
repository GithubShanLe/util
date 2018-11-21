package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

var TenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

/*进制转换的键值查询*/
func findkey(in string) int {
	result := -1
	for k, v := range TenToAny {
		if in == v {
			result = k
		}
	}
	return result
}
func anyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findkey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
func NetMaskSysConvert(s string) string { //子网掩码的格式类型转换
	var tempNetMask string
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[:2], 16)) + "."
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[2:4], 16)) + "."
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[4:6], 16)) + "."
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[6:], 16))
	return tempNetMask
}

func BoolConvertToInt(property bool) int {
	if property {
		return 1
	} else {
		return 0
	}
}

//获取有效slice的长度和slice
func GetIndexByte(b []byte) ([]byte, int) {
	max := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '\u0000' {
			max = i
			break
		}
	}
	if max == 0 {
		max = len(b)
	}
	return b[0:max], max
}
func main() {
	netmask := NetMaskSysConvert("fffff000")
	log.Println(netmask)

	var head []byte
	h := "123"
	k := []byte(h)
	head = make([]byte, len(k)+2)

	for i := 0; i < len(k); i++ {
		head[i] = k[i]
	}
	log.Println(string(head))

	u, a := GetIndexByte(head)
	o, _ := strconv.Atoi(string(u))
	if o == 123 {
		log.Println("hhh", a)
	}
}
