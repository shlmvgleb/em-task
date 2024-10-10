package services

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/shlmvgleb/em-task/pkg/requests"
)

type SongDetailsApiService interface {
	FindSongDetails(ctx context.Context, group string, song string) (*SongDetails, error)
}

type SongDetails struct {
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

type SongDetailsMockApiService struct{}

func NewSongDetailsMockApiService() *SongDetailsMockApiService {
	return &SongDetailsMockApiService{}
}

const (
	apiUrl = "https://some-fancy-url.com"
)

const (
	songInfoRoute = "/info"
)

func (*SongDetailsMockApiService) FindSongDetails(ctx context.Context, group string, song string) (*SongDetails, error) {
	url, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}

	values := url.Query()
	values.Set("group", group)
	values.Set("song", song)

	url.RawQuery = values.Encode()

	// FYI: игнорирую ошибку и результат, из-за мока API-шки
	_, _ = requests.RequestWithJSON[any, SongDetails](
		ctx,
		http.DefaultClient,
		apiUrl+songInfoRoute,
		nil,
		nil,
	)

	date, _ := time.Parse("2006-01-02", "2006-16-07")
	mock := &SongDetails{
		ReleaseDate: date,
		Text:        "Ooh baby, don't you know I suffer?\\nOoh baby, can you hear me moan?\\nYou caught me under false pretenses\\nHow long before you let me go?\\n\\nOoh\\nYou set my soul alight\\nOoh\\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	return mock, nil
}
