package telegram

import (
	"github.com/thanhpk/randstr"
	"time"
)

var CurrentPassword string

func RunPasswordLoop() {
	for {
		CurrentPassword = randstr.String(8)
		time.Sleep(6 * time.Hour)
	}
}
