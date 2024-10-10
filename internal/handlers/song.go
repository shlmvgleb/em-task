package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shlmvgleb/em-task/internal/models"
	"github.com/shlmvgleb/em-task/pkg/exceptions"
	"github.com/sirupsen/logrus"
)

const (
	defaultSongsLimit = 10
	defaultPage       = 1

	defaultVersePage = 1
)

type AddSongPayload struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

// GetSongsWithPagination godoc
// @Summary Получение всех песен с пагинацией
// @Description Возвращает информацию о песне по указанному ID с поддержкой пагинации (если она нужна).
// @Tags songs
// @Accept  json
// @Produce  json
// @Param   page              query     int        false   "Страница"
// @Param   limit             query     int        false   "Количество элементов"
// @Param   search_query      query     string     false   "Полнотекстовый поиск по всем полям сущности Song"
// @Success 200 {object} []services.SongsWithPagination   "Успешный ответ, песня найдена"
// @Failure 400 {object} exceptions.Error                 "Неверный запрос, ID песни не предоставлен или некорректен"
// @Failure 404 {object} exceptions.Error                 "Песня с предоставленным ID не найдена"
// @Failure 500 {object} exceptions.Error                 "Внутренняя ошибка сервера"
// @Router  /songs [get]
func (cntrl *Controller) GetSongsWithPagination(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit == 0 {
		limit = defaultSongsLimit
	}

	if page == 0 {
		page = defaultPage
	}

	searchQuery := c.Query("search_query")
	songs, err := cntrl.songService.GetAllSongsWithPagination(c.Request.Context(), searchQuery, limit, page)
	if err != nil {
		logrus.Debugf("get all songs with pagination error: %s", err)
		exceptions.FetchingSongsError(c)
		return
	}

	c.JSON(http.StatusOK, songs)
}

// GetSongByIdWithVersePagination godoc
// @Summary Получение песни по ID с пагинацией по куплетам
// @Description Возвращает информацию о песне по указанному ID с поддержкой пагинации (если она нужна).
// @Tags songs
// @Accept  json
// @Produce  json
// @Param   id      path     int     true    "ID песни"
// @Param   page    query    int     false   "Страница(порядковый номер куплета)"
// @Success 200 {object} services.SongByIdWithVersePagination  "Песня успешно найдена"
// @Failure 400 {object} exceptions.Error                      "Неверный запрос, ID песни не предоставлен или некорректен"
// @Failure 404 {object} exceptions.Error                      "Песня с предоставленным ID не найдена"
// @Failure 500 {object} exceptions.Error                      "Внутренняя ошибка сервера"
// @Router  /songs/paginated/{id} [get]
func (cntrl *Controller) GetSongByIdWithVersePagination(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		exceptions.SongIdIsNotProvidedError(c)
		return
	}

	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		exceptions.FailedToParseSongIdError(c)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = defaultVersePage
	}

	song, err := cntrl.songService.GetSongById(c.Request.Context(), intId)
	if err != nil {
		logrus.Debugf("get song by id error: %s", err)
		exceptions.SongByIdNotFoundError(c)
		return
	}

	paginated, err := cntrl.songService.CreateVersePagination(song, page)
	if err != nil {
		logrus.Debugf("create verse pagination error: %s", err)
		exceptions.SongVerseNotFoundError(c)
		return
	}

	c.JSON(http.StatusOK, paginated)
}

// GetSongById godoc
// @Summary Получение песни по ID
// @Description Возвращает информацию о песне по указанному ID
// @Tags songs
// @Accept  json
// @Produce json
// @Param   id      path     int     true    "ID песни"
// @Success 200 {object} services.SongByIdWithVersePagination  "Песня успешно найдена"
// @Failure 400 {object} exceptions.Error                      "Неверный запрос, ID песни не предоставлен или некорректен"
// @Failure 404 {object} exceptions.Error                      "Песня с предоставленным ID не найдена"
// @Failure 500 {object} exceptions.Error                      "Внутренняя ошибка сервера"
// @Router  /songs/{id} [get]
func (cntrl *Controller) GetSongById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		exceptions.SongIdIsNotProvidedError(c)
		return
	}

	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		exceptions.FailedToParseSongIdError(c)
		return
	}

	song, err := cntrl.songService.GetSongById(c.Request.Context(), intId)
	if err != nil {
		logrus.Debugf("get song by id error: %s", err)
		exceptions.SongByIdNotFoundError(c)
		return
	}

	c.JSON(http.StatusOK, song)
}

// AddSong godoc
// @Summary Добавление новой песни
// @Description Создает новую запись о песне
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body AddSongPayload true "Данные песни для добавления"
// @Success 201 {object} any              "Песня успешно добавлена"
// @Failure 400 {object} exceptions.Error "Некорректный запрос, неправильный формат данных"
// @Failure 500 {object} exceptions.Error "Внутренняя ошибка сервера"
// @Router /songs [post]
func (cntrl *Controller) AddSong(c *gin.Context) {
	var payload AddSongPayload
	err := c.BindJSON(&payload)
	if err != nil {
		exceptions.InvalidPayloadToCreateASongError(c)
		return
	}

	details, err := cntrl.songDetailsApiService.FindSongDetails(c.Request.Context(), payload.Group, payload.Song)
	if err != nil {
		logrus.Debugf("find song details error: %s", err)
		exceptions.FindSongDetailsError(c)
		return
	}

	song := models.Song{
		Group:       payload.Group,
		Song:        payload.Song,
		Text:        details.Text,
		Link:        details.Link,
		ReleaseDate: details.ReleaseDate,
	}

	err = cntrl.songService.AddSong(c.Request.Context(), &song)
	if err != nil {
		logrus.Debugf("add song error: %s", err)
		exceptions.CreatingSongError(c)
		return
	}

	c.Status(http.StatusCreated)
}

// UpdateSong godoc
// @Summary Обновление данных песни
// @Description Обновляет запись о песне
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body models.Song true "Данные песни для обновления"
// @Success 201 {object} models.Song "Песня успешно обновлена"
// @Failure 400 {object} exceptions.Error "Некорректный запрос, неправильный формат данных"
// @Failure 500 {object} exceptions.Error "Внутренняя ошибка сервера"
// @Router /songs [patch]
func (cntrl *Controller) UpdateSong(c *gin.Context) {
	var songPayload models.Song
	err := c.BindJSON(&songPayload)
	if err != nil {
		exceptions.InvalidPayloadToCreateASongError(c)
		return
	}

	song, err := cntrl.songService.UpdateSong(c.Request.Context(), songPayload.Id, songPayload)
	if err != nil {
		logrus.Debugf("update song error: %s", err)
		exceptions.UpdatingSongError(c)
		return
	}

	c.JSON(http.StatusCreated, song)
}

// DeleteSong godoc
// @Summary Удаление песни
// @Description Удаляет песню по ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "ID песни"
// @Success 204 {object} any              "Песня успешно удалена"
// @Failure 400 {object} exceptions.Error "Некорректный запрос, неправильный формат данных"
// @Failure 500 {object} exceptions.Error "Внутренняя ошибка сервера"
// @Router /songs/{id} [delete]
func (cntrl *Controller) DeleteSong(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		exceptions.SongIdIsNotProvidedError(c)
		return
	}

	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		exceptions.FailedToParseSongIdError(c)
		return
	}

	err = cntrl.songService.DeleteSong(c.Request.Context(), intId)
	if err != nil {
		logrus.Debugf("delete song error: %s", err)
		exceptions.DeletingSongError(c)
		return
	}

	c.Status(http.StatusNoContent)
}
