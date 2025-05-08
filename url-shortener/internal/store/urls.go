package store

import (
	"database/sql"
	"fmt"
	"mirzaadr/url-shortener/internal/slug"
	"time"
)

type Urls struct {
	Slug        string    `json:"slug"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiredAt   time.Time `json:"expired_at"`
}

type UrlsStorage struct {
	db *sql.DB
}

func (s *UrlsStorage) Create(url string) (string, error) {
	query := `
		INSERT INTO urls (slug, original_url)
		VALUES ('', ?) RETURNING id;
	`

	var id int
	err := s.db.QueryRow(query, url).Scan(&id)
	if err != nil {
		return "", err
	}

	// Generate a slug from the auto-incremented id (Base62 encoding)
	slug := slug.GenerateSlug(id)
	// Update the slug field with the generated slug
	updateQuery := `UPDATE urls SET slug = $1 WHERE id = $2`
	_, err = s.db.Exec(updateQuery, slug, id)
	if err != nil {
		return "", err
	}

	return slug, nil
}

func (s *UrlsStorage) GetBySlug(slug string) (string, error) {
	query := `
		SELECT original_url FROM urls WHERE slug = ?;
	`
	var originalURL string
	err := s.db.QueryRow(query, slug).Scan(&originalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no URL found for the provided slug: %s", slug)
		}
		return "", err
	}
	return originalURL, nil
}
