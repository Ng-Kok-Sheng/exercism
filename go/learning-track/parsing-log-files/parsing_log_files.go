package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	matchedPattern := re.FindString(text)

	return matchedPattern != ""
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~|*|=|-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, v := range lines {
		matchedCount := re.FindStringSubmatch(v)
		count += len(matchedCount)
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User )+([ ]*[a-zA-Z0-9]*)`)
	newLines := []string{}

	for _, v := range lines {
		matchedStrings := re.FindStringSubmatch(v)

		if len(matchedStrings) > 0 {
			username := strings.TrimSpace(matchedStrings[2])
			newLines = append(newLines, fmt.Sprintf("[USR] %s %s", username, v))
		} else {
			newLines = append(newLines, v)
		}
	}
	return newLines
}
