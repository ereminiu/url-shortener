package service

import "math/rand"

func shuffle(a []rune) {
	for rep := 0; rep < len(a)/2; rep++ {
		i, j := rand.Intn(len(a)), rand.Intn(len(a))

		a[i], a[j] = a[j], a[i]
	}
}
