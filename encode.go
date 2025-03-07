package main

import (
	"fmt"
	"strconv"
	"strings"
)

// DecimalToWordList converts decimal values to words from the word list.
func DecimalToWordList(decimalValues []int64, wordList []string) []string {
	words := make([]string, len(decimalValues))
	for i, decimalValue := range decimalValues {
		if decimalValue < 0 || decimalValue >= int64(len(wordList)) {
			fmt.Printf("Error: decimal value %d out of range for word list.\n", decimalValue)
			return nil // Return nil on error
		}
		words[i] = wordList[decimalValue]
	}
	return words
}

// BinaryChunksToDecimal converts the binary chunks to decimal values.
func BinaryChunksToDecimal(chunks []string, padding int) []int64 {
	decimalValues := make([]int64, len(chunks))
	for i, chunk := range chunks {
		if i == len(chunks)-1 && padding > 0 {
			chunk = chunk[:11-padding]
		}
		decimalValue, err := strconv.ParseInt(chunk, 2, 64)
		if err != nil {
			fmt.Printf("Error parsing chunk %d: %v\n", i+1, err)
			return nil // Return nil on error
		}
		decimalValues[i] = decimalValue
	}
	return decimalValues
}

// StringToBinary converts a string to its binary representation.
func StringToBinary(inputStr string) string {
	var binaryStr strings.Builder
	for _, char := range inputStr {
		binaryStr.WriteString(fmt.Sprintf("%08b", char))
	}
	return binaryStr.String()
}

// PadAndSplitBinary adds padding to the end and splits a binary string into 11-bit chunks.
func PadAndSplitBinary(binaryStr string) ([]string, int) {
	length := len(binaryStr)
	padding := 11 - (length % 11)
	if padding == 11 {
		padding = 0
	}
	// Padding at the end
	paddedStr := binaryStr + strings.Repeat("0", padding)
	chunks := []string{}

	for i := 0; i < len(paddedStr); i += 11 {
		chunks = append(chunks, paddedStr[i:i+11])
	}
	return chunks, padding
}
