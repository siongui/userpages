package velthuis

import (
	"regexp"
	"strings"
)

var paliWordsInVelthuisScheme = regexp.MustCompile(`[abcdeghijklmnoprstuvyABCDEGHIJKLMNOPRSTUVY"~.]+`)

func velthuis2unicode(str string) (string, bool) {
	previousLetter := ""
	output := ""
	for i := 0; i < len(str); i++ {
		currentLetter := str[i : i+1]
		if i == 0 {
			output += currentLetter
			previousLetter = currentLetter
			continue
		}
		if previousLetter == "." {
			if currentLetter == "n" {
				output = output[0:len(output)-1] + "ṇ"
				previousLetter = currentLetter
				continue
			}
			if currentLetter == "m" {
				output = output[0:len(output)-1] + "ṃ"
				previousLetter = currentLetter
				continue
			}
			if currentLetter == "t" {
				output = output[0:len(output)-1] + "ṭ"
				previousLetter = currentLetter
				continue
			}
			if currentLetter == "d" {
				output = output[0:len(output)-1] + "ḍ"
				previousLetter = currentLetter
				continue
			}
			if currentLetter == "l" {
				output = output[0:len(output)-1] + "ḷ"
				previousLetter = currentLetter
				continue
			}
			return "", false
		}
		if previousLetter == "~" {
			if currentLetter == "n" {
				output = output[0:len(output)-1] + "ñ"
				previousLetter = currentLetter
				continue
			}
			return "", false
		}
		if previousLetter == "\"" {
			if currentLetter == "n" {
				output = output[0:len(output)-1] + "ṅ"
				previousLetter = currentLetter
				continue
			}
			return "", false
		}
		if previousLetter == "a" && currentLetter == "a" {
			output = output[0:len(output)-1] + "ā"
			previousLetter = currentLetter
			continue
		}
		if previousLetter == "i" && currentLetter == "i" {
			output = output[0:len(output)-1] + "ī"
			previousLetter = currentLetter
			continue
		}
		if previousLetter == "u" && currentLetter == "u" {
			output = output[0:len(output)-1] + "ū"
			previousLetter = currentLetter
			continue
		}
		output += currentLetter
		previousLetter = currentLetter
	}
	if output[len(output)-1:len(output)] == "." {
		return "", false
	}
	if output[len(output)-1:len(output)] == "\"" {
		return "", false
	}
	if str == output {
		return "", false
	}
	return output, true
}

func replacePaliWordsInVelthuisSchemeWithUnicode(b []byte) []byte {
	if len(b) == 1 {
		return b
	}
	str := strings.ToLower(string(b))
	if strings.HasSuffix(str, ".net") {
		return b
	}
	if strings.HasSuffix(str, ".tm") {
		return b
	}
	if strings.HasSuffix(str, ".lk") {
		return b
	}
	if output, ok := velthuis2unicode(str); ok {
		println(str + " => " + output)
		return []byte(output)
	}
	return b
}

func ProcessBytes(b []byte) []byte {
	return paliWordsInVelthuisScheme.ReplaceAllFunc(b,
		replacePaliWordsInVelthuisSchemeWithUnicode)
}
