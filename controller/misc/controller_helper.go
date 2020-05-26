package misc

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

const (
	YYYYMMDD_FORMAT    = "2006-01-02"
	YYYMMDDTIME_FORMAT = "2006-01-02 15:04:05"
)

/**
Format date in yyyy-MM-dd HH:mm:ss
*/

func FormatDateTime(date time.Time) string {
	return date.Format(YYYMMDDTIME_FORMAT)
}

/**
format date in yyyy-MM-dd
*/
func FormatDate(date time.Time) string {
	return date.Format(YYYYMMDD_FORMAT)
}

/***
this method should separates longitude and latitude
*/
func SeparateLatLng(latlng string) (string, string) {
	var longitude string
	var latitude string
	val := strings.TrimSuffix(latlng, ")")
	val2 := strings.TrimPrefix(val, "(")
	parts := strings.Split(val2, ",")
	if len(parts) != 2 {
		return latitude, longitude
	}
	latitude = parts[0]
	longitude = parts[1]
	return latitude, longitude
}

func CheckFiles(files []io.Reader) [][]byte {
	var bytelist [][]byte
	for index, reablefile := range files {
		if reablefile != nil {
			reader := bufio.NewReader(reablefile)
			content, _ := ioutil.ReadAll(reader)
			bytelist = append(bytelist, content)
			fmt.Println(" done with file: ", index)
		}
	}
	return bytelist
}
