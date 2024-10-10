package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/shlmvgleb/em-task/internal/models"
)

type SongsWithPagination struct {
	Result      []*models.Song `json:"result"`
	CurrentPage int            `json:"current_page"`
	PagesAmount int            `json:"pages_amount"`
}

type SongByIdWithVersePagination struct {
	Song         *models.Song `json:"song"`
	CurrentVerse int          `json:"current_verse"`
	VersesAmount int          `json:"verses_amount"`
}

type SongService struct {
	repo models.SongRepository
}

func NewSongService(sr models.SongRepository) *SongService {
	return &SongService{
		repo: sr,
	}
}

func (ss *SongService) AddSong(ctx context.Context, song *models.Song) error {
	err := ss.repo.Add(ctx, song)
	if err != nil {
		return fmt.Errorf("database error while creating a song: %w", err)
	}

	return nil
}

func (ss *SongService) GetAllSongsWithPagination(ctx context.Context, searchQuery string, limit int, page int) (*SongsWithPagination, error) {
	offset := (page * limit) - limit
	instances, count, err := ss.repo.GetWithSearchAndPagination(ctx, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	pagesCount := math.Floor(float64(count) / float64(limit))
	if count%limit != 0 {
		pagesCount += 1
	}

	return &SongsWithPagination{
		Result:      instances,
		PagesAmount: int(pagesCount),
		CurrentPage: page,
	}, nil
}

func (ss *SongService) GetSongById(ctx context.Context, id int64) (*models.Song, error) {
	song, err := ss.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (ss *SongService) CreateVersePagination(song *models.Song, page int) (*SongByIdWithVersePagination, error) {
	verses := strings.Split(song.Text, "\\n\\n")

	if len(verses)-1 < page-1 || page < 1 {
		return nil, errors.New("invalid verse page")
	}

	song.Text = verses[page-1]

	return &SongByIdWithVersePagination{
		Song:         song,
		CurrentVerse: page,
		VersesAmount: len(verses),
	}, nil
}

func (ss *SongService) UpdateSong(ctx context.Context, id int64, song models.Song) (*models.Song, error) {
	updated, err := ss.repo.Update(ctx, id, &song)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (ss *SongService) DeleteSong(ctx context.Context, id int64) error {
	err := ss.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
