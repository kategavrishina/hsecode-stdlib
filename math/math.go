package math

import (
	"math"
)

func NthPrime(n int) int {
	if n < 1 {
		panic("n < 1 error")
	}
	if n == 1 {
		return 2
	}
	edge := 2 * (math.Floor(float64(n)*math.Log(float64(n))) + 1)
	integers := make([]bool, int(edge)/2)
	for i := 0; i < len(integers); i++ {
		integers[i] = true
	}
	for p := 3; p*p <= len(integers)*2; p += 2 {
		if integers[((p+1)/2)-2] == true {
			for i := p * 3; i <= (len(integers)*2)+1; i += 2 * p {
				integers[int(math.Floor(float64(i/2)))-1] = false
			}
		}
	}
	var primes []int
	for p := 0; len(primes) < n; p++ {
		if integers[p] == true {
			primes = append(primes, (p*2)+3)
		}
	}
	return primes[n-2]
}
