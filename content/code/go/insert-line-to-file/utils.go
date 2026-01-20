package insert

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func InsertStringToFile(path, str string, index int) error {
	var lines []string
	{
		f, err := os.Open(path)
		if err != nil {
			log.Print(err)
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Print(err)
		}

		defer f.Close()
	}

	lines[index] = str + lines[index]

	result := strings.Join(lines, "\n")

	return ioutil.WriteFile(path, []byte(result), 0644)
}
