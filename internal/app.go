package internal

import (
	"github.com/cake-store/internal/cake"
	"github.com/cake-store/internal/database"
	"github.com/cake-store/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Router *echo.Echo
}

func (a *App) Initialize() {
	a.Router = echo.New()
	a.initializeMiddleware()
	a.initializeRoutes()
}

func (a *App) initializeMiddleware() {
	a.Router.Use(middleware.Logger())
	a.Router.Use(middleware.Recover())
}

func (a *App) initializeRoutes() {
	db := database.NewMySQLDB()
	cakeRepo := cake.NewCakeRepository(db)
	cakeService := cake.NewCakeService(cakeRepo)
	cakeHandler := handler.NewCakeHandler(cakeService)

	a.Router.POST("/cakes", cakeHandler.CreateCake)
	a.Router.GET("/cakes", cakeHandler.GetAllCakes)
	a.Router.GET("/cakes/:id", cakeHandler.GetCakeByID)
	a.Router.PUT("/cakes/:id", cakeHandler.UpdateCake)
	a.Router.DELETE("/cakes/:id", cakeHandler.DeleteCake)
}
