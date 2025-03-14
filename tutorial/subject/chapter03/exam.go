package chapter03

import "tutorial/helper"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	if station == "東京" {
		return 0
	}

	currentStation := "東京"
	distance := 0
	for {
		currentStation = helper.InnerNextStation(currentStation)
		distance += helper.InnerLoopDistance(currentStation)
		if currentStation == station {
			break
		}
	}

	if distance <= 3999 {
		return 140
	} else if 4000 <= distance && distance <= 6999 {
		return 160
	} else if 7000 <= distance && distance <= 10999 {
		return 170
	} else if 11000 <= distance && distance <= 15999 {
		return 200
	} else if 16000 <= distance && distance <= 20999 {
		return 270
	} else if 21000 <= distance && distance <= 25999 {
		return 350
	} else if 26000 <= distance && distance <= 30999 {
		return 420
	} else {
		return 490
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	if station == "東京" {
		return 0
	}

	currentStation := "東京"
	distance := 0
	for {
		currentStation = helper.OuterNextStation(currentStation)
		distance += helper.OuterLoopDistance(currentStation)
		if currentStation == station {
			break
		}
	}

	switch {
	case distance <= 3999:
		return 140
	case 4000 <= distance && distance <= 6999:
		return 160
	case 7000 <= distance && distance <= 10999:
		return 170
	case 11000 <= distance && distance <= 15999:
		return 200
	case 16000 <= distance && distance <= 20999:
		return 270
	case 21000 <= distance && distance <= 25999:
		return 350
	case 26000 <= distance && distance <= 30999:
		return 420
	default:
		return 490
	}
}
