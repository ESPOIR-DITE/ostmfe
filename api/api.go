package api

import (
	"gopkg.in/resty.v1"
	"ostmfe/config"
)

/**
This links to the real server
*/
//const BASE_URL string = "http://159.69.222.82:9000/ostm/"
//const BASE_URL string = "http://155.238.30.11:9000/ostm/"
const BASE_URL string = "http://localhost:9000/ostm/"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
