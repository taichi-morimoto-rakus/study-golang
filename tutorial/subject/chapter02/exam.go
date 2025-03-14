package chapter02

const (
	Coin500         = 500
	Coin100         = 100
	Coin050         = 50
	Coin010         = 10
	Coin005         = 5
	Coin001         = 1
	TaxRate float64 = 0.1
)

func MinimumCoins(price uint) (count500, count100, count050, count010, count005, count001 uint) {
	taxedPrice := uint(float64(price) * (1 + TaxRate))
	count500 = taxedPrice / Coin500
	taxedPrice %= Coin500
	count100 = taxedPrice / Coin100
	taxedPrice %= Coin100
	count050 = taxedPrice / Coin050
	taxedPrice %= Coin050
	count010 = taxedPrice / Coin010
	taxedPrice %= Coin010
	count005 = taxedPrice / Coin005
	taxedPrice %= Coin005
	count001 = taxedPrice / Coin001
	return
}
