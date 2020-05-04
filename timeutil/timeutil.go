package timeutil

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func parseDurationSuffix(x string) (time.Duration, error) {
	x = strings.ToLower(x)
	switch x {
	case "s":
		return time.Second, nil
	case "m":
		return time.Minute, nil
	case "h":
		return time.Hour, nil
	case "d":
		return time.Hour * 24, nil
	default:
		return time.Second, errors.New(string(x) + " is not s/m/h/d")
	}
}

func ParseDuration(x string) (time.Duration, error) {
	if len(x) == 0 {
		return 0, errors.New("empty string")
	}

	suffix := x[len(x)-1 : len(x)]
	suffixDuration, err := parseDurationSuffix(suffix)
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseInt(x[:len(x)-1], 10, 64)
	if err != nil {
		return 0, err
	}

	return suffixDuration * time.Duration(num), nil
}
