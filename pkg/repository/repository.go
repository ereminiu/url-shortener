package repository

import (
	"fmt"

	"github.com/ereminiu/link-shorter/pkg/models"
)

func CreateLink(link models.Link) error {
	query := fmt.Sprintf("insert into links (hcode, original) values ($1, $2)")
	_, err := db.Exec(query, link.Hashcode, link.Original)
	return err
}

func GetLink(code string) (string, error) {
	query := fmt.Sprintf("select hcode, original from links where hcode = $1")
	row := db.QueryRow(query, code)

	var hashCode string
	var original string
	// return empty string and error if link not found
	if err := row.Scan(&hashCode, &original); err != nil {
		return "", err
	}

	return original, nil
}

func LinkExist(code, tab, key string) bool {
	query := fmt.Sprintf("select * from %s where %s = $1", tab, key)
	row := db.QueryRow(query, code)

	var id int
	var hashCode string
	var original string
	if err := row.Scan(&id, &hashCode, &original); err != nil {
		return false
	}

	return true
}

func CreateCustomLink(link models.CustomLink) error {
	query := fmt.Sprintf("insert into customlinks (custom, original) values ($1, $2)")
	_, err := db.Exec(query, link.CustomCode, link.Link)
	return err
}

func GetCustomLink(code string) (string, error) {
	query := fmt.Sprintf("select custom, original from customlinks where custom = $1")
	row := db.QueryRow(query, code)

	var customCode string
	var original string
	// return empty string and error if link not found
	if err := row.Scan(&customCode, &original); err != nil {
		return "", err
	}

	return original, nil
}
