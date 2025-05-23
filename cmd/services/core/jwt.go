package main

import (
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/jwtutil"
	"time"
)

const (
	halfDayHour = 12
	oneDayHour  = 24
)

func (a *application) initJWTSession(secret string) {
	jwtutil.InitJWTSession(
		secret,
		time.Hour*halfDayHour,
		time.Hour*oneDayHour,
	)
}
