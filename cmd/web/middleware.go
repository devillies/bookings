package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// add crf protection
func NoSurf(next http.Handler) http.Handler {
	crfHandler := nosurf.New(next)
	crfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return crfHandler
}

//load and save session on request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
