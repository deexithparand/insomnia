package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func CMDInputRuleSet(input string) bool {
	// allow only numbers [0-9]
	// allow only lowercase characters
	// allow only hypen

	for _, char := range input {
		if (char >= 'a' && char <= 'z') || char == '-' || unicode.IsDigit(char) {
			continue
		}
		return false
	}
	return true
}

func GetCMDInput(promptStr string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	// give user the prompt
	fmt.Print(promptStr)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading the cmd input %v", err)
	}

	// trim spaces
	return strings.TrimSpace(input), nil
}
