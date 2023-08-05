package common

import (
	"fmt"
	"time"
)

func TimerStart() time.Time {
	return time.Now()
}

func TimerEnd(start time.Time) time.Duration {
	return time.Now().Sub(start)
}

func PrintTimer(start time.Time) {
	fmt.Printf("\nTime elapsed: %s\n", TimerEnd(start))
}
