package typeutils

import (
	"fmt"
	"strconv"
)

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func TruncateString(s string, length int, suffix string) string {
	return fmt.Sprintf("%."+strconv.Itoa(length-len(suffix))+"s"+suffix, s)
}
