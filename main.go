package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	infoText()

	// Read user choice
	if scanner.Scan() {
		choiceInput := scanner.Text()
		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			return
		}

		if choice == 1 {
			fmt.Print("Enter binary expression: ")
			if scanner.Scan() {
				expression := scanner.Text()
				result := convertBinaryToHex(expression)
				fmt.Println("Hexadecimal:", result)
			}
		}
		if choice == 2 {
			fmt.Print("Enter binary expression: ")
			if scanner.Scan() {
				expression := scanner.Text()
				result := convertBinaryToAscii(expression)
				fmt.Println("ASCII:", result)
			}
		}
		if choice == 3 {
			fmt.Print("Enter hexadecimal expression: ")
			if scanner.Scan() {
				expression := scanner.Text()
				result := convertHexToBinary(expression)
				fmt.Println("Binary:", result)
			}
		}
		if choice == 4 {
			fmt.Print("Enter hexadecimal expression: ")
			if scanner.Scan() {
				input := scanner.Text()
				result := convertHexToAscii(input)
				fmt.Println("ASCII:", result)
			}
		}
		if choice == 5 {
			fmt.Print("Enter ASCII expression: ")
			if scanner.Scan() {
				expression := scanner.Text()
				result := convertAsciiToBinary(expression)
				fmt.Println("Binary:", result)
			}
		}
		if choice == 6 {
			fmt.Print("Enter ASCII expression: ")
			if scanner.Scan() {
				expression := scanner.Text()
				result := convertAsciiToHex(expression)
				fmt.Println("Hexadecimal:", result)
			}
		}
	}
}

func convertBinaryToHex(input string) string{
	// Convert binary string to int
	decimal, err := strconv.ParseInt(input,2,64)
	if err != nil{
		return "invalid binary input"
	}
	// Convert int to hex
	hex := fmt.Sprintf("%X", decimal)
	return hex
}

func convertBinaryToAscii(input string) string {
	// Split binary string into chunks of 8 bits (1 byte per ASCII character)
	var asciiString strings.Builder
	for i := 0; i < len(input); i += 8 {
		// Ensure we don't exceed the input length
		end := i + 8
		if end > len(input) {
			end = len(input)
		}
		// Parse each 8-bit chunk as a binary number
		decimal, err := strconv.ParseInt(input[i:end], 2, 64)
		if err != nil {
			return "Invalid binary input"
		}
		// Convert to ASCII character
		ascii := fmt.Sprintf("%c", decimal)
		asciiString.WriteString(ascii)
	}
	
	return asciiString.String()
}

func convertHexToBinary(input string)string{
	// Convert hex string to int
	decimal, err := strconv.ParseInt(input, 16, 64)
	if err != nil{
		return "Invalid hex input"
	}
	// Convert int to binary
	binary := fmt.Sprintf("%b", decimal)
	return binary
}


func convertHexToAscii(input string)string{
	hexParts := strings.Split(input, " ")
	var asciiString strings.Builder

	for _, hexPart := range hexParts{
		decimal, err := strconv.ParseInt(hexPart, 16, 64)
		if err != nil{
			return "Invalid hex input"
		}
		ascii := fmt.Sprintf("%c", decimal)
		asciiString.WriteString(ascii)
	}

	return asciiString.String()
}

func convertAsciiToBinary(input string)string{
	var binaryString strings.Builder
	for _, letter := range input{
		binary := fmt.Sprintf("%08b", letter)
		binaryString.WriteString(binary)
	}
	return binaryString.String()
}

func convertAsciiToHex(input string)string{
	var hexString strings.Builder
	for _, letter := range input{
		hex := fmt.Sprintf("%X", letter)
		hexString.WriteString(hex)
	}
	return hexString.String()
}

func infoText(){
	fmt.Println("What kind of conversion do you want to preform?")
	fmt.Println("1. binary -> hex")
	fmt.Println("2. binary -> ascii")
	fmt.Println("3. hex -> binary")
	fmt.Println("4. hex -> ascii")
	fmt.Println("5. ascii -> binary")
	fmt.Println("6. ascii -> hex")
}