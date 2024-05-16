package util

import (
	"fmt"
	"math"
)

func FSize(size float64) string {
	var suffixes = [5]string{"Bytes", "KB", "MB", "GB", "TB"}
	if size == 0 {
		return fmt.Sprintf("%.0f %s", size, suffixes[0])
	}

	base := math.Log(float64(size)) / math.Log(1024)
	converted := math.Pow(1024, base-math.Floor(base))
	suffix := suffixes[int(math.Floor(base))]

	if size > 1023 {
		return fmt.Sprintf("%.1f %s", converted, suffix)
	}

	return fmt.Sprintf("%.0f %s", converted, suffix)
}
