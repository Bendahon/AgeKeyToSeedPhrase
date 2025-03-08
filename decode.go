package main

import (
	"strconv"
	"strings"
)

// WordsToBinaryChunks converts words back to 11-bit binary chunks.
//func WordsToBinaryChunks(words []string, wordMap map[string]int64) []string {
//	chunks := []string{}
//	for _, word := range words {
//		if index, ok := wordMap[word]; ok {
//			chunks = append(chunks, fmt.Sprintf("%011b", index))
//		}
//	}
//	return chunks
//}

// BinaryToString converts a binary string back to a string.
func BinaryToString(binaryStr string) string {
	var result strings.Builder
	for i := 0; i < len(binaryStr); i += 8 {
		if i+8 > len(binaryStr) {
			break
		}
		charBinary := binaryStr[i : i+8]
		charInt, _ := strconv.ParseInt(charBinary, 2, 64)
		result.WriteByte(byte(charInt))
	}
	return result.String()
}
