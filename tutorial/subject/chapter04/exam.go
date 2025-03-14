package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	if card.Balance+card.Point < charge {
		return false
	} else {
		if card.Point < charge {
			card.Balance -= charge - card.Point
			card.Point = 0
		} else {
			card.Point -= charge
		}
		return true
	}
}
