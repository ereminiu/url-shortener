package service

import (
	"log"

	"github.com/ereminiu/link-shorter/pkg/models"
	"github.com/ereminiu/link-shorter/pkg/repository"
)

func CreateLink(link string) (string, error) {
	// here I call hash function
	h := GetHash([]rune(link))
	log.Printf("hash = %s, original = %s \n", h, link)
	return h, repository.CreateLink(models.Link{Hashcode: h, Original: link})
}

func GetLink(code string) (string, error) {
	return repository.GetLink(code)
}

func CreateCustomLink(link models.CustomLink) error {
	return repository.CreateCustomLink(link)
}

func GetCustomLink(code string) (string, error) {
	return repository.GetCustomLink(code)
}

func LinkExist(code, tab, key string) bool {
	return repository.LinkExist(code, tab, key)
}
