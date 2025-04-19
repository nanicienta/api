package model

import "time"

type TokensModel struct {
	Access       string
	AccessExp    time.Time
	RefreshToken string
	RefreshHash  []byte
}
