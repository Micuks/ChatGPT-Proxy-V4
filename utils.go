package main

import (
	"regexp"
)

func SplitByBlankLine(body_bytes []byte) [][]byte {
	str_normalized := regexp.MustCompile("\r\n").ReplaceAllString(string(body_bytes), "\n")
	str_slices := regexp.MustCompile(`\n\s*\n`).Split(str_normalized, -1)

	byte_slices := make([][]byte, len(str_slices))
	for i, s := range str_slices {
		byte_slices[i] = []byte(s)
		// println("[", i, "]", string(byte_slices[i]))
	}

	return byte_slices
}
