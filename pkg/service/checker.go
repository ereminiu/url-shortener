package service

import "regexp"

func ValidLink(s string) bool {
	// links begins with http:// or https://
	re := regexp.MustCompile(`https?://`)

	return re.MatchString(s)
}
