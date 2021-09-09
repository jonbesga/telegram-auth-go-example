package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

var TELEGRAM_BOT_TOKEN string = os.Getenv("TELEGRAM_BOT_TOKEN")

func VerifyTelegramData(data_check_string []byte, hash string) bool {
	secret := sha256.Sum256([]byte(TELEGRAM_BOT_TOKEN))
	h := hmac.New(sha256.New, []byte(secret[:]))
	h.Write([]byte(data_check_string))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha == hash
}
