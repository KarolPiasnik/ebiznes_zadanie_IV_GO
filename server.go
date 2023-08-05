package main

import (
	"myapp/database"
	"myapp/routing"
	"net/http"
)

func main() {
	e := routing.Init()
	database.Init()
	e.Logger.Fatal(e.Start(":1323"))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
