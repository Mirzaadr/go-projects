package store

import "database/sql"

type Storage struct {
	Urls interface {
		Create(url string) (string, error)
		GetBySlug(slug string) (string, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Urls: &UrlsStorage{db},
	}
}
