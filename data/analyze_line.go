package data

import "regexp"

func AnalyzeLine(line string) []string {
	re := regexp.MustCompile(`[.,\s]+`) // " ", ", ", "."
	result := re.Split(line, -1)

	return result
}
