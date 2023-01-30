package typeutils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StringDefault(s string, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	return s
}
func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func TruncateString(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	return fmt.Sprintf("%."+strconv.Itoa(length-len(suffix))+"s"+suffix, s)
}

func StringSnakeCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
