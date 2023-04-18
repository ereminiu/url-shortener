package service

import "math/rand"

func addSold(s []rune) []rune {
	N := 20
	res := make([]rune, len(s), len(s)+N)
	copy(res, s)

	for i := 0; i < N; i++ {
		if rand.Int()%2 == 0 {
			res = append(res, rune('a'+rand.Intn(25)))
		} else {
			res = append(res, rune('A'+rand.Intn(25)))
		}
	}

	return res
}
