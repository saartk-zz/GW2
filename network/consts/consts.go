package consts

import "time"

const (
	Base_GW_URL = "https://api.guildwars2.com"
)

const(
	Version_2 = "/v2"
)

const (
	Account_Path = "/account"
)

const(
	AuthHeaderKey = "Authorization"
	AuthHeaderValue = "Bearer %s"
)

const (
	DefaultHttpClientTimeout = 10 * time.Second
)