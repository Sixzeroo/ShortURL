package main

import (
	"time"
	"math/rand"
	"regexp"
)

// ID 生成：时间戳 + 4位随机 + 4位机器数
func IdGenerator(machineId int) int64 {
	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	randNum := rand.Int() % 1 << 4
	id := (timestamp << 8) | int64(randNum  << 4) | int64(machineId)
	return id
}

// 变成62进制的数
func Int2Str(num int64) string {
	res := ""
	for ; num != 0 ; num /= 62 {
		remain := num % 62
		if remain < 10 {
			char := '0' + remain
			res += string(char)
		} else if remain < 36 {
			char := 'a' + remain - 10
			res += string(char)
		} else {
			char := 'A' + remain - 36
			res += string(char)
		}
	}
	return res
}

// 字符串变62进制数
func Str2Int(numStr string) int64 {
	var res int64
	for i := len(numStr)-1; i >= 0; i-- {
		res *= 62
		if numStr[i] >= '0' && numStr[i] <= '9' {
			res += int64(numStr[i] - '0')
		} else if numStr[i] >= 'a' && numStr[i] <= 'z'{
			res += int64(numStr[i] - 'a' + 10)
		} else {
			res += int64(numStr[i] - 'A' + 36)
		}
	}
	return res
}

func isUrl(url string) bool {
	reg, err := regexp.Compile(`(https?|ftp|file):\/\/[-A-Za-z0-9+&@#\/%?=~_|!:,.;]+[-A-Za-z0-9+&@#\/%=~_|]`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(url))
}


// 生成标识字符串进行判断
func isIdStr(idStr string) bool {
	reg, err := regexp.Compile(`[0-9a-zA-Z]{3,20}`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(idStr))
}