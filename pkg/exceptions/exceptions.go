package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SongIdIsNotProvidedError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Error{
		Code:    http.StatusBadRequest,
		Message: songIdIsNotProvidedErrorMsg,
	})
}

func FailedToParseSongIdError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Error{
		Code:    http.StatusBadRequest,
		Message: failedToParseSongIdErrorMsg,
	})
}

func SongByIdNotFoundError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, Error{
		Code:    http.StatusNotFound,
		Message: songByIdNotFoundErrorMsg,
	})
}

func SongVerseNotFoundError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, Error{
		Code:    http.StatusNotFound,
		Message: songVerseNotFoundErrorMsg,
	})
}

func InvalidPayloadToCreateASongError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Error{
		Code:    http.StatusBadRequest,
		Message: invalidPayloadToCreateASongErrorMsg,
	})
}

func FetchingSongsError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Error{
		Code:    http.StatusInternalServerError,
		Message: fetchingSongsErrorMsg,
	})
}

func CreatingSongError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Error{
		Code:    http.StatusInternalServerError,
		Message: creatingSongErrorMsg,
	})
}

func FindSongDetailsError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Error{
		Code:    http.StatusInternalServerError,
		Message: findSongDetailsErrorMsg,
	})
}

func UpdatingSongError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Error{
		Code:    http.StatusInternalServerError,
		Message: updatingSongErrorMsg,
	})
}

func DeletingSongError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Error{
		Code:    http.StatusInternalServerError,
		Message: deletingSongErrorMsg,
	})
}
