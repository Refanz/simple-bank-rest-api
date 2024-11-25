package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCookie(t *testing.T) {
	const cookieName = "test_cookie"

	GetCookie := func(r *http.Request) *string {
		cookies, err := r.Cookie(cookieName)

		if err != nil {
			return nil
		}

		return &cookies.Value
	}

	t.Run("Cookie exists", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		request.AddCookie(&http.Cookie{
			Name:  cookieName,
			Value: "cookie_value",
		})

		result := GetCookie(request)

		if result == nil || *result != "cookie_value" {
			t.Errorf("Expected cookie value is 'cookie_value', but get %v", result)
		}
	})

	t.Run("Cookie does not exists", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		result := GetCookie(request)

		if result != nil {
			t.Errorf("expected nil, but get %v", result)
		}
	})
}
