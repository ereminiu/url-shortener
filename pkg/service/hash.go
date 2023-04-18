package service

import (
	"math/rand"

	"github.com/ereminiu/url-shortener/pkg/repository"
)

var base = 37
var mod = int(1e8 + 7)
var sold []rune

// polynomial hash of string
func hash(s []rune) int {
	ans := 0
	for _, c := range s {
		ans = (ans*base + int(c-'a'+1)) % mod
	}
	return ans
}

func GetHash(s []rune) string {
	s = addSold(s)

	for true {
		N := 8
		digits := make([]rune, 0, N)
		// shuffle the s to make hash more undetermined
		shuffle(s)

		// translate hash number to string
		h := hash(s)
		for h > 0 {
			cur := rune('a')
			if rand.Int()%2 == 0 {
				cur = rune('a' + h%10 + rand.Intn(10))
			} else {
				cur = rune('A' + h%10 + rand.Intn(10))
			}
			digits = append(digits, cur)
			h /= 10
		}

		// add characters if len is less than N
		ln := len(digits)
		for i := 0; i < N-ln; i++ {
			cur := rune('a')
			if rand.Int()%2 == 0 {
				cur = rune('a' + rand.Intn(25))
			} else {
				cur = rune('A' + rand.Intn(25))
			}
			digits = append(digits, cur)
		}

		// change chars to digits in random positions
		for i := 0; i < 3; i++ {
			pz := rand.Intn(N)
			digits[pz] = rune('0' + rand.Intn(10))
		}

		s := string(digits)

		// check whether hash code exists in data base
		// if it is generate hash again
		if !repository.LinkExist(s, "links", "hcode") {
			return s
		}
	}
	panic("Can't generate unique hash :(")
}
