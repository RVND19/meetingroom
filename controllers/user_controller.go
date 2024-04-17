package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userRepo interfaces.IUserRepo	
}

func NewUserController(userRepo interfaces.IUserRepo) *UserController {
	return &UserController{userRepo: userRepo}
}

func (c *UserController) Route(app *fiber.App) {
	app.Put("/user", c.Create)
}


func (c *UserController) Create(ctx *fiber.Ctx) error {
	userdto := new(entities.CreateUserDto)
	if err := ctx.BodyParser(&userdto); err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	if userdto.Name == "" || userdto.Email == "" || userdto.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "name, email and password are required", "data": nil})
	}

	user:=entities.User{
		Name: userdto.Name,
		Email: userdto.Email,
		Password: userdto.Password,
		IsAdmin: false,
		IsVerified: true, //dibuat auto true dulu
	}

	//random salt
	salt:=make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	user.Salt = base64.StdEncoding.EncodeToString(salt)

	//sha256
	user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password+user.Salt)))


	err = c.userRepo.CreateUser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "User created successfully", "data": user})
}
 
