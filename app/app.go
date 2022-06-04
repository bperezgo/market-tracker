package app

import (
	"markettracker.com/health"

	"github.com/gin-gonic/gin"
)

var appInstance *App

type App struct {
	app *gin.Engine
}

func GetInstance() *App {
	if appInstance != nil {
		return appInstance
	}
	app := gin.Default()
	return &App{app: app}
}

func (a *App) Start() {
	a.app.GET("/api/health", health.Handler)
}

func (a *App) Run(addr string) error {
	return a.app.Run(addr)
}
