package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrf := nosurf.New(next)

	csrf.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProdution,
		SameSite: http.SameSiteLaxMode,
	})

	return csrf
}

//Session Load and save the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}