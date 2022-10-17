package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	web := echo.New()

	// Middleware
	web.Use(middleware.Logger())
	web.Use(middleware.Recover())

	// Routes
	//используем RPC style, не REST
	web.POST("/findTickets", hello)   //поиск билетов по фильтрам
	web.POST("/bookTicket", hello)    //бронирование билета
	web.POST("/buyTicket", hello)     //оплата забронированного билета
	web.POST("/cancelBooking", hello) //отмена бронирования
	web.POST("/changeBooking", hello) //внесение изменений в бронирование

	// Start server
	web.Logger.Fatal(web.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
