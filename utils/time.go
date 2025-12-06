package utils

import (
	"time"

	"github.com/jihanlugas/calendar/constant"
)

func ParseTime(t string) (time.Time, error) {
	return time.Parse(constant.FormatTimeLayout, t)
}
