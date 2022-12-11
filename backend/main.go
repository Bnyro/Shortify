package main

import (
	"github.com/url-shortener/database"
	"github.com/url-shortener/router"
)

func main() {
	database.Init()
	router.Create()
}
