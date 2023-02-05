package blog

import (
	"xzblog/internal/httpserver"
)

type App struct {
}

func (m *App) Run() {

	httpserver.Run()

}

func (m *App) Stop() {

}
