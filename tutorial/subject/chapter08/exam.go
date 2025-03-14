package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

// Car を満たすように実装
type Basic struct{}

func (b Basic) PricePer15Minutes() int {
	return 220
}

func (b Basic) MaxPrice() int {
	return 4290
}

// Car を満たすように実装
type Middle struct{}

func (m Middle) PricePer15Minutes() int {
	return 330
}

func (m Middle) MaxPrice() int {
	return 6490
}

// Car を満たすように実装
type Premium struct{}

func (p Premium) PricePer15Minutes() int {
	return 440
}

func (p Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	overMaxFeeMinutes := minutes % 360
	overMaxFeeCount := minutes / 360
	notIncludeMaxPrice := (overMaxFeeMinutes/15 + 1) * car.PricePer15Minutes()

	if notIncludeMaxPrice >= car.MaxPrice() {
		return car.MaxPrice() * (overMaxFeeCount + 1)
	} else {
		return notIncludeMaxPrice + car.MaxPrice()*overMaxFeeCount
	}
}
