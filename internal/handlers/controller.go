package handlers

import (
	"github.com/shlmvgleb/em-task/internal/services"
)

type Controller struct {
	songService           *services.SongService
	songDetailsApiService services.SongDetailsApiService
}

func NewController(
	songService *services.SongService,
	songDetailsApiService services.SongDetailsApiService,
) *Controller {
	return &Controller{
		songService,
		songDetailsApiService,
	}
}
