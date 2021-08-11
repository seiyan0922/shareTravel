package common

import (
	"net/http"
	"strings"
)

func GetQueryParam(r *http.Request) string {
	query := r.URL.Query().Encode()
	param := strings.Split(query, "=")[1]
	return param
}

// func TimeFormatter(time string) string {

// }
