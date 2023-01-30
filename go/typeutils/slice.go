package typeutils

import "strings"

func EmptyByteSliceNil(b []byte) []byte {
	if len(b) == 0 {
		return nil
	}
	return b
}

func SplitEmptyStringSliceEmpty(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func SplitEmptyStringSliceNil(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
