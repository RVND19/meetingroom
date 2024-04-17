package controllers

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"github.com/gofiber/fiber/v2"
)


type ReservationController struct {
	reservationRepo interfaces.IReservationRepo
	userRepo interfaces.IUserRepo
	roomRepo interfaces.IRoomRepo
}

func NewReservationController(reservationRepo interfaces.IReservationRepo, userRepo interfaces.IUserRepo, roomRepo interfaces.IRoomRepo) *ReservationController {
	return &ReservationController{reservationRepo: reservationRepo , userRepo: userRepo, roomRepo: roomRepo}
}

func (c *ReservationController) Route(app *fiber.App) {
	app.Put("/reservation", c.CreateReservation)
	app.Get("/reservation/:id", c.GetReservationById)
}


func (c *ReservationController) CreateReservation(ctx *fiber.Ctx) error {

	reservationdto := new(entities.CreateReservationDto)
	if err := ctx.BodyParser(&reservationdto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	if reservationdto.RoomId == 0 || reservationdto.Email == "" || reservationdto.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "room id , email and password are required", "data": nil})
	}

	start, err:=time.Parse("2006-01-02 15:04:05", reservationdto.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "parse start date error, please use format: YYYY-MM-DD HH:mm:ss", "data": nil})
	}

	end ,err:=time.Parse("2006-01-02 15:04:05", reservationdto.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "parse end date error, please use format: YYYY-MM-DD HH:mm:ss", "data": nil})
	}

	//check request time 
	if start.After(end) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "start date must be before end date", "data": nil})
	}

	//check user is verified or not
	user,err:=c.userRepo.FindUserByEmail(reservationdto.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	if !user.IsVerified {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "user is not verified", "data": nil})
	}

	//check password
	passwordToCheck:=fmt.Sprintf("%x", sha256.Sum256([]byte(reservationdto.Password+user.Salt)))
	if passwordToCheck != user.Password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "password is incorrect", "data": nil})
	}

	//check room is exist or not
	room,err:=c.roomRepo.GetById(reservationdto.RoomId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "room not found", "data": nil})
	}
	
	//check time slot is available or not
	isAvailable,err:=c.reservationRepo.IsRoomAvailable(room.Id, start, end)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	if !isAvailable {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "room and time is not available", "data": nil})
	}

	reservation:=entities.Reservation{
		RoomId: room.Id,
		UserId: user.Id,
		StartDate: start,
		EndDate: end,
	}

	err=c.reservationRepo.CreateReservation(&reservation)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Reservation created successfully", "data": reservation.Id})

}

func (c *ReservationController) GetReservationById(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	reservation, err := c.reservationRepo.GetById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error(), "data": nil})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": reservation})
}
