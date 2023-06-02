package main

import (
	"kunikida123456/NutritionApp/config"
	"kunikida123456/NutritionApp/handler"

	"github.com/labstack/echo/v4"
)

// @title Nutrition App API
// @description API server for the Nutrition App
// @version 1

func main() {
	e := echo.New()
	config.Connect()
	handler.InitRouting(e)
	e.Logger.Fatal(e.Start(":8088"))
}
