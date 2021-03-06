package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"
)

func RandomString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func SetCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   true,
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, c)
}