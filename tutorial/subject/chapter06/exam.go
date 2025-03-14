package chapter06

import (
	"strconv"
	"strings"

	"tutorial/helper"
)

type AhoNumber int

func (n AhoNumber) String() string {
	return strconv.Itoa(int(n))
}

func (n AhoNumber) aho() string {
	return helper.AhoString(int(n))
}

func (n AhoNumber) Call() string {
	if n%3 == 0 || strings.Contains(strconv.Itoa(int(n)), "3") {
		return n.aho()
	} else {
		return n.String()
	}
}

func Nabeatsu(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		aho := AhoNumber(i).Call()

		res = append(res, aho)
	}
	return res
}
