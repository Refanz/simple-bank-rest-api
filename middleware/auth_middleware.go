package middleware

import "net/http"

func AuthMiddleware(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		_, err := w.Write([]byte("something went wrong"))
		if err != nil {
			panic(err)
		}
		return false
	}

	isValid := (username == "refan") && (password == "123")
	if !isValid {
		_, err := w.Write([]byte("login error"))
		if err != nil {
			return false
		}
	}

	return true
}
