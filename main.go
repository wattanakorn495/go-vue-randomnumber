package main 

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"math/rand"
	"strconv"
)

func main() {
	e := echo.New()
	e.File("/", "public/index.html")
	e.Static("/", "public")
	e.GET("/random", func(c echo.Context) error {
		number := rand.Int31n(1000000)
		output := strconv.Itoa(int(number))
		return c.JSON(http.StatusOK, output)
		})

	e.Use(middleware.Logger())
	e.Start(":8080")
}