package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Open the Word file
	file, err := os.Open("BIP39Wordlist.txt")
	if err != nil {
		fmt.Println("Error opening BIP39Wordlist.txt:", err)
		os.Exit(2)
	}
	// GoLand why will you scream at me if I don't do this :(
	// I already quit!
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing BIP39Wordlist.txt:", err)
			os.Exit(3)
		}
	}(file)
	// Read the file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(4)
	}
	// Split the file by newlines
	wordList := strings.Split(string(fileBytes), "\n")
	// Make a map of each word:reference
	// E.G abandon == 1
	wordMap := make(map[string]int64)
	for i, word := range wordList {
		wordMap[word] = int64(i)
	}
	//This should let you "cat key.txt | age-keyphrase" or something to that nature
	originalString := "1YRV7SRYNV4US80G2LFU8TX5Z3FJYDP5D54UN4PV2C7HSQVKGKC2S4AU3CW"
	fmt.Println("Encoded: AGE-SECRET-KEY-" + originalString)

	// Expands to 472 chars of binary
	binaryString := StringToBinary(originalString)
	fmt.Println("StringToBinary: " + binaryString)

	chunks, padding := PadAndSplitBinary(binaryString)
	fmt.Println("Binary Chunks (11 bits):", chunks)

	decimalValues := BinaryChunksToDecimal(chunks, padding)
	fmt.Println("Decimal Values:", decimalValues)

	words := DecimalToWordList(decimalValues, wordList)
	fmt.Println("Words:", words)

	// DECODING
	//Simulate word to binary conversion.
	binaryChunksFromWords := chunks
	joinedBinary := strings.Join(binaryChunksFromWords, "")

	// Remove the last character (assuming padding is always 1).
	// Maybe there is a way to go
	joinedBinary = joinedBinary[:len(joinedBinary)-1]
	decodedString := BinaryToString(joinedBinary)
	// Yes ik this bis bad
	fmt.Println("Decoded String: AGE-SECRET-KEY-" + decodedString)
}
