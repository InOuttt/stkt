package main

import (
	"fmt"
	"strings"
)

var listString = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

func isAnagram(fs string, ss string) bool {
	// compare length
	if len(fs) != len(ss) {
		return false
	}

	// make it lowercase
	fs = strings.ToLower(fs)
	ss = strings.ToLower(ss)

	// compare each one by byte
	mapByte := make(map[int]int)
	for i := 0; i < len(fs); i++ {
		mapByte[int([]rune(fs)[i])-97]++
		mapByte[int([]rune(ss)[i])-97]--
	}
	for i := 0; i < 26; i++ {
		if mapByte[i] != 0 {
			return false
		}
	}

	return true
}

func groupOfAnagram(ls []string) [][]string {
	var checkedStr = make(map[string]bool)
	var resp [][]string

	for _, fStr := range ls {
		// skip paired string
		if checkedStr[fStr] {
			continue
		}

		var anagram []string
		for _, sStr := range ls {
			// add unpaired and anagram string to array
			if !checkedStr[sStr] && isAnagram(fStr, sStr) {
				anagram = append(anagram, sStr)
				checkedStr[sStr] = true // mark paired string
			}
		}

		resp = append(resp, anagram)
	}

	return resp
}

func main() {
	fmt.Println(groupOfAnagram(listString))
}
