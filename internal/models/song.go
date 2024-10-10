package models

import (
	"context"
	"time"
)

type SongRepository interface {
	GetWithSearchAndPagination(ctx context.Context, searchQuery string, limit int, offset int) ([]*Song, int, error)
	GetById(ctx context.Context, id int64) (*Song, error)
	Add(ctx context.Context, song *Song) error
	Update(ctx context.Context, id int64, song *Song) (*Song, error)
	Delete(ctx context.Context, id int64) error
}

// Song представляет информацию о песне
// @Description Структура, содержащая данные о песне, такие как группа, название песни, текст, дата выпуска и ссылка.
// @Tags songs
type Song struct {
	Id          int64     `json:"id,omitempty"`
	Group       string    `json:"group,omitempty"`
	Song        string    `json:"song,omitempty"`
	Text        string    `json:"text,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Link        string    `json:"link,omitempty"`
}
