package workout

import "math"

type Increase float64

const (
	TenPercent     Increase = 0.1
	TwentyPercent  Increase = 0.2
	ThirtyPercent  Increase = 0.3
	FortyPercent   Increase = 0.4
	FiftyPercent   Increase = 0.5
	SixtyPercent   Increase = 0.6
	SeventyPercent Increase = 0.7
	EightyPercent  Increase = 0.8
	NintyPercent   Increase = 0.9
	HundredPercent Increase = 1.0
)

func CalculateIncrease(number int, increase Increase) int {
	return int(math.Ceil(float64(number) * float64(increase)))
}
