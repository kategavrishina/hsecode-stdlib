package math

import (
	"math"
)

var primes = [70]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
	233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	283, 283, 293, 307, 311, 313, 317, 331, 337, 347}

func NthPrime(n int) int {
	if n < 1 {
		panic("n < 1 error")
	}
	if n <= 70 {
		return primes[n-1]
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
