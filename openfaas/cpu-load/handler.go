package function

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// Handle a serverless request
func Handle(req []byte) string {
	sec, err := strconv.ParseInt(string(req), 10, 64)

	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	if sec > int64(MaxInt) {
		return fmt.Sprintf("The numbeer is too large, bie")
	}

	startStr := fmt.Sprintf("Generating load for %d sec + 1 sec of sleep\n", sec)

	time.Sleep(time.Second)
	sqrtSum := float64(0)
	startTime := time.Now()
	for i := int64(0); time.Since(startTime) < time.Second*time.Duration(sec); i++ {
		sqrtSum += math.Sqrt(float64(i))
		if sqrtSum > 1e9 {
			sqrtSum = 0
		}
	}

	finishStr := fmt.Sprintf("Result sum: %s\n", strconv.FormatFloat(sqrtSum, 'f', 10, 64))

	return startStr + finishStr
}
