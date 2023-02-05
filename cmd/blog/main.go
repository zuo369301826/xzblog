package main

import (
	"xzblog/internal/blog"
	"xzblog/log"
)

func main() {

	log.Info(" ====================================== ")
	log.Info(" =========== service start ============ ")
	log.Info(" ====================================== ")

	app := &blog.App{}
	app.Run()

	log.Info(" ====================================== ")
	log.Info(" ============ service end ============= ")
	log.Info(" ====================================== ")

	return
}
