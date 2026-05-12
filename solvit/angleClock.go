package solvit

import "math"

func angleClock(hour int, minutes int) float64 {

	const (
		oneMinuteAngle = 6.0
		oneHourAngle = 30.0
		oneMinuteForHourAngle = .5
	)

	if hour >= 12 {
		hour -= 12
	}

	minutesAngle := float64(minutes) * oneMinuteAngle
	hourAngle := float64(hour) * oneHourAngle + float64(minutes) * oneMinuteForHourAngle

	angle := math.Abs(hourAngle - minutesAngle)
	if angle > 180 {
		angle = 360 - angle
	}

	return angle

}
