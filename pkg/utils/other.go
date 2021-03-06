package utils

import (
	"fmt"
	"time"
)

// FormatDate formats time into
// DD-MM-YYYY format
func FormatDate(t time.Time) string {
	y, m, d := t.Date()
	y %= 100
	return fmt.Sprintf("%02d/%02d/%02d", d, m, y)
}
