package ascii

import (
	"os"
	"strings"
)

func AsciiLogic(input string, bannerType string) (string, bool) {
	initialUserInput := strings.ReplaceAll(input, "\r\n", "\n")
	// Checking for potential non-ascii values
	isThereAnyOtherCharacters := false
	errorMessage := false
	isThereAnyOtherCharacters, errorMessage = checkForErrors(initialUserInput)
	data := ""
	if isThereAnyOtherCharacters {
		return data, errorMessage
	}
	// if there is any newline split it to the words by newlines
	userInput := strings.Split(initialUserInput, "\n")
	// choose banner type
	banner := bannerType
	if len(banner) == 0 {
		errorMessage = true
	}
	content := []byte{}
	var err error
	switch banner {
	case "standard":
		content, err = os.ReadFile("./files/standard.txt")
		if err != nil {
			errorMessage = true
		}

	case "shadow":
		content, err = os.ReadFile("./files/shadow.txt")
		if err != nil {
			errorMessage = true
		}
	case "thinkertoy":
		content, err = os.ReadFile("./files/thinkertoy.txt")
		content = []byte(strings.ReplaceAll(string(content), "\r\n", "\n"))
		if err != nil {
			errorMessage = true
		}
	case "":
		content, err = os.ReadFile("./thinkertoy.txt")
		if err != nil {
			errorMessage = true
		}
	}

	for _, l := range userInput {
		data += (getStr(l, getMap(string(content)), banner))
	}

	return data, errorMessage
}

func getMap(s string) map[rune]string {
	asciiMap := make(map[rune]string)
	temp := ""
	count := 0
	j := rune(32)

	for _, data := range s {
		temp += string(data)
		if data == '\n' {
			count++
		}
		if count == 9 {
			asciiMap[j] = temp[1 : len(temp)-1]
			temp = ""
			count = 0
			j++
		}
	}
	return asciiMap
}

func getStr(word string, asciiMap map[rune]string, banner string) string {
	temp := [8]string{}
	if word == "" {
		return "\n"
	}

	for _, char := range word {
		for index, value := range strings.Split(asciiMap[char], "\n") {
			temp[index] += value
		}
	}

	result := ""
	for _, l := range temp {
		result += l + "\n"
	}
	return result
}

func checkForErrors(s string) (bool, bool) {
	flag := false
	errorMessage := false

	// Error checking for non-ascii values
	for _, l := range s {
		if l < 32 || l > 126 {
			if l == 10 {
				continue
			}
			flag = true
			errorMessage = true
		}
	}
	return flag, errorMessage
}
