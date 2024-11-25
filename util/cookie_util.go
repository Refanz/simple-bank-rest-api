package util

import (
	"bank-rest-api/constant"
	"net/http"
	"time"
)

func GetCookie(r *http.Request) *string {
	cookies, err := r.Cookie(constant.CookieName)

	if err != nil {
		return nil
	}

	return &cookies.Value
}

func SetCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     constant.CookieName,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(2 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    constant.CookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}

	http.SetCookie(w, cookie)
}
