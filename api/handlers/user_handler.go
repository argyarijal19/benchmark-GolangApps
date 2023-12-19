package handlers

import (
	"golang-apps/helper"
	"golang-apps/models"
	"golang-apps/repository"

	"github.com/gofiber/fiber/v2"
)

type UsersHandlers struct {
	UserTable *repository.UserRepository
}

func (uh *UsersHandlers) GetAllUsers(ctx *fiber.Ctx) error {
	dataUsers, err := uh.UserTable.GetAllUser()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if dataUsers == nil {
		return fiber.NewError(fiber.StatusNotFound, "Users doesn't Exist")
	}

	return helper.SendResponse(ctx, fiber.StatusOK, true, "success get data", dataUsers)
}

func (uh *UsersHandlers) InsertDataUsers(ctx *fiber.Ctx) error {
	userData := new(models.UserModel)
	if err := ctx.BodyParser(userData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Request Body isn't Valid")
	}

	if err := uh.UserTable.RegisterUser(*userData); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helper.SendResponse(ctx, fiber.StatusOK, true, "success insert data", nil)
}
