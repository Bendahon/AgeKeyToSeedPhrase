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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing BIP39Wordlist.txt:", err)
			os.Exit(3)
		}
	}(file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(4)
	}

	wordList := strings.Split(string(fileBytes), "\n")

	wordMap := make(map[string]int64)
	for i, word := range wordList {
		wordMap[word] = int64(i)
	}

	originalString := "13QR"
	fmt.Println("Encoded:", originalString)

	binaryString := StringToBinary(originalString)
	fmt.Println("StringToBinary: " + binaryString)

	chunks, padding := PadAndSplitBinary(binaryString)
	fmt.Println("Binary Chunks (11 bits):", chunks)

	decimalValues := BinaryChunksToDecimal(chunks, padding)
	fmt.Println("Decimal Values:", decimalValues)

	words := DecimalToWordList(decimalValues, wordList)
	fmt.Println("Words:", words)

	//DECODING
	binaryChunksFromWords := WordsToBinaryChunks(words, wordMap)
	fmt.Println("Binary Chunks from words:", binaryChunksFromWords)

	//Remove padding.
	joinedBinary := strings.Join(binaryChunksFromWords, "")
	if padding > 0 {
		joinedBinary = joinedBinary[:len(joinedBinary)-padding]
	}

	fmt.Println("Joined Binary:", joinedBinary)

	//Convert binary back to string.
	decodedString := BinaryToString(joinedBinary)
	fmt.Println("Decoded String:", decodedString)
}
