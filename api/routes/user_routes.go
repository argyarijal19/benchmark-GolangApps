package routes

import (
	"golang-apps/api/handlers"
	"golang-apps/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UsersRouter(router fiber.Router, database *gorm.DB) {
	usersTable := repository.UserTable(database)

	usersHandler := handlers.UsersHandlers{UserTable: usersTable}
	usersRouter := router.Group("users")
	usersRouter.Get("/", usersHandler.GetAllUsers)
	usersRouter.Post("/register", usersHandler.InsertDataUsers)
}
