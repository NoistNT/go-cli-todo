package helpers

import (
	"bufio"
	"fmt"
	"strings"
)

// PrintCommands prints the list of commands
func PrintCommands() {
	fmt.Println("Commands available: [ add | list | update | done | remove | exit | quit ]")
}

// GetInput gets user input using a buffered reader
func GetInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// GetInputInt gets user input using GetInput and converts it to an int
func GetInputInt(reader *bufio.Reader, prompt string) (int, error) {
	input, err := GetInput(reader, prompt)
	if err != nil {
		return 0, err
	}
	var parsedInput int
	fmt.Sscan(input, &parsedInput)
	return parsedInput, nil
}

// ErrorHandler recieves an error and prints the error message
func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
