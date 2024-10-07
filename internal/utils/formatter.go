package utils

func GetBearerHeader(token string) string {
	return "Bearer " + token
}
