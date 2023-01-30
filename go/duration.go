package typeutils

import (
	"errors"
	"fmt"
	"time"
)

// Type Converters
func DurationStringToSeconds(s string) (int, error) {
	var a, b, c int
	n, err := fmt.Sscanf(s, "%d:%d:%d", &a, &b, &c)
	if err != nil {
		return 0, errors.New("error parsing duration string: '" + s + "'")
	}

	var durationString string
	switch n {
	case 3:
		durationString = fmt.Sprintf("%dh%dm%ds", a, b, c)
	case 2:
		durationString = fmt.Sprintf("%dm%ds", a, b)
	case 1:
		durationString = fmt.Sprintf("%ds", a)
	default:
		return 0, errors.New("error parsing duration string '" + s + "'")
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return 0, err
	}
	return int(duration.Seconds()), nil
}

func SecondsToDurationString(seconds int, layout string) string {
	if seconds == 0 {
		return ""
	}
	minutes := seconds / 60
	seconds = seconds % 60
	hours := minutes / 60
	minutes = minutes % 60

	switch layout {
	case "15":
		return fmt.Sprintf("%02d", hours)
	case "15:04":
		return fmt.Sprintf("%02d:%02d", hours, minutes)
	case "15:04:05":
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	case "04:05":
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	case "05":
		return fmt.Sprintf("%02d", seconds)
	default:
		return ""
	}
}
