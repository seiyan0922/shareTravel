package common

import (
	"net/http"
	"strings"
	"time"
)

//クエリパラメータの取得処理
func GetQueryParam(r *http.Request) string {
	query := r.URL.Query().Encode()
	param := strings.Split(query, "=")[1]
	return param
}

//DBから取得した日付データをYYYY/MM/DD形式に変換
func TimeFormatter(datetime time.Time) string {

	formated := datetime.Format(TIME_LAYOUT_YYYYMMDD)

	return formated
}

func TimeFormatterHyphen(datetime string) string {
	arr := strings.Split(datetime, "-")
	year := arr[0]
	month := arr[1]
	arr2 := strings.Split(arr[2], "T")
	day := arr2[0]
	// var time string = strings.Replace(arr2[1], "Z", "", 1)

	formated := year + "-" + month + "-" + day

	return formated
}

func ContainsValueInt(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
