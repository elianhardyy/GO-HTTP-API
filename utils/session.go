package utils

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func Session() {
	key := "mysecret"
	maxAge := 86400 * 30
	isProduction := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.HttpOnly = true
	store.Options.Secure = isProduction
	gothic.Store = store

}