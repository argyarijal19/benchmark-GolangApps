package main

import (
	"encoding/json"
	"golang-apps/api/middleware"
	"golang-apps/api/routes"
	"golang-apps/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: helper.ExceptionHandler,
	})

	databaseConn, err := middleware.InitGormDB()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	app.Use(logger.New())
	configCors := cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Access-Control-Allow-Headers",
	}

	app.Use(cors.New(configCors))
	app.Use(middleware.GormMiddleware(databaseConn))

	routes.UsersRouter(app, databaseConn)
	app.Listen(":8080")
}
