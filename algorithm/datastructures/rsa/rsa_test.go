package rsa

import "testing"

func TestGcd(t *testing.T) {

	t.Log(greatestCommonDivisor1(4, 6))
	t.Log(greatestCommonDivisor2(4, 6))
	t.Log(greatestCommonDivisor3(4, 6))
}
