package math

import (
	"math"
)

func NthPrime(n int) int {

	if n < 1 {
		panic("n < 1 error")
	}
	edge := 2 * (math.Floor(float64(n)*math.Log(float64(n))) + 1)
	integers := make([]bool, int(edge)+1)
	for i := 2; i < int(edge); i++ {
		integers[i] = true
	}

	for p := 2; p*p <= int(edge); p++ {
		if integers[p] == true {
			for i := p * 2; i <= int(edge); i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= int(edge); p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes[n-1]
}
