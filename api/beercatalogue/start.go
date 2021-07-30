package beercatalogue

import (
	"github.com/leumas3003/beercatalogue-go/docs"
	"github.com/leumas3003/beercatalogue-go/internal/pkg/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server BeerCatalogue server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email samuel.martel@derivco.se

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func StartServer() *echo.Echo {

	e := echo.New()
	setupRoutes(e)
	confSwagger()
	e.Logger.Fatal(e.Start(":3001"))
	return e
}

func setupRoutes(e *echo.Echo) {
	e.POST("/addbeer", handler.AddBeer)
	e.GET("/beer/:name", handler.GetBeer)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func confSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "BeerCatalogue API"
	docs.SwaggerInfo.Description = "Webserver using echo. PostgreSQL"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
