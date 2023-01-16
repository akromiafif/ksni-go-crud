package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"ksni.com/crud/database"
	"ksni.com/crud/model/entity"
	"ksni.com/crud/model/request"
	"ksni.com/crud/model/response"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	newUser := entity.User {
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": userResponse,
	})
}

func UserHandlerUpdateById(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	userId := ctx.Params("id")

	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": errUpdate,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func UserHandlerDeleteById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}