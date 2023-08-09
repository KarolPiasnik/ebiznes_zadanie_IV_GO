package main

import (
	"myapp/database"
	"myapp/routing"
)

func main() {
	e := routing.Init()
	database.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
