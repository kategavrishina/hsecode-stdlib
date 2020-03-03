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
	edge := math.Ceil(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))
	integers := make([]bool, int(float64(int(edge)/2)))
	for p := 3; p*p <= len(integers)*2; p += 2 {
		if integers[((p+1)/2)-2] == false {
			for i := p * p; i <= (len(integers)*2)+1; i += 2 * p {
				integers[int(math.Floor(float64(i/2)))-1] = true
			}
		}
	}
	c := 0
	var prime int
	for p := 0; p < len(integers); p++ {
		if integers[p] == false {
			c += 1
		}
		if (c + 1) == n {
			prime = (p * 2) + 3
			break
		}
	}
	return prime
}
