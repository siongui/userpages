package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const js = `
<script type="text/javascript" >
    var test = {};
    test.campaign = "8d26113ba";
    test.isTest = "false";
    test.sitegroup = "Homepage";
</script>
`

func StringToLines(s string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func getCampaignNumber(line string) string {
	tmp := strings.Split(line, "=")[1]
	tmp = strings.TrimSpace(tmp)
	tmp = tmp[1 : len(tmp)-2]
	return tmp
}

func main() {
	lines := StringToLines(js)
	for _, line := range lines {
		if strings.Contains(line, "campaign") {
			result := getCampaignNumber(line)
			println(result)
		}
	}
}
