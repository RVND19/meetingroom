package controllers

import (
	"fmt"

	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"github.com/gofiber/fiber/v2"
)

type RoomController struct {
	roomRepo interfaces.IRoomRepo
}

func NewRoomController(roomRepo interfaces.IRoomRepo) *RoomController {
	return &RoomController{roomRepo: roomRepo}
}

func (c *RoomController) Route(app *fiber.App) {
	app.Put("/room", c.Create)
}


func (c *RoomController) Create(ctx *fiber.Ctx) error {

	roomdto := new(entities.CreateRoomDto)
	if err := ctx.BodyParser(&roomdto); err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	if roomdto.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "name is required", "data": nil})
	}

	room:=entities.Room{
		Name: roomdto.Name,
	}

	err := c.roomRepo.CreateRoom(&room)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Room created successfully", "data": room})
}