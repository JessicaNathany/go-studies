package math

import (
	"fmt"
	"strconv"
)

// Media has a responsibility for calculate that you know

func Media(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	media := total / float64(len(numbers))
	mediaRound, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaRound
}
