package function

import (
	"fmt"
	"math"
	"strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// Handle a serverless request
func Handle(req []byte) string {
	value, err := strconv.ParseInt(string(req), 10, 64)

	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	if value > int64(MaxInt) {
		return fmt.Sprintf("The numbeer is too large")
	}

	startStr := fmt.Sprintf("Starting generating load with value of %d\n", value)

	sqrtSum := float64(0)
	for i := int64(0); i < value; i++ {
		sqrtSum += math.Sqrt(float64(i))
	}
	finishStr := fmt.Sprintf("Result value %s\n", strconv.FormatFloat(sqrtSum, 'f', 10, 64))

	return startStr + finishStr
}
