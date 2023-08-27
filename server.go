package main

import (
	"myapp/database"
	"myapp/routing"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {

	key := "not-so-secret-session-key"
	maxAge := 86400 * 30
	isProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New("your_google_id", "your_google_token", "http://localhost:1323/auth/google/callback", "email", "profile"),
	)

	e := routing.Init()
	database.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
