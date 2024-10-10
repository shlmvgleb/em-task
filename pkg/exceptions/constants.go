package exceptions

const (
	songIdIsNotProvidedErrorMsg         = "Song ID is not provided."
	failedToParseSongIdErrorMsg         = "Failed to parse song ID. Invalid value passed."
	songByIdNotFoundErrorMsg            = "Song with provided ID is not found."
	songVerseNotFoundErrorMsg           = "Provided song verse is not found."
	invalidPayloadToCreateASongErrorMsg = "Passed invalid payload to create a song."
	fetchingSongsErrorMsg               = "Unknown error while fetching songs."
	creatingSongErrorMsg                = "Unknown error while creating a song."
	findSongDetailsErrorMsg             = "Unknown error while finding song details."
	updatingSongErrorMsg                = "Unknown error while updating a song."
	deletingSongErrorMsg                = "Unknown error while deleting a song by id."
)
