package chapter01

import (
	"tutorial/helper"
)

// 初乗り料金
const firstPrice = 500

// 距離ごとに加算される料金
const perPrice = 100

// 初乗りの距離
const firstRideDistance = 1500

// 加算される距離
const perDistance = 250

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	if helper.ParseDistance(distance) <= firstRideDistance {
		return firstPrice, firstPrice * 1.2
	} else {
		normalPrice := firstPrice + perPrice*((helper.ParseDistance(distance)-firstRideDistance)/perDistance)
		latePrice := int(float64(normalPrice) * 1.2)
		return normalPrice, latePrice
	}
}
